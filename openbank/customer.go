package openbank

import (
	"bytes"
	"fmt"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

// 6.1. Харилцагч шалгах
func (o *openbank) CustomerCheck(body model.CustomerCheckReq) (*model.CustomerCheckResp, error) {
	return postEncrypted[*model.CustomerCheckResp](o, "CUSTCHECK", "/v1/customer/check", body, requestOption{})
}

// 6.2. Харилцагчийн утасны дугаар мөн эсхийг шалгах
func (o *openbank) CustomerPhoneCheck(body model.CustomerPhoneCheckReq) (*model.CustomerPhoneCheckResp, error) {
	return postEncrypted[*model.CustomerPhoneCheckResp](o, "CUSTCHPH", "/v1/customer/check/phone", body, requestOption{})
}

// 6.3. Харилцагчийн мэдээлэл харах
func (o *openbank) CustomerInquire(body model.CustomerInquireReq) (*model.CustomerInquireResp, error) {
	return postEncrypted[*model.CustomerInquireResp](o, "RETCUSTINQ", "/v1/customer/inquire", body, requestOption{})
}

// 6.4. Иргэн – Харилцагч бүртгэх (харилцагчийн зураг хуулах)
func (o *openbank) CustomerImageUpload(body model.CustomerImageUploadReq) (*model.CustomerImageUploadResp, error) {
	return postEncrypted[*model.CustomerImageUploadResp](o, "CUSTIMGUP", "/v1/customer/image/upload", body, requestOption{})
}

// 6.4. Иргэн – Харилцагч бүртгэх
func (o *openbank) CustomerRetailAdd(body model.CustomerRetailAddReq) (*model.CustomerAddResp, error) {
	return postEncrypted[*model.CustomerAddResp](o, "RETCUSTADD", "/v1/customer/retail/add", body, requestOption{})
}

// 6.5. Байгууллага - Харилцагч бүртгэх
func (o *openbank) CustomerCorpAdd(body model.CustomerCorpAddReq) (*model.CustomerAddResp, error) {
	return postEncrypted[*model.CustomerAddResp](o, "CORPCUSTADD", "/v1/customer/corporate/add/v1.0", body, requestOption{})
}

// 6.6. Харилцагчийн мэдээлэл бүртгэх
func (o *openbank) CustomerSaveInfo(body model.CustomerSaveInfoReq) (*model.CustomerSaveInfoResp, error) {
	return postEncrypted[*model.CustomerSaveInfoResp](o, "CUSTSVINF", "/v1/customer/save/information", body, requestOption{})
}

// 6.7. Бүртгэгдсэн харилцагчийн мэдээлэл татах
func (o *openbank) CustomerGetInfo(body model.CustomerGetInfoReq) (*model.CustomerGetInfoResp, error) {
	return postEncrypted[*model.CustomerGetInfoResp](o, "CUSTGTINF", "/v1/customer/get/information", body, requestOption{})
}

// 6.8. Харилцагчийн иргэний үнэмлэхний зураг upload хийх /BASE64/
func (o *openbank) CustomerNewImageUpload(body model.CustomerNewImageUploadReq) (*model.CustomerNewImageUploadResp, error) {
	return postEncrypted[*model.CustomerNewImageUploadResp](o, "CUSTNWIMGUP", "/v1/customer/new/image/upload", body, requestOption{})
}

// 6.9. Харилцагчийн иргэний үнэмлэхний зураг upload хийх /FORM-DATA/
func (o *openbank) CustomerNewImageFormUpload(input model.CustomerNewImageFormUploadInput) (*model.CustomerNewImageUploadResp, error) {
	const service = "CUSTNWIMGUPFRMDT"
	if err := o.auth(); err != nil {
		return nil, err
	}
	fileName := input.FileName
	if fileName == "" {
		fileName = input.ImgName
	}
	client := resty.New()
	defer client.Close()
	res, err := client.R().
		SetHeader("X-Golomt-Service", service).
		SetHeader("Authorization", "Bearer "+o.authObject.Token).
		SetQueryParams(map[string]string{
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
		}).
		SetMultipartFormData(map[string]string{
			"infoId":  input.InfoID,
			"imgCat":  input.ImgCat,
			"imgName": input.ImgName,
		}).
		SetFileReader("imgFile", fileName, bytes.NewReader(input.ImgFile)).
		Post(o.url + "/v1/customer/new/image/form-upload")
	if err != nil {
		return nil, err
	}
	response := res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return nil, fmt.Errorf("%s-Golomt CG %s response: %s", time.Now().Format("20060102150405"), service, res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s-Golomt CG %s response: %s: %s", time.Now().Format("20060102150405"), service, errResp.Message, errResp.DebugMessage)
	}
	return parseEncryptedResponse[*model.CustomerNewImageUploadResp](response, o.DecryptAESCBC)
}

// 6.10. Харилцагчийн иргэний үнэмлэхний зураг татах
func (o *openbank) CustomerImageDownload(body model.CustomerImageDownloadReq) (*model.CustomerImageDownloadResp, error) {
	return postEncrypted[*model.CustomerImageDownloadResp](o, "CUSTIMDL", "/v1/customer/image/download", body, requestOption{})
}

// 6.11. Байгууллага - Харилцагчийн мэдээлэл татах
func (o *openbank) CustomerCorpDetail(body model.CustomerCorpDetailReq) (*model.CustomerCorpDetailResp, error) {
	return postEncrypted[*model.CustomerCorpDetailResp](o, "CORPCUSTDTL", "/v1/customer/corporate/detail", body, requestOption{})
}

// 6.12. Байгууллага - Харилцагчийн limit-н мэдээлэл татах
func (o *openbank) CustomerCorpLimitDetail(body model.CustomerCorpLimitDetailReq) (*model.CustomerCorpLimitDetailResp, error) {
	return postEncrypted[*model.CustomerCorpLimitDetailResp](o, "CUSTLMTDTL", "/v1/customer/corporate/customer/limit/detail", body, requestOption{})
}
