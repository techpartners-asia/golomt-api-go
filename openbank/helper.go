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

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
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
		fmt.Println("body is empty")
		return nil
	}
	requestBody := bytes.NewReader(requestByte)
	return requestBody
}
