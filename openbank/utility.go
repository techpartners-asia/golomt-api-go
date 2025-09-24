package openbank

import (
	"fmt"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

// 7.1.	Хот, аймагийн жагсаалт авах
func (o *openbank) StateListInq(body model.StateListReq) (*model.StateListResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "STATEINQ").
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
		Post(o.url + "/v1/utility/state/inq")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG utility area list inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[*model.StateListResp](response, o.DecryptAESCBC)
}

// 7.2.	Сум, дүүргийн жагсаалт авах
func (o *openbank) DistrictListInq(body model.DistrictListReq) (*model.DistrictListResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CITYINQ").
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
		Post(o.url + "/v1/utility/city/inq")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG utility district list inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[*model.DistrictListResp](response, o.DecryptAESCBC)
}

// 7.3.	Категори төрлөөр сонголтын жагсаалт авах
func (o *openbank) CategoryListInq(body model.CategoryReq) (*model.CategoryResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "CATINQ").
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
		Post(o.url + "/v1/utility/category/inq")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG utility category list inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[*model.CategoryResp](response, o.DecryptAESCBC)
}

// 7.4.	Ханшны мэдээлэл авах
func (o *openbank) RateInq(body model.RateReq) (*model.RateResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
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
		errResp, err := parseResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG utility rate inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[*model.RateResp](response, o.DecryptAESCBC)
}

// 7.5.	Салбарын жагсаалт авах
func (o *openbank) BranchListInq(body model.BranchListReq) (*model.BranchListResp, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "SOLINQ").
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
		Post(o.url + "/v1/utility/sol/inq")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		errResp, err := parseResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG utility branch list inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[*model.BranchListResp](response, o.DecryptAESCBC)
}

// 7.6.	Бүтээгдэхүүн лавлах
func (o *openbank) ProductListInq(body model.ProductListReq) (*[]model.ProductData, error) {
	if err := o.auth(); err != nil {
		return nil, err
	}
	client := resty.New()
	defer client.Close()
	var response []byte
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Golomt-Service", "PROCATINQ").
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
		Post(o.url + "/v1/utility/product/category/inq")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG utility product list inq response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG utility product list inq response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	return parseResponse[*[]model.ProductData](response, o.DecryptAESCBC)
}
