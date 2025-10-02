package openbank

import (
	"fmt"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

// 8.1.	Байгууллагын виртуал кредит карт токенжуулах
func (o *openbank) CardTokenize(body model.TokenizeReq) (*model.TokenizeResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CRCORPTOK").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		Post(o.url + "/v1/card/corp/tokenize")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card corporate tokenize response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG card corporate tokenize response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TokenizeResp](response, o.DecryptAESCBC)
}

// 8.2.	Токен цуцлах
func (o *openbank) CardTokenClose(body model.TokenCloseReq) (*model.TokenCloseResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CRTOKCL").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		Post(o.url + "/v1/card/token/close")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card corporate close response: %s", time.Now().Format("20060102150405"), res.Status())
		}
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card corporate close response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG card corporate close response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TokenCloseResp](response, o.DecryptAESCBC)
}

// 8.3.	Токентэй картнаас гүйлгээ гаргах
func (o *openbank) CardPurchase(body model.CardPurchaseReq) (*model.CardPurchaseResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CDTXN").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		Post(o.url + "/v1/card/purchase")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card purchase response: %s", time.Now().Format("20060102150405"), res.Status())
		}
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card purchase response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG card purchase response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.CardPurchaseResp](response, o.DecryptAESCBC)
}

// 8.4.	Картын гүйлгээ шалгах
func (o *openbank) CardPurchaseCheck(body model.CardPurchaseCheckReq) (*model.CardPurchaseCheckResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CRDTXNINQ").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		Post(o.url + "/v1/card/tran/inq")
	if err != nil {
		return nil, err
	}

	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card purchase check response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG card purchase check response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.CardPurchaseCheckResp](response, o.DecryptAESCBC)
}

// 8.5.	Мерчантын хуулга авах

func (o *openbank) CardMerchantStatement(body model.CardMerchantStatementReq, page model.PageReq) (*model.CardMerchantStatementResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "MRCHSTMT").
		SetHeader("X-Golomt-Code", func() string {
			code, err := GenerateCurrentNumberString(o.xGolomtKey)
			if err != nil {
				return ""
			}
			return code
		}()).
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetQueryParams(map[string]string{
			"page_no":   page.PageNo,
			"page_size": page.PageSize,
		}).
		Post(o.url + "/v1/card/merchant/statement")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card merchant statement response: %s", time.Now().Format("20060102150405"), res.Status())
		}
	}
	return parseEncryptedResponse[*model.CardMerchantStatementResp](response, o.DecryptAESCBC)
}

// 8.6.	Кредит картын дэлгэрэнгүй
func (o *openbank) CardCreditDetail(body model.CardCreditDetailReq) (*model.CardCreditDetailResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CCDTLS").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		Post(o.url + "/v1/card/credit/details")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card credit detail response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG card credit detail response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.CardCreditDetailResp](response, o.DecryptAESCBC)
}

// 8.7.	Картын гүйлгээний мэдээлэл татах
func (o *openbank) CardTransaction(body model.CardTransactionReq) (*model.CardTransactionResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CRDTRNDETS").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		Post(o.url + "/v1/card/transaction-details")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card transaction response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG card transaction response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.CardTransactionResp](response, o.DecryptAESCBC)
}

// 8.8.	Кредит карт хуулга харах
func (o *openbank) CardCreditStatement(body model.CardCreditStatementReq) ([]model.CardCreditStatementData, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CCSTATM").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		Post(o.url + "/v1/card/credit/statement")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG card credit statement response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG card credit statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[[]model.CardCreditStatementData](response, o.DecryptAESCBC)
}
