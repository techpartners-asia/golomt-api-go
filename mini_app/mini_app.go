package mini_app

import (
	"encoding/json"
	"fmt"

	"resty.dev/v3"
)

type socialPayMiniApp struct {
	baseUrl         string // example: https://sp-api.golomtbank.com/api
	clientId        string // example: test_cert_id
	base64PublicKey string // example: MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDJ7rnZH30unXZbTNHNX7wjfECxWyaABX88F5cjSqnA5Soo6Uwu72ufzjEzAtoPk8sE9tnfi
}

type SocialPayMiniApp interface {
	GetUserInfo(token string) (*UserInfo, error)
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
	client := resty.New()
	defer client.Close()
	var response *UserInfo
	body := map[string]interface{}{
		"token": token,
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	encryptedToken, err := generateGolomtSignature(string(bodyJson), s.base64PublicKey)
	if err != nil {
		return nil, err
	}
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Signature", encryptedToken).
		SetHeader("X-Golomt-Client-Id", s.clientId).
		SetBody(body).
		SetResult(&response). // or SetResult(LoginResponse{}).
		Post(s.baseUrl + "/utility/miniapp/token/check?language=mn")
	if err != nil {
		return nil, err
	}
	fmt.Println(res.String())
	if res.IsError() {
		return nil, fmt.Errorf("error: %s", res.String())
	}
	return response, nil
}
