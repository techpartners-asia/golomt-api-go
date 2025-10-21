package model

type (
	StatementReq struct {
		// Хуулга авах дансны дугаар
		AccountID string `json:"accountId" validate:"required"`
		// Данс эзэмшигчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Эхлэх огноо байна.
		// ISO Date форматтай байна
		StartDate string `json:"startDate" validate:"required"`
		// Дуусах огноо.
		// ISO Date форматтай байна
		EndDate string `json:"endDate"`
	}
	StatementResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Тухайн дансны дугаар
		AccountID string `json:"accountId"`
		// Дансны хуулганы жагсаалт
		Statements []Statement `json:"statements"`
	}
	Statement struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Дарааллын дугаар
		RecNum float64 `json:"recNum"`
		// Гүйлгээний дугаар
		TransactionId string `json:"tranId"`
		// Гүйлгээний огноо
		TranDate string `json:"tranDate"`
		// Debit – Зарлага
		// Credit – Орлого
		DrOrCr string `json:"drOrCr"`
		// Гүйлгээний дүн
		TranAmount float64 `json:"tranAmount"`
		// Гүйлгээний утга
		TranDesc string `json:"tranDesc"`
		// Гүйлгээ амжилттай болсон огноо
		TranPostedDate string `json:"tranPostedDate"`
		// Гүйлгээний валют
		TranCrnCode string `json:"tranCrnCode"`
		// Гүйлгээний ханш
		ExchRate float64 `json:"exchRate"`
		// Гүйлгээний дараах үлдэгдэл
		Balance string `json:"balance"`
		// Дансны нэр
		AccountName string `json:"accName"`
		// Дансны дугаар
		AccountNumber string `json:"accNum"`
	}
	AccountAddReq struct {
		// Харилцагчийн регисртийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Данс нээх валют
		Currency string `json:"currency" validate:"required"`
		// Дансны нэр
		Name string `json:"name" validate:"required"`
		// Дансы товч нэр 10 хүртэлх тэмдэгт байна
		ShortName string `json:"shortName" validate:"required"`
		// Банкнаас урдьчилан гэрээний үндсэнд гаргаж өгсөн бүтээгдэхүүний код
		// Лавлах төрөл: OPER
		SchemeCode string `json:"schemeCode" validate:"required"`
		// Доод үлдэгдэл шилжүүлэх данс. Бүтээгдэхүүний доод үлдэгдэл 0-с их тохиолдолд
		InitAccount string `json:"initAccount"`
		// Данс нээх доод үлдэгдэл, хоосон илгээсэн тохиолдолд бүтээгдэхүүний доод лимит-ээр данс үүсгэнэ
		InitAmount string `json:"initAmount" validate:"required"`
	}

	AccountAddResp struct {
		// Шинээр нээсэн дансны дугаар
		Account string `json:"account"`
	}

	AccountListReq struct {
		// Харилцагчийн регисртийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	AccountListResp struct {
		// Харилцах дансны жагсаалт
		OperAccounts []Account `json:"operAccounts"`
		// Хадгаламжийн дансны жагсаалт
		DepoAccounts []Account `json:"depoAccounts"`
		// Зээл дансны жагсаалт
		LoanAccounts []Account `json:"loanAccounts"`
	}

	Account struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Дансны дугаар
		AccountID string `json:"accountId"`
		// Дансы нэр
		AccountName string `json:"accountName"`
		// Дансы товч нэр
		ShortName string `json:"shortName"`
		// Дансы валют
		Currency string `json:"currency"`
		// Үндсэн салбарын дугаар
		BranchID string `json:"branchId"`
		// Social pay холбогдсон эсэх. Y - тийм, N - үгүй
		IsSocialPayConnected string `json:"isSocialPayConnected"`
		// Дансны төрөл
		AccountType AccountType `json:"accountType"`
	}

	AccountType struct {
		// Бүтээгдэхүүний код
		SchemeCode string `json:"schemeCode"`
		// Бүтээгдэхүүний төрөл
		SchemeType string `json:"schemeType"`
	}

	AccountTypeInqReq struct {
		// Дансны дугаар
		AccountID string `json:"accountId"`
	}

	AccountTypeInqResp struct {
		// Дансы төрөл. Дараах төрлүүд байна.
		//
		// 1.	DEPO – хадгаламж
		//
		// 2.	OPER – харилцах
		//
		// 3.	LOAN – зээл
		AccountType AccountTypeEnum `json:"accountType"`
	}

	AccountRenameReq struct {
		// Харилцагчийн регисртийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Дансны дугаар
		AccountID string `json:"accountId" validate:"required"`
		// Дансы товч нэр. 10 хүртэлх тэмдэгт байна.
		ShortName string `json:"shortName" validate:"required"`
	}
	AccountRenameResp struct {
		// Дансны дугаар
		AccountID string `json:"accountId"`
		// Амжилттай шинэчилсэн мессеж
		Message string `json:"message"`
		// Амжилттай шинэчилсэн эсхээс хамаарч SUCCESS эсвэл FAILED хариу буцаана
		Status string `json:"status"`
	}

	AccountDetailReq struct {
		// Дансны дугаар
		AccountID string `json:"accountId" validate:"required"`
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	AccountDetailResp struct {
		// Дансны дугаар
		AccountNumber string `json:"accountNumber"`
		// Дансы нэр
		AccoutnName string `json:"accountName"`
		// Дансы товч нэр
		AccountShortName string `json:"accountShortName"`
		// Дансы валют
		Currency string `json:"currency"`
		// Дансы эзэмшигчийн бүтэн нэр
		CustomerName string `json:"customerName"`
		// Хүндэтгэл. MR, MRS, MS, DR, SIR, PROF, MISS, etc.
		TitlePrefix string `json:"titlePrefix"`
		// Битүүмжийн код. T-total, C-credit, D-debit
		FreezeStatusCode FreezeStatusCodeEnum `json:"freezeStatusCode"`
		// Битүүмж хийгдсэн шалтгаан
		FreezeReasonCode string `json:"freezeReasonCode"`
		// Данс нээгдсэн огноо.
		// Формат: yyyy-MM-dd
		OpenDate string `json:"openDate"`
		// Дансны төлөв. Active, Inactive, Dormant
		Status AccountStatusEnum `json:"status"`
		// Харилцах дансны бүтээгдэхүүний нэршил
		ProductName string `json:"productName"`
		// Дансны хүүний ханш
		IntRate int `json:"intRate"`
		// Данс нь дээр хамтран бүртгэлтэй эсэх. Y - тийм, N - үгүй
		IsRelParty string `json:"isRelParty"`
		// бүтээгдэхүүний мэдээлэл
		Type AccountType `json:"type"`
	}

	AccountDepositDetailResp struct {
		// Дансны дугаар
		AccountNumber string `json:"accountNumber"`
		// Дансы нэр
		AccoutnName string `json:"accountName"`
		// Дансы товч нэр
		AccountShortName string `json:"accountShortName"`
		// Дансы валют
		Currency string `json:"currency"`
		// Дансы эзэмшигчийн бүтэн нэр
		CustomerName string `json:"customerName"`
		// Хүндэтгэл. MR, MRS, MS, DR, SIR, PROF, MISS, etc.
		TitlePrefix string `json:"titlePrefix"`
		// Битүүмжийн код. T-total, C-credit, D-debit
		FreezeStatusCode FreezeStatusCodeEnum `json:"freezeStatusCode"`
		// Битүүмж хийгдсэн шалтгаан
		FreezeReasonCode string `json:"freezeReasonCode"`
		// Данс нээгдсэн огноо.
		// Формат: yyyy-MM-dd
		OpenDate string `json:"openDate"`
		// Дансны төлөв. Active, Inactive, Dormant
		Status AccountStatusEnum `json:"status"`
		// Харилцах дансны бүтээгдэхүүний нэршил
		ProductName string `json:"productName"`
		// Дансны хүүний ханш
		IntRate int `json:"intRate"`
		// Данс нь дээр хамтран бүртгэлтэй эсэх. Y - тийм, N - үгүй
		IsRelParty string `json:"isRelParty"`
		// бүтээгдэхүүний мэдээлэл
		Type AccountType `json:"type"`
		// Хадгаламжийн хугацаа
		Term AccountTerm `json:"term,omitempty"`
		// Авто сунгалт бүртгэлтэй эсэх. Y - тийм, N - үгүй
		AutoRenewalFlag string `json:"autoRenewalFlag"`
		// Дансны IBAN дугаар
		IBAN string `json:"iban"`
		// Татварын хувь
		WTaxPercent int `json:"wTaxPcnt"`
		// Хугацаа нь дуусах огноо
		// Формат: yyyy-MM-dd
		MaturityDate string `json:"maturityDate"`
		// Хугацаа нь дуусах дүн
		MaturityAmount AmountDetail `json:"maturityAmt"`
		//
		OriginalMaturityAmount AmountDetail `json:"origMaturityAmt"`
	}

	AccountTerm struct {
		Month int `json:"month"`
		Day   int `json:"day"`
	}

	AccountDepositAddReq struct {
		// Харилцагчийн регисртийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Данс нээх валют
		Currency string `json:"currency" validate:"required"`
		// Дансны нэр
		Name string `json:"name" validate:"required"`
		// Дансы товч нэр 10 хүртэлх тэмдэгт байна
		ShortName string `json:"shortName" validate:"required"`
		// Банкнаас урдьчилан гэрээний үндсэнд гаргаж өгсөн бүтээгдэхүүний код
		// Лавлах төрөл: OPER
		SchemeCode string `json:"schemeCode" validate:"required"`
		// Тус бүтээгдэхүүнийг нээх боломжтой сар байна
		TermMonth int `json:"termMonth" validate:"required"`
		// Данс нээх мөнгөн дүн буюу Бүтээгдэхүүний доод дүнгээс багагүй байна
		Amount float64 `json:"amount" validate:"required"`
		// Доод үлдэгдэл шилжүүлэх данс. Бүтээгдэхүүний доод үлдэгдэл 0-с их тохиолдолд
		InitAccount string `json:"initAccount"`
	}

	AccountCustomerDetailReq struct {
		// Дансны дугаар
		AccountID string `json:"accountId" validate:"required"`
		// Тухайн банкны код.
		BankCode string `json:"bankCode"`
	}
	AccountOtherBankCustomerDetailResp struct {
		// Тухайн дансны дугаар зөв эсэх
		Vrfctn bool `json:"vrfctn"`
		// харилцагчийн утасны дугаар
		PhoneNumber string `json:"phneNb"`
		// харилцагчийн утасны дугаар
		MobileNumber string `json:"mobNb"`
		// харилцагчийн имэйл
		Email string `json:"emailAdr"`
		// Данс эзэмшигчийн нэр /masked/
		MaskedAccountName string `json:"maskedAccountName"`
		// Данс
		Account string `json:"account"`
		// Тухайн банкны код
		BankId string `json:"bankId"`
		// Алдааны тайлбар
		ErrorDesc string `json:"errDesc"`
		// Группын төлөв
		GroupStatus string `json:"grpSts"`
		// Шалтгаан
		Reason string `json:"rsn"`
	}
	AccountCustomerDetailResp struct {
		// Данс эзэмшигчийн нэр /masked/
		MarkedAccountName string `json:"markedAccountName"`
		// Дансны төлөв. A - Active, I - Inactive, D - Dormant
		Status string `json:"status"`
	}

	AccountBalcInqReq struct {
		//Дансны дугаар
		AccountID string `json:"accountId" validate:"required"`
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}

	AccountBalcInqResp struct {
		// Тухайн дансны дугаар
		AccountID string `json:"accountId"`
		// Данс эзэмшигчийн нэр
		AccountName string `json:"accountName"`
		// Дансны валют
		Currency string `json:"currency"`
		// Дансны үлдэгдлийн жагсаалт
		BalanceLL []Balance `json:"balanceLL"`
	}
	Balance struct {
		// Үлдэгдлийн төрөл.
		// AVAIL - Боломжит үлдэгдэл.
		Type string `json:"type"`
		// Дансны үлдэгдэл дүн
		Amount AmountDetail `json:"amount"`
	}
	AmountDetail struct {
		// Тоон утга
		Value float64 `json:"value"`
		// Валют
		Currency string `json:"currency"`
		// Шилжүүлэг хийх банкны код буюу Голомт банк (15) байна. Лавлах төрөл: BANK
		Bank string `json:"bank"`
	}
)
