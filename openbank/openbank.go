package openbank

import (
	"fmt"
	"time"

	"resty.dev/v3"
)

type openbank struct {
	organizationName string
	username         string
	password         string
	ivKey            string
	sessionKey       string
	url              string
	registerNo       string
	expireTime       time.Time
	authObject       *AuthResp
	clientID         string
	state            string
	scope            string
}

type Openbank interface {
	Statement(body StatementReq) (*StatementResp, error)                // Харилцах дансны хуулга харах
	AccountList(body AccountListReq) (*AccountListResp, error)          // Дансны жагсаалт татах
	AccountTypeInq(body AccountTypeInqReq) (*AccountTypeInqResp, error) // Дансны төрөл шалгах
	AccountBalcInq(body AccountBalcInqReq) (*AccountBalcInqResp, error) // Дансны үдэгдэл шалгах
	UtilityRateInq() (*RateResp, error)                                 // Ханшны мэдээлэл авах
}

func New(input BaseInput) Openbank {
	return &openbank{
		organizationName: input.OrganizationName,
		username:         input.Username,
		password:         input.Password,
		ivKey:            input.IvKey,
		sessionKey:       input.SessionKey,
		url:              input.Url,
		authObject:       nil,
		expireTime:       time.Time{},
		registerNo:       input.RegisterNo,
		clientID:         input.ClientID,
		state:            "",
		scope:            "",
	}
}

func (o openbank) auth() error {
	if o.authObject != nil {
		if o.expireTime.Before(time.Now()) {
			return nil
		}
		client := resty.New()
		defer client.Close()
		var response *AuthResp
		var errResp *ErrorResp
		res, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("X-Golomt-Service", "LGIN").
			SetResult(&response).
			SetError(&errResp).
			Get(o.url + "/v1/auth/refresh")
		if err != nil {
			return err
		}
		if res.IsError() {
			return fmt.Errorf("%s-Golomt CG auth response: %s", time.Now().Format("20060102150405"), errResp.Message)
		}
		o.authObject = response
		o.expireTime = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)
		return nil
	}
	var aesError error
	request := AuthReq{
		Name: o.username,
		Password: func() string {
			pass, err := o.EncryptAESCBC(o.password)
			if err != nil {
				aesError = err
				return ""
			}
			return pass
		}(),
	}
	if aesError != nil {
		return fmt.Errorf("%s-Golomt CG auth response: %s", time.Now().Format("20060102150405"), aesError.Error())
	}
	client := resty.New()
	defer client.Close()
	var response *AuthResp
	var errResp *ErrorResp
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "LGIN").
		SetBody(request).
		SetResult(&response).
		SetError(&errResp).
		Post(o.url + "/v1/auth/login")
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("%s-Golomt CG auth response: %s", time.Now().Format("20060102150405"), errResp.Message)
	}
	o.authObject = response
	o.expireTime = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)
	return nil
}

func (o openbank) Statement(body StatementReq) (*StatementResp, error) {
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
		SetHeader("Authorization", o.authObject.Token).
		SetBody(body).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetResult(&response).
		Post(o.url + "/v1/account/operative/statement/")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := convertResponse[*ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return convertResponse[*StatementResp](response, o.DecryptAESCBC)
}

func (o openbank) AccountList(body AccountListReq) (*AccountListResp, error) {
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
		SetHeader("Authorization", o.authObject.Token).
		SetBody(body).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetResult(&response).
		Post(o.url + "/v1/account/list")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := convertResponse[*ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account list response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return convertResponse[*AccountListResp](response, o.DecryptAESCBC)
}

func (o openbank) AccountTypeInq(body AccountTypeInqReq) (*AccountTypeInqResp, error) {
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
		SetHeader("Authorization", o.authObject.Token).
		SetBody(body).
		SetResult(&response).
		Post(o.url + "/v1/account/type/inq")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := convertResponse[*ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account type inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return convertResponse[*AccountTypeInqResp](response, o.DecryptAESCBC)
}

func (o openbank) AccountBalcInq(body AccountBalcInqReq) (*AccountBalcInqResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "ACCTBALINQ").
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
		Post(o.url + "/v1/account/balance/inq")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := convertResponse[*ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG account balance inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return convertResponse[*AccountBalcInqResp](response, o.DecryptAESCBC)
}

func (o openbank) UtilityRateInq() (*RateResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	body := RateReq{
		Currency: "MNT",
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "RATEINQ").
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
		Post(o.url + "/v1/utility/rate/inq")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := convertResponse[*ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG utility rate inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return convertResponse[*RateResp](response, o.DecryptAESCBC)
}
