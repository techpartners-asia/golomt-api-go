package openbank

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

func PKCS7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func (g openbank) EncryptAESCBC(text string) (string, error) {
	block, err := aes.NewCipher([]byte(g.sessionKey))
	if err != nil {
		return "", err
	}

	if len(g.ivKey) != aes.BlockSize {
		return "", fmt.Errorf("IV length must be %d bytes", aes.BlockSize)
	}

	// Pad the plaintext
	plaintext := PKCS7Pad([]byte(text), aes.BlockSize)

	encrypted := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, []byte(g.ivKey))
	mode.CryptBlocks(encrypted, plaintext)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func PKCS7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	padding := int(data[length-1])
	if padding > length || padding == 0 {
		return nil, fmt.Errorf("invalid PKCS#7 padding")
	}
	return data[:length-padding], nil
}

func (g openbank) DecryptAESCBC(ciphertext string) (string, error) {
	block, err := aes.NewCipher([]byte(g.sessionKey))
	if err != nil {
		return "", err
	}

	if len(g.ivKey) != aes.BlockSize {
		return "", fmt.Errorf("IV length must be %d bytes", aes.BlockSize)
	}

	// Decode the base64-encoded ciphertext
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(decodedCiphertext))
	mode := cipher.NewCBCDecrypter(block, []byte(g.ivKey))
	mode.CryptBlocks(decrypted, decodedCiphertext)

	// Unpad the decrypted data
	unpaddedData, err := PKCS7Unpad(decrypted)
	if err != nil {
		return "", err
	}

	return string(unpaddedData), nil
}

// Checksum (X–Golomt–Checksum) үүсгэх
func (o openbank) bodyChecksum(body interface{}) (string, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(jsonBody)
	hex := hex.EncodeToString(hash[:])
	checkSum, err := o.EncryptAESCBC(hex)
	if err != nil {
		return "", err
	}
	return checkSum, nil
}

func parseEncryptedResponse[T any](response []byte, decryptFunc func(string) (string, error)) (T, error) {
	var result T
	responseData, err := decryptFunc(string(response))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(responseData), &result)
	return result, err
}

func parseResponse[T any](response []byte) (T, error) {
	var result T
	err := json.Unmarshal(response, &result)
	return result, err
}

// Харилцагч АМЖИЛТТАЙ нэвтэрсэн үед гуравдагч системийн тус сервисийн
// хариуг хүлээж авахаар өгөгдсөн Redirect URL авто дуудагдана.
// Тухайн хариу дээрх утгуудыг хүлээж авах функц
func ParseOathResponse(response []byte) (*model.OAuthResp, error) {
	var result *model.OAuthResp
	err := json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func bodyReader(body interface{}) *bytes.Reader {
	requestByte, _ := json.Marshal(body)
	if len(requestByte) == 0 {
		return nil
	}
	requestBody := bytes.NewReader(requestByte)
	return requestBody
}

// requestOption нь postEncrypted хүсэлтийн нэмэлт тохиргоо
type requestOption struct {
	// X-Golomt-Code (TOTP) header нэмэх эсэх
	withCode bool
	// Нэмэлт query параметрүүд (хуудаслалт г.м)
	query map[string]string
}

// postEncrypted нь client_id/state/scope query, checksum-тай POST хүсэлт илгээж,
// нууцлагдсан хариуг decrypt хийн T төрөлд хөрвүүлнэ.
func postEncrypted[T any](o *openbank, service, path string, body interface{}, opt requestOption) (T, error) {
	var result T
	if err := o.auth(); err != nil {
		return result, err
	}
	checksum, err := o.bodyChecksum(body)
	if err != nil {
		return result, err
	}
	client := resty.New()
	defer client.Close()
	req := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", service).
		SetHeader("X-Golomt-Checksum", checksum).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
		}).
		SetQueryParams(opt.query).
		SetBody(bodyReader(body))
	if opt.withCode {
		code, err := GenerateCurrentNumberString(o.xGolomtKey)
		if err != nil {
			return result, err
		}
		req.SetHeader("X-Golomt-Code", code)
	}
	res, err := req.Post(o.url + path)
	if err != nil {
		return result, err
	}
	response := res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return result, fmt.Errorf("%s-Golomt CG %s response: %s", time.Now().Format("20060102150405"), service, res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return result, err
		}
		return result, fmt.Errorf("%s-Golomt CG %s response: %s: %s", time.Now().Format("20060102150405"), service, errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[T](response, o.DecryptAESCBC)
}

// postPlain нь /v1/utility сервисүүдэд зориулсан (хариу нууцлагдаагүй) POST хүсэлт
func postPlain[T any](o *openbank, service, path string, body interface{}) (T, error) {
	var result T
	if err := o.auth(); err != nil {
		return result, err
	}
	checksum, err := o.bodyChecksum(body)
	if err != nil {
		return result, err
	}
	client := resty.New()
	defer client.Close()
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", service).
		SetHeader("X-Golomt-Checksum", checksum).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		Post(o.url + path)
	if err != nil {
		return result, err
	}
	response := res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return result, fmt.Errorf("%s-Golomt CG %s response: %s", time.Now().Format("20060102150405"), service, res.Status())
		}
		errResp, err := parseResponse[*model.ErrorResp](response)
		if err != nil {
			return result, err
		}
		return result, fmt.Errorf("%s-Golomt CG %s response: %s: %s", time.Now().Format("20060102150405"), service, errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[T](response)
}
