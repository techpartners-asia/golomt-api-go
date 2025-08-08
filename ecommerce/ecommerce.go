package ecommerce

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/techpartners-asia/golomt-api-go/utils"
	"resty.dev/v3"
)

type golomtEcommerce struct {
	baseUrl     string // example: https://ecommerce.golomtbank.com/
	secret      string //
	bearerToken string //
	merchantId  string //
}

type GolomtEcommerce interface {
	CreateInvoice(input CreateInvoiceInput) (*CreateInvoiceResponse, error)                      // Нэхэмжлэх үүсгэх
	CheckTokenPayment(transactionId string) (*CheckTokenPaymentResponse, error)                  // Токентэй гүйлгээ шалгах
	PayTokenPayment(input PayTokenInput) (*PayTokenPaymentResponse, error)                       // Токентэй гүйлгээ хийх
	Inquiry(transactionId string) (*InquiryResponse, error)                                      // Нэхэмжлэх шалгах
	GetInvoiceUrl(input GetInvoiceInput) string                                                  // Нэхэмжлэхийн дугаарыг /Нэхэмжлэх үүсгэсэн үед/ ашиглан харилцагчийн browser дээр карт бөглөх хуудсыг дуудах
	ParsePushNotificationResponse(body []byte) (*InquiryResponse, error)                         // Push notification хүсэлт хувиргах
	GetConfirmationUrl(input GetConfirmationInput) string                                        // Нэмэмжлэхийн дугаарыг /Токен үүсгэх үед/ ашиглан харилцагчийн browser дээр карт бөглөх хуудсыг дуудах
	CreateToken(input CreateTokenInput) (*CreateTokenResponse, error)                            // Токен үүсгэх
	CheckToken(transactionId string) (*CheckTokenResponse, error)                                // Токен шалгах
	GetSettlementDetails(input GetSettlementDetailsInput) (*GetSettlementDetailsResponse, error) // Гүйлгээний өндөрлөгөө хийсэн дэлгэрэнгүй мэдээлэл авах
}

func New(baseUrl, secret, bearerToken string) GolomtEcommerce {
	return &golomtEcommerce{
		baseUrl:     baseUrl,
		secret:      secret,
		bearerToken: bearerToken,
	}
}

// GetInvoiceUrl - gets the URL by invoice ID
//
// input: GetInvoiceInput
// return: URL
func (g *golomtEcommerce) GetInvoiceUrl(input GetInvoiceInput) string {
	return fmt.Sprintf("https://ecommerce.golomtbank.com/%v/%v/%v", input.PaymentMethod, input.Lang, input.Invoice)
}

// PayByToken - pays by token
//
// input: ByTokenInput
// return: ByTokenResponse, error
func (g *golomtEcommerce) PayTokenPayment(input PayTokenInput) (*PayTokenPaymentResponse, error) {
	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(input.Amount, input.TransactionID, input.Token))
	request := PayTokenPaymentRequest{
		Amount:        fmt.Sprintf("%.2f", input.Amount),
		Checksum:      checksum,
		TransactionID: input.TransactionID,
		Token:         input.Token,
		Lang:          string(input.Lang),
	}

	client := resty.New()
	defer client.Close()
	var response *PayTokenPaymentResponse
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+g.bearerToken).
		SetBody(request).
		SetResult(&response).
		Post(g.baseUrl + "/api/pay")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.ErrorCode != "000" {
		return nil, fmt.Errorf("%s", response.ErrorDesc)
	}
	// Verify checksum
	expectedChecksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(
		response.TransactionID,
		response.ErrorCode,
		response.Amount,
	))
	if expectedChecksum != response.Checksum {
		return nil, fmt.Errorf("checksum verification failed")
	}
	return response, nil
}

// CheckTokenPayment - checks if the token payment is successful
//
// transactionId: Transaction ID
// return: CheckTokenPaymentResponse, error
func (g *golomtEcommerce) CheckTokenPayment(transactionId string) (*CheckTokenPaymentResponse, error) {
	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(transactionId, transactionId))
	request := CheckTokenPaymentRequest{
		Checksum:      checksum,
		TransactionID: transactionId,
	}

	client := resty.New()
	defer client.Close()
	var response *CheckTokenPaymentResponse
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+g.bearerToken).
		SetBody(request).
		SetResult(&response).
		Post(g.baseUrl + "/api/pay")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.ErrorCode != "000" {
		return nil, fmt.Errorf("%s", response.ErrorDesc)
	}
	// Verify checksum
	var expectedChecksum string
	if response.Token != "" {
		expectedChecksum = utils.GenerateHMAC(g.secret, utils.AppendAsString(
			response.TransactionID,
			response.ErrorCode,
			response.Amount,
			response.Token,
		))
	} else {
		expectedChecksum = utils.GenerateHMAC(g.secret, utils.AppendAsString(
			response.TransactionID,
			response.ErrorCode,
			response.Amount,
		))
	}
	if expectedChecksum != response.Checksum {
		return nil, fmt.Errorf("checksum verification failed")
	}
	return response, nil
}

