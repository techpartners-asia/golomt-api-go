package socialpay

import (
	"errors"
	"fmt"

	"github.com/techpartners-asia/golomt-api-go/utils"
	"resty.dev/v3"
)

type socialPay struct {
	terminal string
	secret   string
	endpoint string
}

type SocialPay interface {
	CreateInvoiceQR(input InvoiceInput) (*CommonResponse, error)         // Qr кодтой нэхэмжлэх үүсгэх
	CreateInvoicePhone(input InvoicePhoneInput) (*CommonResponse, error) // Утасны дугаараар нэхэмжлэх үүсгэх
	CancelInvoice(input InvoiceInput) (*CommonResponse, error)           // Нэхэмжлэх цуцлах
	CheckInvoice(input InvoiceInput) (*InvoiceResponse, error)           // Нэхэмжлэх шалгах
	CancelPayment(input InvoiceInput) (*InvoiceResponse, error)          // Төлбөр цуцлах
	Settlement(settlementId string) (*SettlementResponse, error)         // Гүйлгээний өндөрлөгөө шалгах
}

func New(terminal, secret, endpoint string) SocialPay {
	return socialPay{
		terminal: terminal,
		secret:   secret,
		endpoint: endpoint,
	}
}

func (s socialPay) CreateInvoiceQR(input InvoiceInput) (*CommonResponse, error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, input.Invoice, input.Amount))
	request := InvoiceRequest{
		Amount:   fmt.Sprintf("%.2f", input.Amount),
		Invoice:  input.Invoice,
		Terminal: s.terminal,
		Checksum: checksum,
	}
	client := resty.New()
	defer client.Close()
	var response *Response
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(s.endpoint + "/pos/invoice/qr")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.Header.Code != 200 {
		errorResponse := mapToErrorResponse(response.Body.Error)
		return nil, errors.New(errorResponse.ErrorDescription)
	}
	responseData := mapToCommonResponse(response.Body.Response)
	return &responseData, nil
}

func (s socialPay) CreateInvoicePhone(input InvoicePhoneInput) (*CommonResponse, error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, input.Invoice, input.Amount, input.Phone))
	request := InvoicePhoneRequest{
		Amount:   fmt.Sprintf("%.2f", input.Amount),
		Invoice:  input.Invoice,
		Phone:    input.Phone,
		Terminal: s.terminal,
		Checksum: checksum,
	}
	client := resty.New()
	defer client.Close()
	var response *Response
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(s.endpoint + "/pos/invoice/phone")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.Header.Code != 200 {
		errorResponse := mapToErrorResponse(response.Body.Error)
		return nil, errors.New(errorResponse.ErrorDescription)
	}
	responseData := mapToCommonResponse(response.Body.Response)
	return &responseData, nil
}

func (s socialPay) CancelInvoice(input InvoiceInput) (*CommonResponse, error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, input.Invoice, input.Amount))
	request := InvoiceRequest{
		Amount:   fmt.Sprintf("%.2f", input.Amount),
		Invoice:  input.Invoice,
		Terminal: s.terminal,
		Checksum: checksum,
	}
	client := resty.New()
	defer client.Close()
	var response *Response
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(s.endpoint + "/pos/invoice/cancel")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.Header.Code != 200 {
		errorResponse := mapToErrorResponse(response.Body.Error)
		return nil, errors.New(errorResponse.ErrorDescription)
	}
	responseData := mapToCommonResponse(response.Body.Response)
	return &responseData, nil
}

func (s socialPay) CheckInvoice(input InvoiceInput) (*InvoiceResponse, error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, input.Invoice, input.Amount))
	request := InvoiceRequest{
		Invoice:  input.Invoice,
		Amount:   fmt.Sprintf("%.2f", input.Amount),
		Terminal: s.terminal,
		Checksum: checksum,
	}
	client := resty.New()
	defer client.Close()
	var response *Response
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(s.endpoint + "/pos/invoice/check")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.Header.Code != 200 {
		errorResponse := mapToErrorResponse(response.Body.Error)
		return nil, errors.New(errorResponse.ErrorDescription)
	}
	responseData := mapToInvoiceResponse(response.Body.Response)
	return &responseData, nil
}

func (s socialPay) CancelPayment(input InvoiceInput) (*InvoiceResponse, error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, input.Invoice, input.Amount))
	request := InvoiceRequest{
		Invoice:  input.Invoice,
		Amount:   fmt.Sprintf("%.2f", input.Amount),
		Terminal: s.terminal,
		Checksum: checksum,
	}
	client := resty.New()
	defer client.Close()
	var response *Response
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(s.endpoint + "/pos/payment/cancel")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.Header.Code != 200 {
		errorResponse := mapToErrorResponse(response.Body.Error)
		return nil, errors.New(errorResponse.ErrorDescription)
	}
	responseData := mapToInvoiceResponse(response.Body.Response)
	return &responseData, nil
}

func (s socialPay) Settlement(settlementId string) (*SettlementResponse, error) {
	checksum := utils.GenerateHMAC(s.secret, utils.AppendAsString(s.terminal, settlementId))
	request := SettlementRequest{
		SettlementId: settlementId,
		Terminal:     s.terminal,
		Checksum:     checksum,
	}
	client := resty.New()
	defer client.Close()
	var response *Response
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(s.endpoint + "/pos/settlement")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.Header.Code != 200 {
		errorResponse := mapToErrorResponse(response.Body.Error)
		return nil, errors.New(errorResponse.ErrorDescription)
	}
	responseData := mapToSettlementResponse(response.Body.Response)
	return &responseData, nil
}
