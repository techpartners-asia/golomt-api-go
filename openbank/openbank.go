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
	TransactionInternal(body model.TransactionReq) (*model.TransactionResp, error)
	// 6.2.	 Бусад банк хоорондын гүйлгээ
	TransactionOtherBank(body model.TransactionReq) (*model.TransactionResp, error)
	// 6.3.	Гүйлгээний төлөв шалгах
	TransactionCheck(body model.TransactionCheckReq) (*model.TransactionCheckResp, error)
	// 7.4.	Ханшны мэдээлэл авах
	RateInq(body model.RateReq) (*model.RateResp, error)
	// 7.5.	Салбарын жагсаалт авах
	BranchListInq(body model.BranchListReq) (*model.BranchListResp, error)
	// 7.6.	Бүтээгдэхүүн лавлах
	ProductListInq(body model.ProductListReq) (*[]model.ProductData, error)
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
	}
}
