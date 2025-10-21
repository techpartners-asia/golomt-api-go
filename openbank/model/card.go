package model

type (
	TokenizeReq struct {
		// Карт эзэмшигч иргэний РД
		CivilRegisterNo string `json:"civilRegisterNo" validate:"required"`
		// Байгууллагын РД
		CorporateRegisterNo string `json:"corpRegisterNo" validate:"required"`
	}
	TokenizeResp struct {
		// Статус
		Status string `json:"status"`
		// Картын маскласан дугаар  /379892*****1234/
		CardNumber string `json:"cardNumber"`
		// Токен
		Token string `json:"token"`
	}

	TokenCloseReq struct {
		// Токен
		Token string `json:"token" validate:"required"`
		// Карт эзэмшигч иргэний РД
		CivilRegisterNo string `json:"civilRegisterNo" validate:"required"`
		// Байгууллагын РД
		CorporateRegisterNo string `json:"corporateRegisterNo" validate:"required"`
	}
	TokenCloseResp struct {
		// Статус
		Status string `json:"status"`
		// Токен
		Token string `json:"token"`
	}

	CardPurchaseReq struct {
		// Мэдээлэлийг нь харах боломжтой банканд бүртгэлтэй харилцагчийн регистрийн дугаар буюу байгууллагын РД
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Гүйлгээний дүн
		Amount float64 `json:"amount" validate:"required"`
		// Гүйлгээний валют
		CurrencyCode string `json:"crncyCode" validate:"required"`
		// Тухайн гүйлгээ гаргах терминал дугаар
		TerminalID string `json:"terminalId" validate:"required"`
		// Урьдчилсан байдалаар Гуравдагч систем руу гаргаж өгсөн secretkey –ийн
		// тусламжтай тухайн гүйлгээ хийх мөчид гаргаж авсан 6 оронтоай код байна.
		ApproveCode string `json:"approveCode" validate:"required"`
	}
	CardPurchaseResp struct {
		// Гүйлгээний огноо
		PruchaseDate string `json:"prchDate"`
		// Гүйлгээний төлөв
		PruchaseStatus string `json:"prchStatus"`
		// Гүйлгээний approval code
		TransactionApprovalCode string `json:"tranApprovalCode"`
		// Гүйлгээний reference code
		TransactionReferenceCode string `json:"tranReferenceCode"`
	}

	CardPurchaseCheckReq struct {
		ReferenceNo string       `json:"refNo" validate:"required"`
		Amount      AmountDetail `json:"amount" validate:"required"`
	}
	CardPurchaseCheckResp struct {
		// Гүйлгээний төлөв
		TransactionID string `json:"tranId"`
		// Гүйлгээний RRN
		RRN string `json:"rrn"`
		// Гүйлгээ хийхэд ашигласан дугаар
		ApprovalCode string `json:"approvalCode"`
		// Хариу код. 00 - амжилттай
		ResponseCode string `json:"respCode"`
		// Тайлбар
		ResponseDesc string `json:"respDesc"`
	}

	CardMerchantStatementReq struct {
		// Тухайн байгууллагын регисртийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Хуулганы төрөл
		Type CardMerchantStatementTypeEnum `json:"type" validate:"required"`
		// Байгууллагын мерчантын дугаар
		Merchant string `json:"merchant" validate:"required"`
		// Тухайн хуулга авах терминал дугаар. Хоосон орхивол тухайн merchant дээр байгаа бүх terminal-н хуулгыг авна
		Terminal string `json:"terminal"`
		// Эхлэх огноо. YYYY-MM-DD
		StartDate string `json:"startDate" validate:"required"`
		// Дуусах огноо. YYYY-MM-DD
		EndDate string `json:"endDate" validate:"required"`
	}
	CardMerchantStatementResp struct {
		// Гүйлгээний дугаар
		ReferenceNo string `json:"referenceNo"`
		// Терминал дугаар
		Terminal string `json:"terminal"`
		// Картын дугаар
		CardNumber string `json:"cardNumber"`
		// Хариу тайлбар
		ResponseDesc string `json:"respDesc"`
		// Гүйлгээний дүн
		Amount float64 `json:"amount"`
		// Суваг
		Channel string `json:"channel"`
		// Гүйлгээний огноо
		Date string `json:"date"`
		// Хариу код
		ApprovalCode string `json:"approvalCode"`
		// Гүйлгээний шимтгэл
		Fee float64 `json:"fee"`
		// Гүйлгээний шимтгэл хасагдсан дүн
		NetAmount float64 `json:"netAmount"`
	}

	CardCreditDetailReq struct {
		// Тухайн байгууллагын регисртийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
	}
	CardCreditDetailResp struct {
		// Картын дугаар
		CardNumber string `json:"cardNumber"`
		// Картын нэр
		EmbossName string `json:"embossName"`
		// Картын төлөв
		Status string `json:"status"`
		// Картын дуусах огноо
		ExpireDate string `json:"expireDate"`
		// Картын пластикийн тайлбар
		ProductGroupDescription string `json:"productGroupDescription"`
		// Сүүлд төлбөр төлсөн огноо
		LastDueDate string `json:"lastDueDate"`
		// Сүүлд төлсөн мөнгөн дүн
		LastDueAmount string `json:"lastDueAmount"`
		// Төлөлт хийх боломжтой бага дүн
		MinimumPaymentDueAmount string `json:"minimumPaymentDueAmount"`
		// Картын боломжит үлдэгдэл
		AccountAvailableLimit string `json:"accountAvailableLimit"`
		// Картын зарцуулалт
		AccountOutstandingBalance string `json:"accountOutstandingBalance"`
	}

	CardTransactionReq struct {
		// Тухайн байгууллагын регисртийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
	}
	CardTransactionResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Карт эзэмшигчийн нэр
		CardName string `json:"cardName"`
		// Картын дугаар
		CardNumber string `json:"cardNumber"`
		// Гүйлгээний мэдээлэл
		Statements []CardStatement `json:"statements"`
	}
	CardStatement struct {
		// Гүйлгээ хийгдсэн огноо. Формат: yyyy-mm-dd
		TransactionDate string `json:"transactionDate"`
		// Гүйлгээ баталгаажсан огноо. Формат: yyyy-mm-dd
		PostDate string `json:"postDate"`
		// Гүйлгээний утга
		Description string `json:"description"`
		// Нэхэмжлэхийн дүн
		BillingAmount float64 `json:"billingAmount"`
		// Гүйлгээний дүн
		TransactionAmount float64 `json:"transactionAmount"`
	}

	CardCreditStatementReq struct {
		// Тухайн байгууллагын регисртийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Хуулга харах сарын интервалын бага утга. 1-12 хүртэл тоо утга байна.
		// Формат: yyyy.MM
		MonthStart string `json:"monthStart" validate:"required"`
		// Хуулга харах сарын интервалын дээд утга. 1-12 хүртэл тоо утга байна.
		// Формат: yyyy.MM
		MonthEnd string `json:"monthEnd"`
	}
	CardCreditStatementData struct {
		// Картын дугаар
		CardNumber     string          `json:"cardNumber"`
		CardName       string          `json:"cardName"`
		CreditLimit    float64         `json:"creditLimit"`
		MinimumPayment float64         `json:"minPmnt"`
		OpenBalance    float64         `json:"openBal"`
		CurrentBalance float64         `json:"currBal"`
		Month          int             `json:"month"`
		Statements     []CardStatement `json:"statements"`
	}
)
