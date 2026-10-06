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
	// 9.9.	Картын гүйлгээний мэдээлэл татах
	CardTransaction(body model.CardTransactionReq) (*model.CardTransactionResp, error)
	// 9.11. Кредит карт хуулга харах
	CardCreditStatement(body model.CardCreditStatementReq) ([]model.CardCreditStatementData, error)
	// 9.15. Токентэй картнаас гүйлгээ гаргах
	CardPurchase(body model.CardPurchaseReq) (*model.CardPurchaseResp, error)
	// 9.23. Мерчантын хуулга авах
	CardMerchantStatement(body model.CardMerchantStatementReq, page model.PageReq) (*model.CardMerchantStatementResp, error)
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
	CategoryListInq(body model.CategoryReq) (*model.CategoryResp, error)
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
