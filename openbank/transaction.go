package openbank

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

// 6.1.	Голомт Банк хоорондын гүйлгээ
func (o *openbank) TransactionInBank(body model.TransactionReq) (*model.TransactionResp, error) {
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
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		Post(o.url + "/v1/transaction/internal")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG transaction internal response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction internal response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TransactionResp](response, o.DecryptAESCBC)
}

// 6.2.	 Бусад банк хоорондын гүйлгээ
func (o *openbank) TransactionOtherBank(body model.TransactionReq) (*model.TransactionResp, error) {
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
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		Post(o.url + "/v1/transaction/interbank")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG transaction other bank response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction other bank response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TransactionResp](response, o.DecryptAESCBC)
}

// 6.3. Байгууллага өөрийн дансаас гүйлгээ хийх
func (o *openbank) TransactionSelf(body model.TransactionSelfReq) (*model.TransactionSelfResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CGWTXNADD").
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
		Post(o.url + "/v1/transaction/cgw/transfer")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG transaction self response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction self response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TransactionSelfResp](response, o.DecryptAESCBC)
}

// 6.4. Гүйлгээ буцаах
func (o *openbank) TransactionRefund(body model.TransactionRefundReq) (*model.TransactionRefundResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "TXNREV").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/transaction/ref/rev")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG transaction refund response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction refund response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TransactionRefundResp](response, o.DecryptAESCBC)
}

// 6.5. Гүйлгээ шалгах
func (o *openbank) TransactionCheck(body model.TransactionCheckReq) (*model.TransactionCheckResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "TXNCHK").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/transaction/ref/check")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG transaction check response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction check response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TransactionCheckResp](response, o.DecryptAESCBC)
}

// 6.6. Багц гүйлгээ хийх
func (o *openbank) TransactionBatch(body model.TransactionBatchReq) (*model.TransactionBatchResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()

	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CGWBLKTXN").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetHeader("X-Golomt-Code", func() string {
			code, err := GenerateCurrentNumberString(o.xGolomtKey)
			if err != nil {
				return ""
			}
			return code
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/transaction/cgw/bulk")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG transaction batch response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction batch response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TransactionBatchResp](response, o.DecryptAESCBC)
}

// 6.7. Багц гүйлгээний төлөв шалгах
func (o *openbank) TransactionBatchCheck(body model.TransactionBatchCheckReq, page model.PageReq) (*model.TransactionBatchCheckResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()

	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CGWBLKINQ").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetQueryParams(map[string]string{
			"page_no":   page.PageNo,
			"page_size": page.PageSize,
			"sort":      page.Sort,
			"sort_by":   page.SortBy,
		}).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/transaction/cgw/bulk/inq")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG transaction batch check response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction batch check response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TransactionBatchCheckResp](response, o.DecryptAESCBC)
}

// 6.8. Гүйлгээний төлөв шалгах
func (o *openbank) TransactionConfirm(body model.TransactionConfirmReq) (*model.TransactionConfirmResp, error) {
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
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/transaction/confirm")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG transaction status response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction status response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TransactionConfirmResp](response, o.DecryptAESCBC)
}

// 6.9. Багц гүйлгээ файлаар хийх
func (o *openbank) TransactionBatchFile(input model.TransactionBatchFileInput) (*model.TransactionBatchFileResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	fileName := "batch_transaction_" + time.Now().Format("20060102150405") + ".json"
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Encode with indentation for readability

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "")
	if err := encoder.Encode(input.File); err != nil {
		panic(err)
	}
	file.Close()

	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetHeader("X-Golomt-Service", "CGWTTUM").
		SetHeader("X-Golomt-Code", func() string {
			code, err := GenerateCurrentNumberString(o.xGolomtKey)
			if err != nil {
				return ""
			}
			return code
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetFormData(map[string]string{
			"registerNumber": input.RegisterNo,
			"fileCode":       input.FileCode,
			"remarks":        input.Remarks,
			"file":           fileName,
		}).
		Post(o.url + "/v1/transaction/cgw/ttum")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG transaction batch file response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG transaction batch file response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.TransactionBatchFileResp](response, o.DecryptAESCBC)
}
