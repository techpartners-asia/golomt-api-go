package openbank

import (
	"fmt"
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
	"resty.dev/v3"
)

// 9.33. Байгууллагын виртуал кредит карт токенжуулах
func (o *openbank) CardTokenize(body model.TokenizeReq) (string, error) {
	if err := o.auth(); err != nil {
		return "", err
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
			"client_id": o.clientID,
			"state":     o.state,
			"scope":     o.scope,
		}).
		Post(o.url + "/v1/card/corp/tokenize")
	if err != nil {
		return "", err
	}

	response = res.Bytes()
	if res.StatusCode() != 200 {
		if len(response) == 0 {
			return "", fmt.Errorf("%s-Golomt CG card corporate tokenize response: %s", time.Now().Format("20060102150405"), res.Status())
		}
		errResp, err := parseEncryptedResponse[*model.ErrorResp](response, o.DecryptAESCBC)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("%s-Golomt CG card corporate tokenize response: %s: %s", time.Now().Format("20060102150405"), errResp.Message, errResp.DebugMessage)
	}
	// return parseEncryptedResponse[string](response, o.DecryptAESCBC)

	responseData, err := o.DecryptAESCBC(string(response))
	if err != nil {
		return "", err
	}
	return responseData, nil
}

// 9.34. Токен цуцлах
func (o *openbank) CardTokenClose(body model.TokenCloseReq) (*model.TokenCloseResp, error) {
	return postEncrypted[*model.TokenCloseResp](o, "CRTOKCL", "/v1/card/token/close", body, requestOption{})
}

// 9.15. Токентэй картнаас гүйлгээ гаргах
func (o *openbank) CardPurchase(body model.CardPurchaseReq) (*model.CardPurchaseResp, error) {
	return postEncrypted[*model.CardPurchaseResp](o, "CDTXN", "/v1/card/purchase/", body, requestOption{})
}

// 9.32. Картын гүйлгээ шалгах
func (o *openbank) CardPurchaseCheck(body model.CardPurchaseCheckReq) (*model.CardPurchaseCheckResp, error) {
	return postEncrypted[*model.CardPurchaseCheckResp](o, "CRDTXNINQ", "/v1/card/tran/inq", body, requestOption{})
}

// 9.23. Мерчантын хуулга авах
func (o *openbank) CardMerchantStatement(body model.CardMerchantStatementReq, page model.PageReq) (*model.CardMerchantStatementResp, error) {
	return postEncrypted[*model.CardMerchantStatementResp](o, "MRCHSTMT", "/v1/card/merchant/statement", body, requestOption{
		withCode: true,
		query: map[string]string{
			"page_no":   page.PageNo,
			"page_size": page.PageSize,
		},
	})
}

// 9.1. Кредит картын дэлгэрэнгүй
func (o *openbank) CardCreditDetail(body model.CardCreditDetailReq) (*model.CardCreditDetailResp, error) {
	return postEncrypted[*model.CardCreditDetailResp](o, "CCDTLS", "/v1/card/credit/details", body, requestOption{})
}

// 9.9. Картын гүйлгээний мэдээлэл татах
func (o *openbank) CardTransaction(body model.CardTransactionReq) (*model.CardTransactionResp, error) {
	return postEncrypted[*model.CardTransactionResp](o, "CRDTRNDETS", "/v1/card/transaction-details", body, requestOption{})
}

// 9.11. Кредит карт хуулга харах
func (o *openbank) CardCreditStatement(body model.CardCreditStatementReq) ([]model.CardCreditStatementData, error) {
	return postEncrypted[[]model.CardCreditStatementData](o, "CCSTATM", "/v1/card/credit/statement", body, requestOption{})
}

// 9.2. Картын жагсаалт (Дебит, Кредит)
func (o *openbank) CardList(body model.CardListReq) ([]model.CardListData, error) {
	return postEncrypted[[]model.CardListData](o, "CRDTLST", "/v1/card/list", body, requestOption{})
}

// 9.35. UnionPay QR үүсгэх
func (o *openbank) UnionPayQRGenerate(body model.UnionPayQRReq) (*model.UnionPayQRResp, error) {
	return postEncrypted[*model.UnionPayQRResp](o, "CUNQRG", "/v1/card/union/qrgen", body, requestOption{})
}

