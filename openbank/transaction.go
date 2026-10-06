package openbank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

// 8.1.	Голомт Банк хоорондын гүйлгээ
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
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
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

// 8.2.	Бусад банк хоорондын гүйлгээ
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
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
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

// 8.3. Байгууллага өөрийн дансаас гүйлгээ хийх
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
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
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

// 8.6. Гүйлгээ буцаах
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
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
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

// 8.7. Гүйлгээ шалгах
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
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
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

// 8.8. Багц гүйлгээ хийх
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
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
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

// 8.9. Багц гүйлгээний төлөв шалгах
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
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
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

// 8.10. Гүйлгээний төлөв шалгах
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

// 8.11. Багц гүйлгээ файлаар хийх
func (o *openbank) TransactionBatchFile(input model.TransactionBatchFileInput) (*model.TransactionBatchFileResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	fileContent, err := json.Marshal(input.File)
	if err != nil {
		return nil, err
	}
	fileName := "batch_transaction_" + time.Now().Format("20060102150405") + ".json"

	var response []byte
	res, err := client.R().
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
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
		}).
		SetMultipartFormData(map[string]string{
			"registerNumber": input.RegisterNo,
			"fileCode":       input.FileCode,
			"remarks":        input.Remarks,
		}).
		SetFileReader("file", fileName, bytes.NewReader(fileContent)).
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

// 8.4. Гаалийн гүйлгээ хийх
func (o *openbank) TransactionCustomsPay(body model.CustomsPayReq) (*model.CustomsPayResp, error) {
	return postEncrypted[*model.CustomsPayResp](o, "CUSPAY", "/v1/payment/custom/pay", body, requestOption{})
}

// 8.5. Татварын гүйлгээ хийх
func (o *openbank) TransactionTaxPay(body model.TaxPayReq) (*model.TaxPayResp, error) {
	return postEncrypted[*model.TaxPayResp](o, "TAXITR", "/v1/transaction/general/tax", body, requestOption{})
}

// 8.12. Татварын төлбөрийн жагсаалт харах TIN
func (o *openbank) TaxListByTIN(body model.TaxTINInqReq) ([]model.TaxTINInqData, error) {
	return postEncrypted[[]model.TaxTINInqData](o, "TAXTININQ", "/v1/payment/taxtin/inq", body, requestOption{})
}

// 8.13. Татварын төлбөрийн жагсаалт харах PIN
func (o *openbank) TaxListByPIN(body model.TaxPINInqReq) ([]model.TaxPINInqData, error) {
	return postEncrypted[[]model.TaxPINInqData](o, "TAXPININQ", "/v1/payment/taxpin/inq", body, requestOption{})
}

// 8.14. Татварын төлбөрийн нэхэмжлэх / цахим төлбөрийн даалгаврын дугаараар лавлагаа авах
func (o *openbank) TaxInvoiceInq(body model.TaxInvoiceInqReq) (*model.TaxInvoiceInqResp, error) {
	return postEncrypted[*model.TaxInvoiceInqResp](o, "TAXINVINQ", "/v1/payment/taxinv/inq", body, requestOption{})
}

// 8.15. Гаалийн төлбөрийн нэхэмжлэх / цахим төлбөрийн даалгаврын дугаараар лавлагаа авах
func (o *openbank) CustomsInvoiceInq(body model.CustomsInvoiceInqReq) (*model.CustomsInvoiceInqResp, error) {
	return postEncrypted[*model.CustomsInvoiceInqResp](o, "CGAINQ", "/v1/payment/cgainv/inq", body, requestOption{})
}

// 8.16. Файлаар хийсэн багц гүйлгээний дэлгэрэнгүй татах
func (o *openbank) TransactionBatchFileInq(body model.TransactionBatchFileInqReq) (*model.TransactionBatchFileInqResp, error) {
	return postEncrypted[*model.TransactionBatchFileInqResp](o, "CGWTTUMINQ", "/v1/transaction/cgw/ttum/inq", body, requestOption{})
}
