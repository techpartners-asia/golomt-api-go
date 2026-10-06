package openbank

import (
	"time"

	"github.com/techpartners-asia/golomt-api-go/openbank/model"
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
	authObject       *model.AuthResp
	xGolomtKey       string
	clientID         string
	state            string
	scope            string
}

type Openbank interface {
	// 4.4.	Сервисүүдийг багцаар авах
	ServicesAccess(body model.ServiceListReq) (*model.ServiceListResp, error)
	// 4.5.	Бүртгэлтэй дугаар татах
	GetPhone() (*model.GetPhoneResp, error)
	// 4.6.	Бүртгэлтэй дугаар руу OTP код илгээх
	OTPSend(phone string) (*model.OTPSendResp, error)
	// 4.7.	OTP шалгах
	OTPVerify(body model.OTPVerifyReq) (*model.OTPVerifyResp, error)
	// 4.8.	ХУР систем OTP илгээх
	XypOTPSend() (*model.OTPSendResp, error)
	// 4.9.	ХУР систем OTP шалгах
	XypOTPVerify(body model.OTPXypVerifyReq) (*model.OTPVerifyResp, error)
	// 4.10. Тоон гарын үсгээр баталгаажуулах
	DigitalSignature() (*model.DigitalSignatureResp, error)
	// 5.1.	Дансны үлдэгдэл
	AccountBalcInq(body model.AccountBalcInqReq) (*model.AccountBalcInqResp, error)
	// 5.2.	Дансны төрөл шалгах
	AccountTypeInq(body model.AccountTypeInqReq) (*model.AccountTypeInqResp, error)
	// 5.3.	Дансны товч нэр солих
	AccountRename(body model.AccountRenameReq) (*model.AccountRenameResp, error)
	// 5.4.	Харилцах дансны дэлгэрэнгүй мэдээлэл харах
	AccountDetail(body model.AccountDetailReq) (*model.AccountDetailResp, error)
	// 5.5.	Харилцах дансны хуулга харах
	AccountStatement(body model.StatementReq) (*model.StatementResp, error)
	// 5.6.	Харилцах дансны хуулга хуудаслалтай харах
	AccountStatementPage(body model.StatementPageReq) (*model.StatementPageResp, error)
	// 5.7.	Харилцах данс нээх
	AccountAdd(body model.AccountAddReq) (*model.AccountAddResp, error)
	// 5.8.	Хадгаламжийн дансны дэлгэрэнгүй
	AccountDepositDetail(body model.AccountDetailReq) (*model.AccountDepositDetailResp, error)
	// 5.9.	Хадгаламжийн дансны хуулга харах
	AccountDepositStatement(body model.StatementReq) (*model.AccountAddResp, error)
	// 5.10. Хадгаламжийн данс нээх
	AccountDepositAdd(body model.AccountDepositAddReq) (*model.AccountAddResp, error)
	// 5.11. Дансны жагсаалт татах
	AccountList(body model.AccountListReq) (*model.AccountListResp, error)
	// 5.12.a Данс эзэмшигчнийн мэдээллэл авах /Голомт/
	AccountCustomerDetail(body model.AccountCustomerDetailReq) (*model.AccountCustomerDetailResp, error)
	// 5.12.b Данс эзэмшигчнийн мэдээллэл авах /Голомт бус/
	AccountOtherBankCustomerDetail(body model.AccountCustomerDetailReq) (*model.AccountOtherBankCustomerDetailResp, error)

	// 6.1. Харилцагч шалгах
	CustomerCheck(body model.CustomerCheckReq) (*model.CustomerCheckResp, error)
	// 6.2. Харилцагчийн утасны дугаар мөн эсхийг шалгах
	CustomerPhoneCheck(body model.CustomerPhoneCheckReq) (*model.CustomerPhoneCheckResp, error)
	// 6.3. Харилцагчийн мэдээлэл харах
	CustomerInquire(body model.CustomerInquireReq) (*model.CustomerInquireResp, error)
	// 6.4. Иргэн – Харилцагч бүртгэх (харилцагчийн зураг хуулах)
	CustomerImageUpload(body model.CustomerImageUploadReq) (*model.CustomerImageUploadResp, error)
	// 6.4. Иргэн – Харилцагч бүртгэх
	CustomerRetailAdd(body model.CustomerRetailAddReq) (*model.CustomerAddResp, error)
	// 6.5. Байгууллага - Харилцагч бүртгэх
	CustomerCorpAdd(body model.CustomerCorpAddReq) (*model.CustomerAddResp, error)
	// 6.6. Харилцагчийн мэдээлэл бүртгэх
	CustomerSaveInfo(body model.CustomerSaveInfoReq) (*model.CustomerSaveInfoResp, error)
	// 6.7. Бүртгэгдсэн харилцагчийн мэдээлэл татах
	CustomerGetInfo(body model.CustomerGetInfoReq) (*model.CustomerGetInfoResp, error)
	// 6.8. Харилцагчийн иргэний үнэмлэхний зураг upload хийх /BASE64/
	CustomerNewImageUpload(body model.CustomerNewImageUploadReq) (*model.CustomerNewImageUploadResp, error)
	// 6.9. Харилцагчийн иргэний үнэмлэхний зураг upload хийх /FORM-DATA/
	CustomerNewImageFormUpload(input model.CustomerNewImageFormUploadInput) (*model.CustomerNewImageUploadResp, error)
	// 6.10. Харилцагчийн иргэний үнэмлэхний зураг татах
	CustomerImageDownload(body model.CustomerImageDownloadReq) (*model.CustomerImageDownloadResp, error)
	// 6.11. Байгууллага - Харилцагчийн мэдээлэл татах
	CustomerCorpDetail(body model.CustomerCorpDetailReq) (*model.CustomerCorpDetailResp, error)
	// 6.12. Байгууллага - Харилцагчийн limit-н мэдээлэл татах
	CustomerCorpLimitDetail(body model.CustomerCorpLimitDetailReq) (*model.CustomerCorpLimitDetailResp, error)

	// 7.1. Зээлийн дансны дэлгэрэнгүй
	LoanDetail(body model.LoanDetailReq) (*model.LoanDetailResp, error)
	// 7.2. Зээлийн дансны график
	LoanSchedule(body model.LoanDetailReq) (*model.LoanScheduleResp, error)
	// 7.3. Цалингийн зээл болон хэрэглээний зээлийн хүсэлт
	LoanRequest(body model.LoanRequestReq) (*model.LoanRequestResp, error)
	// 7.4. Хамтран нэмж зээл судлуулах
	LoanJointAdd(body model.LoanJointAddReq) (*model.LoanJointAddResp, error)
	// 7.5. Кредит картын хүсэлт
	LoanCardRequest(body model.LoanCardRequestReq) (*model.LoanRequestResp, error)
	// 7.6. Тэтгэврийн зээлийн хүсэлт
	LoanPensionRequest(body model.LoanRequestReq) (*model.LoanRequestResp, error)
	// 7.7. Зээл хаах
	LoanCorporateClosure(body model.LoanCorporateClosureReq) (*model.LoanCorporateClosureResp, error)
	// 7.8. Байгууллагын зээл
	LoanCorporateRequest(body model.LoanCorporateReq) (*model.LoanCorporateResp, error)
	// 7.9. Харилцагч авах зээлийн дүнг баталгаажуулах
	LoanConfirm(body model.LoanConfirmReq) (*model.LoanConfirmResp, error)
	// 7.10. Харилцагчийн зээлийн хүсэлтийн жагсаалт
	LoanList(body model.LoanListReq) ([]model.LoanListData, error)
	// 7.11. Зээлийн онлайн тооцооллын төлөв лавлах
	LoanStatus(body model.LoanAppReq) (*model.LoanStatusResp, error)
	// 7.12. Зээлийн дэд гэрээ татах
	LoanContract(body model.LoanAppReq) (*model.LoanContractResp, error)
	// 7.13. Зээлийн хүсэлт баталгаажуулах
	LoanVerify(body model.LoanAppReq) (*model.LoanRequestResp, error)
	// 7.14. Зээлийн хүсэлт цуцлах
	LoanCancel(body model.LoanCancelReq) (*model.LoanRequestResp, error)
	// 7.15. Зээлийн идэвхитэй хүсэлт байгаа эсэхийг шалгах
	LoanCheckRecord(body model.LoanCheckRecordReq) (*model.LoanCheckRecordResp, error)
	// 7.16. (ParseLoanCallback) болон 7.17. (ParseLoanGreenProductRequest) нь
	// Голомтоос ирэх хүсэлт тул package-level функцээр хэрэгжсэн.

	// 8.1.	Голомт Банк хоорондын гүйлгээ
	TransactionInBank(body model.TransactionReq) (*model.TransactionResp, error)
	// 8.2.	Бусад банк хоорондын гүйлгээ
	TransactionOtherBank(body model.TransactionReq) (*model.TransactionResp, error)
	// 8.3.	Байгууллага өөрийн дансаас гүйлгээ хийх
	TransactionSelf(body model.TransactionSelfReq) (*model.TransactionSelfResp, error)
	// 8.4.	Гаалийн гүйлгээ хийх
	TransactionCustomsPay(body model.CustomsPayReq) (*model.CustomsPayResp, error)
	// 8.5.	Татварын гүйлгээ хийх
	TransactionTaxPay(body model.TaxPayReq) (*model.TaxPayResp, error)
	// 8.6.	Гүйлгээ буцаах
	TransactionRefund(body model.TransactionRefundReq) (*model.TransactionRefundResp, error)
	// 8.7.	Гүйлгээ шалгах
	TransactionCheck(body model.TransactionCheckReq) (*model.TransactionCheckResp, error)
	// 8.8.	Багц гүйлгээ хийх
	TransactionBatch(body model.TransactionBatchReq) (*model.TransactionBatchResp, error)
	// 8.9.	Багц гүйлгээний төлөв шалгах
	TransactionBatchCheck(body model.TransactionBatchCheckReq, page model.PageReq) (*model.TransactionBatchCheckResp, error)
	// 8.10. Гүйлгээний төлөв шалгах
	TransactionConfirm(body model.TransactionConfirmReq) (*model.TransactionConfirmResp, error)
	// 8.11. Багц гүйлгээ файлаар хийх
	TransactionBatchFile(input model.TransactionBatchFileInput) (*model.TransactionBatchFileResp, error)
	// 8.12. Татварын төлбөрийн жагсаалт харах TIN
	TaxListByTIN(body model.TaxTINInqReq) ([]model.TaxTINInqData, error)
	// 8.13. Татварын төлбөрийн жагсаалт харах PIN
	TaxListByPIN(body model.TaxPINInqReq) ([]model.TaxPINInqData, error)
	// 8.14. Татварын төлбөрийн нэхэмжлэхийн дугаараар лавлагаа авах
	TaxInvoiceInq(body model.TaxInvoiceInqReq) (*model.TaxInvoiceInqResp, error)
	// 8.15. Гаалийн төлбөрийн нэхэмжлэхийн дугаараар лавлагаа авах
	CustomsInvoiceInq(body model.CustomsInvoiceInqReq) (*model.CustomsInvoiceInqResp, error)
	// 8.16. Файлаар хийсэн багц гүйлгээний дэлгэрэнгүй татах
	TransactionBatchFileInq(body model.TransactionBatchFileInqReq) (*model.TransactionBatchFileInqResp, error)

	// 9.1.	Кредит картын дэлгэрэнгүй
	CardCreditDetail(body model.CardCreditDetailReq) (*model.CardCreditDetailResp, error)
	// 9.2.	Картын жагсаалт (Дебит, Кредит)
	CardList(body model.CardListReq) ([]model.CardListData, error)
	// 9.3. Кредит карт захиалах (Скоринг) — 7.5 LoanCardRequest-ийг ашиглана
	// 9.4. Кредит карт захиалах (Pre-approved virtual credit card)
	CardCreditOrder(body model.CardCreditOrderReq) (*model.CardOrderResp, error)
	// 9.5. Дебит карт захиалах
	CardDebitOrder(body model.CardDebitOrderReq) (*model.CardDebitOrderResp, error)
	// 9.6. Их сургуулийн дебит карт захиалах
	CardUniversityDebitOrder(body model.CardUniversityDebitOrderReq) (*model.CardUniversityDebitOrderResp, error)
	// 9.7. Prepaid карт захиалах
	CardPrepaidOrder(body model.CardPrepaidOrderReq) (*model.CardOrderResp, error)
	// 9.8. Prepaid карт цэнэглэлт
	CardPrepaidTopup(body model.CardPrepaidTopupReq) (*model.CardPrepaidTopupResp, error)
	// 9.9.	Картын гүйлгээний мэдээлэл татах
	CardTransaction(body model.CardTransactionReq) (*model.CardTransactionResp, error)
	// 9.10. Хүүхдийн карт захиалах
	CardChildOrder(body model.CardChildOrderReq) (*model.CardChildOrderResp, error)
	// 9.11. Кредит карт хуулга харах
	CardCreditStatement(body model.CardCreditStatementReq) ([]model.CardCreditStatementData, error)
	// 9.12. Кредит карт төлөв солих
	CardCreditStatusChange(body model.CardCreditStatusChangeReq) (*model.CardMessageResp, error)
	// 9.13. Карт идэвхжүүлэх
	CardActivate(body model.CardActivateReq) (*model.CardMessageResp, error)
	// 9.14. Кредит картын орлого хийх
	CardCreditPayment(body model.CardCreditPaymentReq) (*model.CardStatusResp, error)
	// 9.15. Токентэй картнаас гүйлгээ гаргах
	CardPurchase(body model.CardPurchaseReq) (*model.CardPurchaseResp, error)
	// 9.16. Void transaction
	CardVoid(body model.CardVoidReq) (*model.CardVoidResp, error)
	// 9.17. Байгууллага өөрийн картаар хийсэн гүйлгээг буцаах (void transaction)
	CardCorpVoid(body model.CardVoidReq) (*model.CardCorpVoidResp, error)
	// 9.18. Нэхэмжлэх төлөх
	CardInvoicePay(body model.CardInvoicePayReq) (*model.CardInvoicePayResp, error)
	// 9.19. Карт токенжуулах
	CardTokenizeCustomer(body model.CardTokenizeCustomerReq) (*model.CardTokenizeCustomerResp, error)
	// 9.20. Мерчант бүртгэх
	CardMerchantAdd(body model.CardMerchantAddReq) (*model.CardMerchantAddResp, error)
	// 9.21. Терминал бүртгэх
	CardTerminalAdd(body model.CardTerminalAddReq) (*model.CardTerminalAddResp, error)
	// 9.22. Reversal гүйлгээ
	CardReversal(body model.CardReversalReq) (*model.CardReversalResp, error)
	// 9.23. Мерчантын хуулга авах
	CardMerchantStatement(body model.CardMerchantStatementReq, page model.PageReq) (*model.CardMerchantStatementResp, error)
	// 9.24. Терминалын жагсаалт авах
	CardTerminalList(body model.CardTerminalListReq) (*model.CardTerminalListResp, error)
	// 9.25. Картын мэдээлэл Openbank web дээр харах
	CardWebView(body model.CardWebViewReq) (*model.CardWebViewResp, error)
	// 9.26. Харилцагч дээр тухайн product-тай карт байгаа эсэх
	CardProductCheck(body model.CardProductCheckReq) (*model.CardProductCheckResp, error)
	// 9.27. Картын token replace хийх
	CardTokenReplace(body model.CardTokenReplaceReq) (*model.CardStatusResp, error)
	// 9.28. Байгууллага токентэй гүйлгээ хийх
	CardCorpPurchase(body model.CardCorpPurchaseReq, device model.CardDeviceInfo) (*model.CardPurchaseResp, error)
	// 9.29. Картын шилжүүлэг хийх
	CardTransfer(body model.CardTransferReq) (*model.CardTransferResp, error)
	// 9.30. Карт руу орлого оруулах
	CardDeposit(body model.CardDepositReq) (*model.CardDepositResp, error)
	// 9.31. Картын орлогын гүйлгээний төлөв шалгах
	CardDepositCheck(body model.CardDepositCheckReq) (*model.CardDepositCheckResp, error)
	// 9.32. Картын гүйлгээ шалгах
	CardPurchaseCheck(body model.CardPurchaseCheckReq) (*model.CardPurchaseCheckResp, error)
	// 9.33. Байгууллагын виртуал кредит карт токенжуулах
	CardTokenize(body model.TokenizeReq) (string, error)
	// 9.34. Токен цуцлах
	CardTokenClose(body model.TokenCloseReq) (*model.TokenCloseResp, error)
	// 9.35. UnionPay QR үүсгэх
	UnionPayQRGenerate(body model.UnionPayQRReq) (*model.UnionPayQRResp, error)
	// 9.36. UnionPay токен үүсгэх
	UnionPayTokenCreate(body model.UnionPayTokenCreateReq) (*model.UnionPayTokenCreateResp, error)
	// 9.37. UnionPay токен өөрчлөх
	UnionPayTokenUpdate(body model.UnionPayTokenUpdateReq) (*model.UnionPayTokenUpdateResp, error)

	// 10.1. Хот, аймагийн жагсаалт авах
	StateListInq(body model.StateListReq) ([]model.StateListResp, error)
	// 10.2. Сум, дүүргийн жагсаалт авах
	DistrictListInq(body model.DistrictListReq) ([]model.DistrictListResp, error)
	// 10.3. Категори төрлөөр сонголтын жагсаалт авах
	CategoryListInq(body model.CategoryReq) ([]model.CategoryResp, error)
	// 10.4. Ханшны мэдээлэл авах
	RateInq(body model.RateReq) (*model.RateResp, error)
	// 10.5. Салбарын жагсаалт авах
	BranchListInq(body model.BranchListReq) ([]model.BranchListResp, error)
	// 10.6. Бүтээгдэхүүн лавлах
	ProductListInq(body model.ProductListReq) ([]model.ProductData, error)
	// 10.7. RateCode лавлах
	RateCodeInq(body model.RateCodeReq) (*model.RateCodeResp, error)
	// 10.8. Exchange Rate лавлах
	ExchangeRateInq(body model.ExchangeRateReq) (*model.ExchangeRateResp, error)
}

func New(input model.OpenbankInput) Openbank {
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
		xGolomtKey:       input.XGolomtKey,
	}
}

func (o *openbank) SetOAuthResponse(response model.OAuthResp) {
	o.clientID = response.ClientID
	o.state = response.State
	o.scope = response.Scope
}