// 9.36. UnionPay токен үүсгэх
func (o *openbank) UnionPayTokenCreate(body model.UnionPayTokenCreateReq) (*model.UnionPayTokenCreateResp, error) {
	return postEncrypted[*model.UnionPayTokenCreateResp](o, "CUNTCR", "/v1/card/union/token/create", body, requestOption{})
}

// 9.37. UnionPay токен өөрчлөх
func (o *openbank) UnionPayTokenUpdate(body model.UnionPayTokenUpdateReq) (*model.UnionPayTokenUpdateResp, error) {
	return postEncrypted[*model.UnionPayTokenUpdateResp](o, "CUNTUP", "/v1/card/union/token/update", body, requestOption{})
}

// 9.3. Кредит карт захиалах (Скоринг хийх кредит карт) нь тусдаа endpoint-гүй,
// зээлийн хэсгийн кредит картын хүсэлтийг ашиглана.

// 9.4. Кредит карт захиалах (Pre-approved virtual credit card)
func (o *openbank) CardCreditOrder(body model.CardCreditOrderReq) (*model.CardOrderResp, error) {
	return postEncrypted[*model.CardOrderResp](o, "CRECRDORD", "/v1/card/credit/order", body, requestOption{})
}

// 9.5. Дебит карт захиалах
func (o *openbank) CardDebitOrder(body model.CardDebitOrderReq) (*model.CardDebitOrderResp, error) {
	return postEncrypted[*model.CardDebitOrderResp](o, "DBCRDORD", "/v1/card/debit/order", body, requestOption{})
}

// 9.6. Их сургуулийн дебит карт захиалах
func (o *openbank) CardUniversityDebitOrder(body model.CardUniversityDebitOrderReq) (*model.CardUniversityDebitOrderResp, error) {
	return postEncrypted[*model.CardUniversityDebitOrderResp](o, "UNIDBCRD", "/v1/card/debit/order/university", body, requestOption{})
}

// 9.7. Prepaid карт захиалах
func (o *openbank) CardPrepaidOrder(body model.CardPrepaidOrderReq) (*model.CardOrderResp, error) {
	return postEncrypted[*model.CardOrderResp](o, "PPCDORD", "/v1/card/prepaid/order", body, requestOption{withCode: true})
}

// 9.8. Prepaid карт цэнэглэлт
func (o *openbank) CardPrepaidTopup(body model.CardPrepaidTopupReq) (*model.CardPrepaidTopupResp, error) {
	return postEncrypted[*model.CardPrepaidTopupResp](o, "PREPREPAY", "/v1/card/prepaid/repay", body, requestOption{})
}

// 9.10. Хүүхдийн карт захиалах
// Өмнө нь харилцагчийн зураг хуулах (CUSTIMGUP, /v1/customer/image/upload) хүсэлт хийгдсэн байх шаардлагатай.
func (o *openbank) CardChildOrder(body model.CardChildOrderReq) (*model.CardChildOrderResp, error) {
	return postEncrypted[*model.CardChildOrderResp](o, "CLRCRDORD", "/v1/card/color/order", body, requestOption{})
}

// 9.12. Кредит карт төлөв солих
func (o *openbank) CardCreditStatusChange(body model.CardCreditStatusChangeReq) (*model.CardMessageResp, error) {
	return postEncrypted[*model.CardMessageResp](o, "CRCDSTTCH", "/v1/card/credit/status/change", body, requestOption{})
}

// 9.13. Карт идэвхжүүлэх
func (o *openbank) CardActivate(body model.CardActivateReq) (*model.CardMessageResp, error) {
	return postEncrypted[*model.CardMessageResp](o, "CRDACT", "/v1/card/activate", body, requestOption{})
}

// 9.14. Кредит картын орлого хийх
func (o *openbank) CardCreditPayment(body model.CardCreditPaymentReq) (*model.CardStatusResp, error) {
	return postEncrypted[*model.CardStatusResp](o, "CCREPAY", "/v1/card/credit/repay", body, requestOption{})
}

// 9.16. Void transaction
func (o *openbank) CardVoid(body model.CardVoidReq) (*model.CardVoidResp, error) {
	return postEncrypted[*model.CardVoidResp](o, "CDTXNVD", "/v1/card/purchase/void", body, requestOption{withCode: true})
}

// 9.17. Байгууллага өөрийн картаар хийсэн гүйлгээг буцаах (void transaction)
func (o *openbank) CardCorpVoid(body model.CardVoidReq) (*model.CardCorpVoidResp, error) {
	return postEncrypted[*model.CardCorpVoidResp](o, "CGWCDVD", "/v1/card/purchase/void/cgw", body, requestOption{withCode: true})
}

