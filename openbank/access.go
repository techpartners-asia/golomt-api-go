package openbank

import (
	"fmt"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

func (o *openbank) auth() error {
	if o.authObject != nil {
		// 4.2.	Холболт шинэчлэх
		if o.expireTime.After(time.Now()) {
			return nil
		}
		client := resty.New()
		defer client.Close()
		var response *model.AuthResp
		var errResp *model.ErrorResp
		res, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("X-Golomt-Service", "LGIN").
			SetHeader("Authorization", "Bearer "+o.authObject.RefreshToken).
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
	request := model.AuthReq{
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
	// 4.1.	Нууц үгээр нэвтрэх
	client := resty.New()
	defer client.Close()
	var response *model.AuthResp
	var errResp *model.ErrorResp
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

// 4.4.	Сервисүүдийг багцаар авах
func (o *openbank) ServicesAccess(body model.ServiceListReq) (*model.ServiceListResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "BSRVACC").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(body).
		SetResult(&response).
		Post(o.url + "/v1/auth/services/access")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.ServiceListResp](response, o.DecryptAESCBC)
}

// 4.5.	Бүртгэлтэй дугаар татах
func (o *openbank) GetPhone() (*model.GetPhoneResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "PHINQ").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(nil)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetResult(&response).
		Get(o.url + "/v1/auth/authorize/getphone")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.GetPhoneResp](response, o.DecryptAESCBC)
}

// 4.6.	Бүртгэлтэй дугаар руу OTP код илгээх
func (o *openbank) OTPSend(phone string) (*model.OTPSendResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "OTPCD").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(nil)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetQueryParams(map[string]string{
			"phone": phone,
		}).
		SetResult(&response).
		Get(o.url + "/v1/auth/authorize/otpsend")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.OTPSendResp](response, o.DecryptAESCBC)
}

// 4.7.	OTP шалгах
func (o *openbank) OTPVerify(body model.OTPVerifyReq) (*model.OTPVerifyResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "OTPVAL").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(body).
		SetResult(&response).
		Post(o.url + "/v1/auth/authorize/otp")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.OTPVerifyResp](response, o.DecryptAESCBC)
}

// 4.8.	ХУР систем OTP илгээх
func (o *openbank) XypOTPSend() (*model.OTPSendResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "OTPXYPCD").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(nil)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetResult(&response).
		Get(o.url + "/v1/auth/xyp/otpsend")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.OTPSendResp](response, o.DecryptAESCBC)
}

// 4.9.	ХУР систем OTP шалгах
func (o *openbank) XypOTPVerify(body model.OTPXypVerifyReq) (*model.OTPVerifyResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "OTPXYPVAL").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(body)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetBody(body).
		SetResult(&response).
		Post(o.url + "/v1/auth/xyp/otp")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.OTPVerifyResp](response, o.DecryptAESCBC)
}

// 4.10. Тоон гарын үсгээр баталгаажуулах
func (o *openbank) DigitalSignature() (*model.DigitalSignatureResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}

	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "DSIGN").
		SetHeader("X-Golomt-Checksum", func() string {
			checksum, err := o.bodyChecksum(nil)
			if err != nil {
				return ""
			}
			return checksum
		}()).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"clientId": o.clientID,
			"state":    o.state,
			"scope":    o.scope,
		}).
		SetResult(&response).
		Get(o.url + "/v1/auth/authorize/signature")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG statement response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.DigitalSignatureResp](response, o.DecryptAESCBC)
}
