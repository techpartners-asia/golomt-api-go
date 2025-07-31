package socialpay

type (
	UPointCheckUserInfoInput struct {
		CardNumber string `json:"card_number"`
		Mobile     string `json:"mobile"`
		PinCode    string `json:"pin_code"`
	}

	UpointProcessTransactionInput struct {
		Mobile     string `json:"mobile"`
		CardNumber string `json:"card_number"`
	}

	InvoicePhoneInput struct {
		Phone   string  `json:"phone"`
		Amount  float64 `json:"amount"`
		Invoice string  `json:"invoice"`
	}

	InvoicePhoneRequest struct {
		Phone    string `json:"phone"`
		Amount   string `json:"amount"`
		Invoice  string `json:"invoice"`
		Terminal string `json:"terminal"`
		Checksum string `json:"checksum"`
	}

	InvoiceInput struct {
		Amount  float64 `json:"amount"`
		Invoice string  `json:"invoice"`
	}

	InvoiceRequest struct {
		Amount   string `json:"amount"`
		Invoice  string `json:"invoice"`
		Terminal string `json:"terminal"`
		Checksum string `json:"checksum"`
	}

	SettlementRequest struct {
		SettlementId string `json:"settlementId"`
		Checksum     string `json:"checksum"`
		Terminal     string `json:"terminal"`
	}

	CommonResponse struct {
		Description string `json:"desc"`
		Status      string `json:"status"`
	}

	InvoiceResponse struct {
		ApprovalCode        string  `json:"approval_code"`
		Amount              float64 `json:"amount"`
		CardNumber          string  `json:"card_number"`
		ResponseDescription string  `json:"resp_desc"`
		ResponseCode        string  `json:"resp_code"`
		Terminal            string  `json:"terminal"`
		Invoice             string  `json:"invoice"`
		Checksum            string  `json:"checksum"`
	}

	SettlementResponse struct {
		Amount float64 `json:"amount"`
		Count  int     `json:"count"`
		Status string  `json:"status"`
	}

	ErrorResponse struct {
		ErrorDescription string `json:"errorDesc"`
		ErrorType        string `json:"errorType"`
	}

	Header struct {
		Status string `json:"status"`
		Code   int    `json:"code"`
	}

	Body struct {
		Response map[string]interface{} `json:"response"`
		Error    map[string]interface{} `json:"error"`
	}

	Response struct {
		Header Header `json:"header"`
		Body   Body   `json:"body"`
	}
)