// 9.18. Нэхэмжлэх төлөх
func (o *openbank) CardInvoicePay(body model.CardInvoicePayReq) (*model.CardInvoicePayResp, error) {
	return postEncrypted[*model.CardInvoicePayResp](o, "CDTXNINST", "/v1/card/instore", body, requestOption{})
}

// 9.19. Карт токенжуулах
func (o *openbank) CardTokenizeCustomer(body model.CardTokenizeCustomerReq) (*model.CardTokenizeCustomerResp, error) {
	return postEncrypted[*model.CardTokenizeCustomerResp](o, "CDTKN", "/v1/card/tokenization", body, requestOption{})
}

// 9.20. Мерчант бүртгэх
func (o *openbank) CardMerchantAdd(body model.CardMerchantAddReq) (*model.CardMerchantAddResp, error) {
	return postEncrypted[*model.CardMerchantAddResp](o, "CPTMERCH", "/v1/card/merchant/capture", body, requestOption{})
}

// 9.21. Терминал бүртгэх
func (o *openbank) CardTerminalAdd(body model.CardTerminalAddReq) (*model.CardTerminalAddResp, error) {
	return postEncrypted[*model.CardTerminalAddResp](o, "CPTTRML", "/v1/card/terminal/capture", body, requestOption{})
}

// 9.22. Reversal гүйлгээ
func (o *openbank) CardReversal(body model.CardReversalReq) (*model.CardReversalResp, error) {
	return postEncrypted[*model.CardReversalResp](o, "CDRVRSL", "/v1/card/reversal", body, requestOption{})
}

// 9.24. Терминалын жагсаалт авах
func (o *openbank) CardTerminalList(body model.CardTerminalListReq) (*model.CardTerminalListResp, error) {
	return postEncrypted[*model.CardTerminalListResp](o, "LSTTRML", "/v1/card/terminal/list", body, requestOption{withCode: true})
}

// 9.25. Картын мэдээлэл Openbank web дээр харах
func (o *openbank) CardWebView(body model.CardWebViewReq) (*model.CardWebViewResp, error) {
	return postEncrypted[*model.CardWebViewResp](o, "CDMNDTLS", "/v1/card/main/details", body, requestOption{})
}

// 9.26. Харилцагч дээр тухайн product-тай карт байгаа эсэх
func (o *openbank) CardProductCheck(body model.CardProductCheckReq) (*model.CardProductCheckResp, error) {
	return postEncrypted[*model.CardProductCheckResp](o, "CDCHK", "/v1/card/check", body, requestOption{})
}

// 9.27. Картын token replace хийх
func (o *openbank) CardTokenReplace(body model.CardTokenReplaceReq) (*model.CardStatusResp, error) {
	return postEncrypted[*model.CardStatusResp](o, "CDTKNRPLC", "/v1/card/token/replace", body, requestOption{})
}

// 9.28. Байгууллага токентэй гүйлгээ хийх
func (o *openbank) CardCorpPurchase(body model.CardCorpPurchaseReq, device model.CardDeviceInfo) (*model.CardPurchaseResp, error) {
	return postEncrypted[*model.CardPurchaseResp](o, "CDTXNDV", "/v1/card/purchase/device", body, requestOption{
		headers: map[string]string{
			"X-DEVICE-ID":   device.ID,
			"X-DEVICE-NAME": device.Name,
			"X-DEVICE-IP":   device.IP,
		},
	})
}

// 9.29. Картын шилжүүлэг хийх
func (o *openbank) CardTransfer(body model.CardTransferReq) (*model.CardTransferResp, error) {
	return postEncrypted[*model.CardTransferResp](o, "CDTRANSF", "/v1/card/transfer", body, requestOption{})
}

// 9.30. Карт руу орлого оруулах
func (o *openbank) CardDeposit(body model.CardDepositReq) (*model.CardDepositResp, error) {
	return postEncrypted[*model.CardDepositResp](o, "CDINUIP", "/v1/card/income/uip", body, requestOption{})
}

// 9.31. Картын орлогын гүйлгээний төлөв шалгах
func (o *openbank) CardDepositCheck(body model.CardDepositCheckReq) (*model.CardDepositCheckResp, error) {
	return postEncrypted[*model.CardDepositCheckResp](o, "CDINUIPD", "/v1/card/income/uip/detail", body, requestOption{})
}
