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
	// 5.6.	Харилцах данс нээх
	AccountAdd(body model.AccountAddReq) (*model.AccountAddResp, error)
	// 5.7.	 Хадгаламжийн дансны дэлгэрэнгүй
	AccountDepositDetail(body model.AccountDetailReq) (*model.AccountDepositDetailResp, error)
	// 5.8.	Хадгаламжийн дансны хуулга харах
	// TODO: maybe response is not correct
	AccountDepositStatement(body model.StatementReq) (*model.AccountAddResp, error)
	// 5.9.	Хадгаламжийн данс нээх
	AccountDepositAdd(body model.AccountDepositAddReq) (*model.AccountAddResp, error)

	// 5.10. Дансны жагсаалт татах
	AccountList(body model.AccountListReq) (*model.AccountListResp, error)
	// 5.11.a Данс эзэмшигчнийн мэдээллэл авах /Голомт/
	AccountCustomerDetail(body model.AccountCustomerDetailReq) (*model.AccountCustomerDetailResp, error)
	// 5.11.b Данс эзэмшигчнийн мэдээллэл авах /Голомт бус/
	AccountOtherBankCustomerDetail(body model.AccountCustomerDetailReq) (*model.AccountOtherBankCustomerDetailResp, error)

	// 6.1.	Голомт Банк хоорондын гүйлгээ
	TransactionInBank(body model.TransactionReq) (*model.TransactionResp, error)
	// 6.2.	 Бусад банк хоорондын гүйлгээ
	TransactionOtherBank(body model.TransactionReq) (*model.TransactionResp, error)
	// 6.3. Байгууллага өөрийн дансаас гүйлгээ хийх
	TransactionSelf(body model.TransactionSelfReq) (*model.TransactionSelfResp, error)
	// 6.8. Гүйлгээний төлөв шалгах
	TransactionConfirm(body model.TransactionConfirmReq) (*model.TransactionConfirmResp, error)
	// 7.1.	Хот, аймагийн жагсаалт авах
	StateListInq(body model.StateListReq) ([]model.StateListResp, error)
	// 7.2.	Сум, дүүргийн жагсаалт авах
	DistrictListInq(body model.DistrictListReq) ([]model.DistrictListResp, error)
	// 7.3.	Категори төрлөөр сонголтын жагсаалт авах
	CategoryListInq(body model.CategoryReq) (*model.CategoryResp, error)
	// 7.4.	Ханшны мэдээлэл авах
	RateInq(body model.RateReq) (*model.RateResp, error)
	// 7.5.	Салбарын жагсаалт авах
	BranchListInq(body model.BranchListReq) ([]model.BranchListResp, error)
	// 7.6.	Бүтээгдэхүүн лавлах
	ProductListInq(body model.ProductListReq) ([]model.ProductData, error)

	// 8.1.	Картын токен авах
	CardTokenize(body model.TokenizeReq) (*model.TokenizeResp, error)
	// 8.2.	Картын токеныг хаах
	CardTokenClose(body model.TokenCloseReq) (*model.TokenCloseResp, error)
	// 8.3.	Картын гүйлгээ хийх
	CardPurchase(body model.CardPurchaseReq) (*model.CardPurchaseResp, error)
	// 8.4.	Картын гүйлгээний дүн шалгах
	CardPurchaseCheck(body model.CardPurchaseCheckReq) (*model.CardPurchaseCheckResp, error)
	// 8.5.	Картын хуулга харах
	CardMerchantStatement(body model.CardMerchantStatementReq, page model.PageReq) (*model.CardMerchantStatementResp, error)
	// 8.6.	Кредит картын дэлгэрэнгүй
	CardCreditDetail(body model.CardCreditDetailReq) (*model.CardCreditDetailResp, error)
	// 8.7.	Картын гүйлгээний мэдээлэл татах
	CardTransaction(body model.CardTransactionReq) (*model.CardTransactionResp, error)
	// 8.8.	Кредит карт хуулга харах
	CardCreditStatement(body model.CardCreditStatementReq) ([]model.CardCreditStatementData, error)
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
