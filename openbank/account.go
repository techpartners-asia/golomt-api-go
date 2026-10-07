package openbank

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

// 5.1.	Дансны үлдэгдэл
//
// SPEC 5 (алхам 2-7) ба 4.3: эхний хүсэлт client_id/state/scope хоосон явна,
// банк OAuth хариу (clientId, state, scope) буцаана; тэдгээрийг ашиглаж
// хүсэлтийг дахин илгээхэд үлдэгдэл ирнэ. Хариунд url ирвэл харилцагч
// банкны хуудсаар баталгаажуулах шаардлагатай тул алдаа буцаана.
func (o *openbank) AccountBalcInq(body model.AccountBalcInqReq) (*model.AccountBalcInqResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	plain, err := o.postAccountBalance(body)
	if err != nil {
		return nil, err
	}
	if grant, ok := oauthGrant(plain); ok {
		if grant.State == "" || grant.Scope == "" {
			if grant.Url != "" {
				return nil, fmt.Errorf("%s-Golomt CG account balance inq: харилцагчийн баталгаажуулалт шаардлагатай: %s", time.Now().Format("20060102150405"), grant.Url)
			}
			return nil, fmt.Errorf("%s-Golomt CG account balance inq: OAuth хариунд state/scope алга: %s", time.Now().Format("20060102150405"), truncate(plain, 300))
		}
		o.SetOAuthResponse(*grant)
		if plain, err = o.postAccountBalance(body); err != nil {
			return nil, err
		}
		if _, again := oauthGrant(plain); again {
			return nil, fmt.Errorf("%s-Golomt CG account balance inq: state/scope хүлээн авагдсангүй: %s", time.Now().Format("20060102150405"), truncate(plain, 300))
		}
	}
	var result *model.AccountBalcInqResp
	if err := json.Unmarshal(plain, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// postAccountBalance sends one ACCTBALINQ request and returns the decrypted
// JSON body. Until a grant is held, client_id/state/scope go out empty, as
// in the bank's Postman and Go examples.
func (o *openbank) postAccountBalance(body model.AccountBalcInqReq) ([]byte, error) {
	clientID := o.clientID
	if o.state == "" {
		clientID = ""
	}
	checksum, err := o.bodyChecksum(body)
	if err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "ACCTBALINQ").
		SetHeader("X-Golomt-Checksum", checksum).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		SetQueryParams(map[string]string{
			"client_id": clientID,
			"state":     o.state,
			"scope":     o.scope,
		}).
		Post(o.url + "/v1/account/balance/inq")
	if err != nil {
		return nil, err
	}
	response := res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil || errResp == nil {
			return nil, fmt.Errorf("%s-Golomt CG account balance inq response: http %d: %s", time.Now().Format("20060102150405"), res.StatusCode(), truncate(response, 300))
		}
		return nil, fmt.Errorf("%s-Golomt CG account balance inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return decodeBody(response, o.DecryptAESCBC)
}

// 5.2.	Дансны төрөл шалгах
func (o *openbank) AccountTypeInq(body model.AccountTypeInqReq) (*model.AccountTypeInqResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "ACCTTYPEINQ").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/account/type/inq")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account type inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountTypeInqResp](response, o.DecryptAESCBC)
}

// 5.3.	Дансны товч нэр солих
func (o *openbank) AccountRename(body model.AccountRenameReq) (*model.AccountRenameResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "ACCTRNM").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/account/rename")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account rename response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountRenameResp](response, o.DecryptAESCBC)
}

// 5.4.	Харилцах дансны дэлгэрэнгүй мэдээлэл харах
func (o *openbank) AccountDetail(body model.AccountDetailReq) (*model.AccountDetailResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "OPERACCTDET").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
		}).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/account/operative/details")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account detail response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountDetailResp](response, o.DecryptAESCBC)
}

// 5.5.	Харилцах дансны хуулга харах
func (o *openbank) AccountStatement(body model.StatementReq) (*model.StatementResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "OPERACCTSTA").
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
		Post(o.url + "/v1/account/operative/statement/")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG statement response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.StatementResp](response, o.DecryptAESCBC)
}

// 5.6.	Харилцах дансны хуулга хуудаслалтай харах
func (o *openbank) AccountStatementPage(body model.StatementPageReq) (*model.StatementPageResp, error) {
	return postEncrypted[*model.StatementPageResp](o, "OPERACCSTAINQ", "/v1/account/operative/statement/inquiry", body, requestOption{})
}

// 5.7.	Харилцах данс нээх
func (o *openbank) AccountAdd(body model.AccountAddReq) (*model.AccountAddResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "OPERACCADD").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
		}).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/account/operative/add")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account add response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountAddResp](response, o.DecryptAESCBC)
}

// 5.8.	Хадгаламжийн дансны дэлгэрэнгүй
func (o *openbank) AccountDepositDetail(body model.AccountDetailReq) (*model.AccountDepositDetailResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "DEPACCDTLS").
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
		Post(o.url + "/v1/account/deposit/details")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account deposit detail response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountDepositDetailResp](response, o.DecryptAESCBC)
}

// 5.9.	Хадгаламжийн дансны хуулга харах
// SPEC: хариу нь `account` талбартай гэж тодорхойлогдсон
func (o *openbank) AccountDepositStatement(body model.StatementReq) (*model.AccountAddResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "DEPACCSTTM").
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
		Post(o.url + "/v1/account/deposit/statement")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account deposit statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountAddResp](response, o.DecryptAESCBC)
}

// 5.10. Хадгаламжийн данс нээх
func (o *openbank) AccountDepositAdd(body model.AccountDepositAddReq) (*model.AccountAddResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "DEPACCADD").
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
		Post(o.url + "/v1/account/deposit/add")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account deposit add response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountAddResp](response, o.DecryptAESCBC)
}

// 5.11. Дансны жагсаалт татах
func (o *openbank) AccountList(body model.AccountListReq) (*model.AccountListResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "ACCTLST").
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
		Post(o.url + "/v1/account/list")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG account list response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account list response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountListResp](response, o.DecryptAESCBC)
}

// 5.12.a Данс эзэмшигчнийн мэдээллэл авах /Голомт/
func (o *openbank) AccountCustomerDetail(body model.AccountCustomerDetailReq) (*model.AccountCustomerDetailResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "ACCCHK").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
		}).
		SetBody(bodyReader(body)).
		Post(o.url + "/v1/account/check/account")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account customer detail response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountCustomerDetailResp](response, o.DecryptAESCBC)
}

// 5.12.b Данс эзэмшигчнийн мэдээллэл авах /Голомт бус/
func (o *openbank) AccountOtherBankCustomerDetail(body model.AccountCustomerDetailReq) (*model.AccountOtherBankCustomerDetailResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "ACCCHK").
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
		Post(o.url + "/v1/account/check/account")
	if err != nil {
		return nil, err
	}
	response = res.Bytes()
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account customer detail response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.AccountOtherBankCustomerDetailResp](response, o.DecryptAESCBC)
}