// CreateInvoice - creates an invoice
//
// input: CreateInvoiceInput
// return: CreateInvoiceResponse, error
func (g *golomtEcommerce) CreateInvoice(input CreateInvoiceInput) (*CreateInvoiceResponse, error) {
	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(
		input.TransactionID,
		fmt.Sprintf("%.2f", input.Amount),
		input.ReturnType,
		input.Callback,
	))
	request := CreateInvoiceRequest{
		Amount:         fmt.Sprintf("%.2f", input.Amount),
		Checksum:       checksum,
		TransactionID:  input.TransactionID,
		ReturnType:     string(input.ReturnType),
		Callback:       input.Callback,
		GenerateToken:  utils.BoolToString(input.GetToken),
		SocialDeeplink: utils.BoolToString(input.SocialDeeplink),
	}
	client := resty.New()
	defer client.Close()
	var response *CreateInvoiceResponse
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+g.bearerToken).
		SetBody(request).
		SetResult(&response).
		Post(g.baseUrl + "/api/invoice")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	// if response.ErrorCode != "000" {
	// 	return nil, fmt.Errorf("%s", response.ErrorDesc)
	// }
	// Verify checksum
	expectedChecksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(
		response.Invoice,
		response.TransactionID,
	))
	if expectedChecksum != response.Checksum {
		return nil, fmt.Errorf("checksum verification failed")
	}
	return response, nil
}

// Inquiry - inquires about a transaction
//
// transactionId: Transaction ID
// return: InquiryResponse, error
func (g *golomtEcommerce) Inquiry(transactionId string) (*InquiryResponse, error) {
	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(
		transactionId,
		transactionId,
	))
	request := InquiryRequest{
		Checksum:      checksum,
		TransactionID: transactionId,
	}

	client := resty.New()
	defer client.Close()
	var response *InquiryResponse
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+g.bearerToken).
		SetBody(request).
		SetResult(&response).
		Post(g.baseUrl + "/api/inquiry")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.ErrorCode != "000" {
		return nil, fmt.Errorf("%s", response.ErrorDesc)
	}

	// Verify checksum
	var expectedChecksum string
	if response.Token != "" {
		expectedChecksum = utils.GenerateHMAC(g.secret, utils.AppendAsString(
			response.TransactionID,
			response.ErrorCode,
			response.Amount,
			response.Token,
		))
	} else {
		expectedChecksum = utils.GenerateHMAC(g.secret, utils.AppendAsString(
			response.TransactionID,
			response.ErrorCode,
			response.Amount,
		))
	}
	if expectedChecksum != response.Checksum {
		return nil, fmt.Errorf("checksum verification failed")
	}
	return response, nil
}

func (g *golomtEcommerce) ParsePushNotificationResponse(body []byte) (*InquiryResponse, error) {
	var response *InquiryResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorCode != "000" {
		return nil, fmt.Errorf("%s", response.ErrorDesc)
	}
	// Verify checksum
	var expectedChecksum string
	if response.Token != "" {
		expectedChecksum = utils.GenerateHMAC(g.secret, utils.AppendAsString(
			response.TransactionID,
			response.ErrorCode,
			response.Amount,
			response.Token,
		))
	} else {
		expectedChecksum = utils.GenerateHMAC(g.secret, utils.AppendAsString(
			response.TransactionID,
			response.ErrorCode,
			response.Amount,
		))
	}
	if expectedChecksum != response.Checksum {
		return nil, fmt.Errorf("checksum verification failed")
	}
	return response, nil
}

func (g *golomtEcommerce) CreateToken(input CreateTokenInput) (*CreateTokenResponse, error) {
	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(
		input.TransactionID,
		input.ReturnType,
		input.Callback,
	))
	request := CreateTokenRequest{
		Checksum:      checksum,
		TransactionID: input.TransactionID,
		ReturnType:    string(input.ReturnType),
		Callback:      input.Callback,
	}
	client := resty.New()
	defer client.Close()
	var response *CreateTokenResponse
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+g.bearerToken).
		SetBody(request).
		SetResult(&response).
		Post(g.baseUrl + "/api/confirmation")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	// Verify checksum
	expectedChecksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(
		response.Invoice,
		response.TransactionID,
	))
	if expectedChecksum != response.Checksum {
		return nil, fmt.Errorf("checksum verification failed")
	}
	return response, nil
}

func (g *golomtEcommerce) GetConfirmationUrl(input GetConfirmationInput) string {
	return fmt.Sprintf("https://ecommerce.golomtbank.com/confirmation/%v/%v", input.Lang, input.Invoice)
}

func (g *golomtEcommerce) CheckToken(transactionId string) (*CheckTokenResponse, error) {
	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(transactionId, transactionId))
	request := CheckTokenRequest{
		Checksum:      checksum,
		TransactionID: transactionId,
	}

	client := resty.New()
	defer client.Close()
	var response *CheckTokenResponse
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+g.bearerToken).
		SetBody(request).
		SetResult(&response).
		Post(g.baseUrl + "/api/get/token")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	if response.ErrorCode != "000" {
		return nil, fmt.Errorf("%s", response.ErrorDesc)
	}
	// Verify checksum
	expectedChecksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(
		response.TransactionID,
		response.Token,
	))
	if expectedChecksum != response.Checksum {
		return nil, fmt.Errorf("checksum verification failed")
	}
	return response, nil
}

func (g *golomtEcommerce) GetSettlementDetails(input GetSettlementDetailsInput) (*GetSettlementDetailsResponse, error) {
	checksum := utils.GenerateHMAC(g.secret, utils.AppendAsString(g.merchantId, g.bearerToken))
	request := GetSettlementDetailsRequest{
		Checksum:   checksum,
		Token:      g.bearerToken, //! maybe not bearer token
		MerchantId: g.merchantId,
	}
	client := resty.New()
	defer client.Close()
	var response *GetSettlementDetailsResponse
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+g.bearerToken).
		SetQueryParams(map[string]string{
			"start": input.Start,
			"end":   input.End,
			"page":  strconv.Itoa(input.Page),
			"size":  strconv.Itoa(input.Size),
		}).
		SetBody(request).
		SetResult(&response).
		Post(g.baseUrl + "/api/getSettlementDetails/details")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(string))
	}
	return response, nil
}
