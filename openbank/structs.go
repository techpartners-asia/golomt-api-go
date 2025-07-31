package openbank

type (
	API struct {
		Url         string
		Method      string
		ServiceName string
	}

	BaseInput struct {
		OrganizationName string `json:"organizationName"`
		Username         string `json:"username"`
		Password         string `json:"password"`
		IvKey            string `json:"ivKey"`
		SessionKey       string `json:"sessionKey"`
		Url              string `json:"url"`
		RegisterNo       string `json:"registerNo"`
		ClientID         string `json:"clientId"`
	}
	AuthReq struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	AuthResp struct {
		AuthMode     string `json:"authMode"`
		RequestId    string `json:"requestId"`
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
		TokenType    string `json:"tokenType"`
		ExpiresIn    int    `json:"expiresIn"`
	}
	StatementReq struct {
		AccountID  string `json:"accountId"`
		RegisterNo string `json:"registerNo"`
		StartDate  string `json:"startDate"`
		EndDate    string `json:"endDate"`
	}
	StatementResp struct {
		RequestID  string      `json:"requestId"`
		AccountID  string      `json:"accountId"`
		Statements []Statement `json:"statements"`
	}
	Statement struct {
		RequestID      string  `json:"requestId"`
		RecNum         float64 `json:"recNum"`
		TranId         string  `json:"tranId"`
		TranDate       string  `json:"tranDate"`
		DrOrCr         string  `json:"drOrCr"`
		TranAmount     float64 `json:"tranAmount"`
		TranDesc       string  `json:"tranDesc"`
		TranPostedDate string  `json:"tranPostedDate"`
		TranCrnCode    string  `json:"tranCrnCode"`
		ExchRate       float64 `json:"exchRate"`
		Balance        string  `json:"balance"`
		AccName        string  `json:"accName"`
		AccNum         string  `json:"accNum"`
	}
	AccountListReq struct {
		RegisterNo string `json:"registerNo"`
	}
	AccountListResp struct {
		OperAccounts []Account `json:"operAccounts"`
		DepoAccounts []Account `json:"depoAccounts"`
		LoanAccounts []Account `json:"loanAccounts"`
	}

	Account struct {
		RequestID            string      `json:"requestId"`
		AccountID            string      `json:"accountId"`
		AccountName          string      `json:"accountName"`
		ShortName            string      `json:"shortName"`
		Currency             string      `json:"currency"`
		BranchID             string      `json:"branchId"`
		IsSocialPayConnected string      `json:"isSocialPayConnected"`
		AccountType          AccountType `json:"accountType"`
	}

	AccountType struct {
		SchemeCode string `json:"schemeCode"`
		SchemeType string `json:"schemeType"`
	}

	ServiceListReq struct {
		RegisterNo string   `json:"registerNo"`
		Services   []string `json:"services"`
		Code       string   `json:"code"`
	}

	ServiceListResp struct {
		ClientID     string `json:"clientId"`
		ResponseType string `json:"responseType"`
		RedirectUri  string `json:"redirectUri"`
		Scope        string `json:"scope"`
		State        string `json:"state"`
	}

	AccountTypeInqReq struct {
		AccountID string `json:"accountId"`
	}

	AccountTypeInqResp struct {
		AccountType AccountTypeEnum `json:"accountType"`
	}

	ErrorResp struct {
		Status       string     `json:"status"`
		Timestamp    string     `json:"timestamp"`
		Message      string     `json:"message"`
		DebugMessage string     `json:"debugMessage"`
		SubErrors    []SubError `json:"subErrors"`
	}
	SubError struct {
		Type string `json:"type"`
		Desc string `json:"desc"`
		Code string `json:"code"`
	}

	OAuthReponse struct {
		ClientID     string `json:"clientId"`
		ResponseType string `json:"responseType"`
		RedirectUri  string `json:"redirectUri"`
		Scope        string `json:"scope"`
		State        string `json:"state"`
		RequestID    string `json:"requestId"`
		Url          string `json:"url"`
		XsrvID       string `json:"xsrvId"`
	}

	AccountBalcInqReq struct {
		AccountID  string `json:"accountId"`
		RegisterNo string `json:"registerNo"`
	}

	AccountBalcInqResp struct {
		AccountID   string    `json:"accountId"`
		AccountName string    `json:"accountName"`
		Currency    string    `json:"currency"`
		BalanceLL   []Balance `json:"balanceLL"`
	}
	Balance struct {
		Type   string     `json:"type"`
		Amount AmountType `json:"amount"`
	}
	AmountType struct {
		Value    float64 `json:"value"`
		Currency string  `json:"currency"`
	}

	// OAuthReponse struct {
	// 	ClientID     string `json:"clientId"`     // Тус байгууллагыг таних зорилготой дахин давдагдашгүй дугаар
	// 	ResponseType string `json:"responseType"` // “Access Grant Response” буцаалтаар ирсэн code параметрийн утга (grant code).
	// 	RedirectUri  string `json:"redirectUri"`  // Тус байгууллагын Буцах URL байна.Энэхүү URL –ийн банканд өгч бүртгүүлэх шаардлагатай.
	// 	Scope        string `json:"scope"`        // Тус хүсэлтийг encrypt хийсэн утга
	// 	State        string `json:"state"`        // Дахин давтагдашгүй дугаар бөгөөд тус хүсэлтийг илтгэнэ.
	// }
	RateReq struct {
		Currency string `json:"currency"` // Валют
	}
	RateResp struct {
		RequestID  string     `json:"requestId"`
		Date       string     `json:"date"`
		Sequence   int        `json:"sequence"`
		Currencies []RateData `json:"currencies"`
	}
	RateData struct {
		RequestID        string `json:"requestId"`
		CurrencyCode     string `json:"currencyCode"`     // Валют код
		CurrencyName     string `json:"currencyName"`     // Валют нэр
		CashValueSell    string `json:"cashValueSell"`    // Бэлэн ханш авах
		CashValueBuy     string `json:"cashValueBuy"`     // Бэлэн ханш зарах
		NonCashValueSell string `json:"nonCashValueSell"` // Бэлэн бус ханш авах
		NonCashValueBuy  string `json:"nonCashValueBuy"`  // Бэлэн бус ханш зарах
	}
)
