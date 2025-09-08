package mini_app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/techpartners-asia/golomt-api-go/utils"
)

type socialPayMiniApp struct {
	baseUrl         string // example: https://sp-api.golomtbank.com/api
	clientId        string // example: test_cert_id
	base64PublicKey string // example: MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJ7rnZH30unXZbTNHNX7wjfECxWyaABX88F5cjSqnA5Soo6Uwu72ufzjEzAtoPk8sE9tnfi
}

type SocialPayMiniApp interface {
	GetUserInfo(token string) (*UserInfo, error)
	SendNotification(input SendNotificationInput) error
}

// New creates a new SocialPayMiniApp instance
// baseUrl: example: https://sp-api.golomtbank.com/api
// clientId: example: test_cert_id
// base64PublicKey: example: MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJ7rnZH30unXZbTNHNX7wjfECxWyaABX88F5cjSqnA5Soo6Uwu72ufzjEzAtoPk8sE9tnfi
func New(baseUrl, clientId, base64PublicKey string) SocialPayMiniApp {
	return &socialPayMiniApp{
		baseUrl:         baseUrl,
		clientId:        clientId,
		base64PublicKey: base64PublicKey,
	}
}

// GetUserInfo gets user info from SocialPay MiniApp
// token: example: 1234567890
// return: UserInfo
func (s *socialPayMiniApp) GetUserInfo(token string) (*UserInfo, error) {
	jsonPayload := `{"token":"` + token + `"}` // store payload as string
	payload := strings.NewReader(jsonPayload)  // create reader from string
	encryptedToken, err := utils.EncryptRSA(jsonPayload, s.base64PublicKey)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, s.baseUrl+"/utility/miniapp/token/check?language=mn", payload)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("X-Golomt-Signature", encryptedToken)
	req.Header.Add("X-Golomt-Cert-Id", s.clientId)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := &UserInfo{}
	err = json.Unmarshal(body, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}

func (s *socialPayMiniApp) SendNotification(input SendNotificationInput) error {

	payload, err := json.Marshal(input)
	if err != nil {
		return err
	}

	stringPayload := string(payload)

	jsonPayload := strings.NewReader(stringPayload)

	encryptedToken, err := utils.EncryptRSA(stringPayload, s.base64PublicKey)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, s.baseUrl+"/utility/notification/push?language=mn", jsonPayload)
	if err != nil {
		return err
	}

	req.Header.Add("X-Golomt-Signature", encryptedToken)
	req.Header.Add("X-Golomt-Cert-Id", s.clientId)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	response := &SendNotificationResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}

	if response.Message != "Амжилттай" {
		return errors.New(response.Message)
	}

	return nil
}
