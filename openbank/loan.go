package openbank

import (
	"encoding/json"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
)

// 7.1. Зээлийн дансны дэлгэрэнгүй
func (o *openbank) LoanDetail(body model.LoanDetailReq) (*model.LoanDetailResp, error) {
	return postEncrypted[*model.LoanDetailResp](o, "ACCTLOANDET", "/v1/account/loan/details", body, requestOption{})
}

// 7.2. Зээлийн дансны график
// SPEC: хариу нь `account` талбартай гэж тодорхойлогдсон
func (o *openbank) LoanSchedule(body model.LoanDetailReq) (*model.LoanScheduleResp, error) {
	return postEncrypted[*model.LoanScheduleResp](o, "ACCTLOANGRP", "/v1/account/loan/graphic", body, requestOption{})
}

// 7.3. Цалингийн зээл болон хэрэглээний зээлийн хүсэлт
func (o *openbank) LoanRequest(body model.LoanRequestReq) (*model.LoanRequestResp, error) {
	return postEncrypted[*model.LoanRequestResp](o, "LNRQTADD", "/v1/loan/request", body, requestOption{})
}

// 7.4. Хамтран нэмж зээл судлуулах
func (o *openbank) LoanJointAdd(body model.LoanJointAddReq) (*model.LoanJointAddResp, error) {
	return postEncrypted[*model.LoanJointAddResp](o, "LNJNTADD", "/v1/loan/joint/add", body, requestOption{})
}

// 7.5. Кредит картын хүсэлт
func (o *openbank) LoanCardRequest(body model.LoanCardRequestReq) (*model.LoanRequestResp, error) {
	return postEncrypted[*model.LoanRequestResp](o, "CCDRQTADD", "/v1/loan/card", body, requestOption{})
}

// 7.6. Тэтгэврийн зээлийн хүсэлт
// loanInfo.isPensionScoringRequest: Y - scoring хүсэлт, N - зээлийн хүсэлт
func (o *openbank) LoanPensionRequest(body model.LoanRequestReq) (*model.LoanRequestResp, error) {
	return postEncrypted[*model.LoanRequestResp](o, "LNRQTADD", "/v1/loan/request", body, requestOption{})
}

// 7.7. Зээл хаах
func (o *openbank) LoanCorporateClosure(body model.LoanCorporateClosureReq) (*model.LoanCorporateClosureResp, error) {
	return postEncrypted[*model.LoanCorporateClosureResp](o, "CLOCLS", "/v1/loan/corporate/closure", body, requestOption{})
}

// 7.8. Байгууллагын зээл
func (o *openbank) LoanCorporateRequest(body model.LoanCorporateReq) (*model.LoanCorporateResp, error) {
	return postEncrypted[*model.LoanCorporateResp](o, "CLOADD", "/v1/loan/corporate/request", body, requestOption{})
}

// 7.9. Харилцагч авах зээлийн дүнг баталгаажуулах
func (o *openbank) LoanConfirm(body model.LoanConfirmReq) (*model.LoanConfirmResp, error) {
	return postEncrypted[*model.LoanConfirmResp](o, "LNSTTCNF", "/v1/loan/confirm", body, requestOption{})
}

// 7.10. Харилцагчийн зээлийн хүсэлтийн жагсаалт
func (o *openbank) LoanList(body model.LoanListReq) ([]model.LoanListData, error) {
	return postEncrypted[[]model.LoanListData](o, "OLNINQ", "/v1/loan/list", body, requestOption{noScope: true})
}

// 7.11. Зээлийн онлайн тооцооллын төлөв лавлах
func (o *openbank) LoanStatus(body model.LoanAppReq) (*model.LoanStatusResp, error) {
	return postEncrypted[*model.LoanStatusResp](o, "LNSTTINQ", "/v1/loan/status", body, requestOption{noScope: true})
}

// 7.12. Зээлийн дэд гэрээ татах
func (o *openbank) LoanContract(body model.LoanAppReq) (*model.LoanContractResp, error) {
	return postPlain[*model.LoanContractResp](o, "CNTINQ", "/v1/utility/contract/consumer/loan", body)
}

// 7.13. Зээлийн хүсэлт баталгаажуулах
func (o *openbank) LoanVerify(body model.LoanAppReq) (*model.LoanRequestResp, error) {
	return postEncrypted[*model.LoanRequestResp](o, "APPVERIF", "/v1/loan/verify", body, requestOption{noScope: true})
}

// 7.14. Зээлийн хүсэлт цуцлах
func (o *openbank) LoanCancel(body model.LoanCancelReq) (*model.LoanRequestResp, error) {
	return postEncrypted[*model.LoanRequestResp](o, "APPCANCL", "/v1/loan/cancel", body, requestOption{noScope: true})
}

// 7.15. Зээлийн идэвхитэй хүсэлт байгаа эсэхийг шалгах
func (o *openbank) LoanCheckRecord(body model.LoanCheckRecordReq) (*model.LoanCheckRecordResp, error) {
	return postEncrypted[*model.LoanCheckRecordResp](o, "LNAPPCHKREC", "/v1/loan/check-record", body, requestOption{noScope: true})
}

// 7.16. Зээлийн онлайн тооцооллын callback хүсэлт
// Голомт банкнаас гуравдагч системийн callbackUrl руу илгээх POST хүсэлтийн
// body-г хүлээж авах функц. Хариуд model.LoanCallbackResp буцаана.
func ParseLoanCallback(request []byte) (*model.LoanCallbackReq, error) {
	var result *model.LoanCallbackReq
	if err := json.Unmarshal(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// 7.17. Хэрэглээний зээл дээр ногоон бүтээгдэхүүний жагсаалт татах
// Голомт банкнаас гуравдагч системийн URL руу илгээх хүсэлтийн body-г
// хүлээж авах функц. Хариуд model.LoanGreenProductResp буцаана.
func ParseLoanGreenProductRequest(request []byte) (*model.LoanGreenProductReq, error) {
	var result *model.LoanGreenProductReq
	if err := json.Unmarshal(request, &result); err != nil {
		return nil, err
	}
	return result, nil
}
