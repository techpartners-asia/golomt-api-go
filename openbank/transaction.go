package openbank

import (
	"fmt"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

// 6.1.	Голомт Банк хоорондын гүйлгээ
func (o openbank) TransactionInternal(body model.TransactionReq) (*model.TransactionResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "TXNADD").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", o.authObject.Token).
		SetBody(body).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetResult(&response).
		Post(o.url + "/v1/transaction/internal")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction internal response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[*model.TransactionResp](response, o.DecryptAESCBC)
}

// 6.2.	 Бусад банк хоорондын гүйлгээ
func (o openbank) TransactionOtherBank(body model.TransactionReq) (*model.TransactionResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	// TODO: check bank code
	if body.BankCode == "" {
		return nil, fmt.Errorf("bank code is required")
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "PMTADD").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", o.authObject.Token).
		SetBody(body).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetResult(&response).
		Post(o.url + "/v1/transaction/interbank")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction other bank response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[*model.TransactionResp](response, o.DecryptAESCBC)
}

// 6.3.	Гүйлгээний төлөв шалгах
func (o openbank) TransactionCheck(body model.TransactionCheckReq) (*model.TransactionCheckResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "TXNREF").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", o.authObject.Token).
		SetBody(body).
		SetResult(&response).
		Post(o.url + "/v1/transaction/confirm")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction status response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[*model.TransactionCheckResp](response, o.DecryptAESCBC)
}
