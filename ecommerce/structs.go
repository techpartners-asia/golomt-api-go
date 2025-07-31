package ecommerce

type (
	CreateInvoiceInput struct {
		Amount         float64    `json:"amount"`         // Картаас хасалт хийх мөнгөн дүн
		TransactionID  string     `json:"transactionId"`  // Мерчантын гүйлгээний дугаар байна.
		ReturnType     ReturnType `json:"returnType"`     // Мерчантын url рүү redirect хийх method-н POST GET утга байна. MOBILE үед app-н deeplink дуудна.
		Callback       string     `json:"callback"`       // Харилцагчийн browser дээрээс callback буюу redirect хийх мерчантын url байна
		GetToken       bool       `json:"getToken"`       // TRUE үед гүйлгээ хийгдсэний дараа тухайн картыг илэрхийлэх token авна. FALSE үед token авахгүй.
		SocialDeeplink bool       `json:"socialDeeplink"` // TRUE үед socialpay deeplink авна, FALSE үед авахгүй
	}

	CreateInvoiceResponse struct {
		Invoice        string `json:"invoice"`        // Тухайн гүйлгээний дугаар дээр үүссэн нэмэмжлэхийн дугаар
		Checksum       string `json:"checksum"`       // Банк талаас үүсгэх checksum. checksum = invoice + transactionId
		TransactionID  string `json:"transactionId"`  // Мерчантын гүйлгээний дугаар байна
		SocialDeeplink string `json:"socialDeeplink"` // Socialpay deeplink байна
		Status         string `json:"status"`         // SENT=Төлбөрийн мэдээлэл илгээсэн, PENDING=Төлбөр хүлээгдэж байгаа
		ErrorDesc      string `json:"errorDesc"`      // Статус кодны тайлбар
		ErrorCode      string `json:"errorCode"`      // Гүйлгээний статус код
	}

	CreateInvoiceRequest struct {
		Amount         string `json:"amount"`         // Картаас хасалт хийх мөнгөн дүн
		Checksum       string `json:"checksum"`       // Checksum
		TransactionID  string `json:"transactionId"`  // Мерчантын гүйлгээний дугаар байна.
		ReturnType     string `json:"returnType"`     // Мерчантын url рүү redirect хийх method-н POST GET утга байна. MOBILE үед app-н deeplink дуудна.
		Callback       string `json:"callback"`       // Харилцагчийн browser дээрээс callback буюу redirect хийх мерчантын url байна
		GenerateToken  string `json:"genToken"`       // Y үед гүйлгээ хийгдсэний дараа тухайн картыг илэрхийлэх token авна. N үед token авахгүй.
		SocialDeeplink string `json:"socialDeeplink"` // Y үед socialpay deeplink авна, N үед авахгүй

	}

	InquiryResponse struct {
		Amount        string `json:"amount"`        // Картаас хасалт хийсэн мөнгөн дүн
		Bank          string `json:"bank"`          // Карт гаргагч банкны нэр
		Status        string `json:"status"`        // SENT=Төлбөрийн мэдээлэл илгээсэн, PENDING=Төлбөр хүлээгдэж байгаа
		ErrorDesc     string `json:"errorDesc"`     // Статус кодны тайлбар
		ErrorCode     string `json:"errorCode"`     // Гүйлгээний статус код
		CardHolder    string `json:"cardHolder"`    // Карт эзэмшигчийн нэр
		CardNumber    string `json:"cardNumber"`    // Гүйлгээ хийгдсэн картын дугаар
		TransactionID string `json:"transactionId"` // Мерчантын гүйлгээний дугаар байна
		Token         string `json:"token"`         // Картыг илэрхийлэх токен дугаар
		Checksum      string `json:"checksum"`      // checksum = transactionId + errorCode + amount + token. Токен үүсээгүй үед token талбарыг оролцуулахгүй байна.
	}

	InquiryRequest struct {
		Checksum      string `json:"checksum"`      // checksum = transactionId + transactionId
		TransactionID string `json:"transactionId"` // Мерчантын гүйлгээний дугаар байна
	}
	PayTokenInput struct {
		Amount        float64 `json:"amount"`        // Гүйлгээний хийх мөнгөн дүн
		TransactionID string  `json:"transactionId"` // Мерчантын гүйлгээний дугаар байна
		Token         string  `json:"token"`         // Картын мэдээллийг илэрхийлэх банкнаас явуулсан утга
		Lang          Lang    `json:"lang"`          // Гүйлгээний дэлгэрэнгүйг авах хэл. MN Монгол хэл, EN Англи хэл
	}
	PayTokenPaymentRequest struct {
		Amount        string `json:"amount"`        // Гүйлгээний хийх мөнгөн дүн
		Checksum      string `json:"checksum"`      // checksum = amount + transactionId + token
		TransactionID string `json:"transactionId"` // Мерчантын гүйлгээний дугаар байна
		Token         string `json:"token"`         // Картын мэдээллийг илэрхийлэх банкнаас явуулсан утга
		Lang          string `json:"lang"`          // Гүйлгээний дэлгэрэнгүйг авах хэл. MN Монгол хэл, EN Англи хэл
	}

	PayTokenPaymentResponse struct {
		Amount        string `json:"amount"`        // Картаас хасалт хийсэн мөнгөн дүн
		ErrorDesc     string `json:"errorDesc"`     // Статус кодны тайлбар
		ErrorCode     string `json:"errorCode"`     // Гүйлгээний статус код
		TransactionID string `json:"transactionId"` // Мерчантын гүйлгээний дугаар байна
		Checksum      string `json:"checksum"`      // checksum = transactionId + errorCode + amount
		CardNumber    string `json:"cardNumber"`    // Гүйлгээ хийгдсэн картын дугаар масктай ирнэ
	}

	CheckTokenPaymentRequest struct {
		Checksum      string `json:"checksum"`
		TransactionID string `json:"transactionId"`
	}

	CheckTokenPaymentResponse struct {
		Amount        string `json:"amount"`        // Картаас хасалт хийсэн мөнгөн дүн
		ErrorDesc     string `json:"errorDesc"`     // Статус кодны тайлбар
		ErrorCode     string `json:"errorCode"`     // Гүйлгээний статус код
		TransactionID string `json:"transactionId"` // Мерчантын гүйлгээний дугаар
		Checksum      string `json:"checksum"`      // checksum = transactionId + errorCode + amount + token. Токен үүсээгүй үед token талбарыг оролцуулахгүй байна.
		CardNumber    string `json:"cardNumber"`    // Гүйлгээ хийгдсэн картын дугаар
		Token         string `json:"token"`         // Картыг илэрхийлэх токен дугаар
	}

	CreateTokenInput struct {
		TransactionID string     `json:"transactionId"` // Мерчантын гүйлгээний дугаар байна
		ReturnType    ReturnType `json:"returnType"`    // Мерчантын url рүү redirect хийх method-н POST GET утга байна. MOBILE үед app-н deeplink дуудна.
		Callback      string     `json:"callback"`      // Харилцагчийн browser дээрээс callback буюу redirect хийх мерчантын url байна
	}

	CreateTokenRequest struct {
		TransactionID string `json:"transactionId"` // Мерчантын гүйлгээний дугаар байна
		Checksum      string `json:"checksum"`      // Checksum
		ReturnType    string `json:"returnType"`    // Мерчантын url рүү redirect хийх method-н POST GET утга байна. MOBILE үед app-н deeplink дуудна.
		Callback      string `json:"callback"`      // Харилцагчийн browser дээрээс callback буюу redirect хийх мерчантын url байна
	}

	CreateTokenResponse struct {
		Checksum      string `json:"checksum"`      // Банк талаас үүсгэх checksum. checksum = invoice + transactionId
		TransactionID string `json:"transactionId"` // Мерчантын гүйлгээний дугаар байна
		Invoice       string `json:"invoice"`       // Тухайн гүйлгээний дугаар дээр үүссэн нэмэмжлэхийн дугаар
	}

	GetInvoiceInput struct {
		Invoice       string        `json:"invoice"`       // Тухайн гүйлгээний дугаар дээр банкнаас үүсгэж өгсөн нэмэмжлэхийн дугаар
		Lang          Lang          `json:"lang"`          // Гүйлгээний дэлгэрэнгүйг авах хэл. MN Монгол хэл, EN Англи хэл
		PaymentMethod PaymentMethod `json:"paymentMethod"` // Гүйлгээний төрөл. "payment" = Картаар төлөх, "socialpay" = SocialPay-ээр төлөх
	}

	GetConfirmationInput struct {
		Invoice string `json:"invoice"` // Тухайн гүйлгээний дугаар дээр банкнаас үүсгэж өгсөн нэмэмжлэхийн дугаар
		Lang    Lang   `json:"lang"`    // Гүйлгээний дэлгэрэнгүйг авах хэл. MN Монгол хэл, EN Англи хэл
	}

	CheckTokenRequest struct {
		Checksum      string `json:"checksum"`
		TransactionID string `json:"transactionId"`
	}

	CheckTokenResponse struct {
		BankCode      string `json:"bankCode"`      // Карт гаргагч банкны дугаар
		Bank          string `json:"bank"`          // Карт гаргагч банкны нэр
		ErrorDesc     string `json:"errorDesc"`     // Статус кодны тайлбар
		Checksum      string `json:"checksum"`      // Checksum
		ErrorCode     string `json:"errorCode"`     // Гүйлгээний статус код
		Cardholder    string `json:"cardHolder"`    // Карт эзэмшигчийн нэр
		TransactionID string `json:"transactionId"` // Мерчантын гүйлгээний дугаар байна
		CardNumber    string `json:"cardNumber"`    // Гүйлгээ хийгдсэн картын дугаар масктай ирнэ
		Token         string `json:"token"`         // Картыг илэрхийлэх токен дугаар
	}

	GetSettlementDetailsInput struct {
		Start string `json:"start"`
		End   string `json:"end"`
		Page  int    `json:"page"`
		Size  int    `json:"size"`
	}

	GetSettlementDetailsRequest struct {
		Checksum   string `json:"checksum"`   // checksum = merchantId + tokenId
		Token      string `json:"token"`      // Картын мэдээллийг илэрхийлэх банкнаас явуулсан утга.
		MerchantId string `json:"merchantId"` // Мерчантын дугаар байна.
	}

	GetSettlementDetailsResponse struct {
		Content  []SettlementDetailContent `json:"content"`
		Pageable Pageable                  `json:"pageable"`
	}

	Pageable struct {
		First         bool `json:"first"`
		Last          bool `json:"last"`
		TotalPages    int  `json:"totalPages"`
		TotalElements int  `json:"totalElements"`
	}

	SettlementDetailContent struct {
		ID              int    `json:"id"`
		MerchantId      string `json:"merchantId"`
		TerminalId      string `json:"terminalId"`
		BatchNumber     string `json:"batchNumber"`
		InvoiceNumber   string `json:"invoiceNumber"`
		CardNumber      string `json:"cardNumber"`
		Token           string `json:"token"`
		Amount          string `json:"amount"`
		MerchantInvoice string `json:"merchantInvoice"`
		RespCode        string `json:"respCode"`
		RespDesc        string `json:"respDesc"`
		TranTime        string `json:"tranTime"`
		TranDate        string `json:"tranDate"`
		SysReference    string `json:"sysReference"`
		ApprovalCode    string `json:"approvalCode"`
		FreeText1       string `json:"freeText1"`
		FreeText2       string `json:"freeText2"`
		FreeText3       string `json:"freeText3"`
		RequestDate     string `json:"requestDate"`
		ReversedDate    string `json:"reversedDate"`
		CheckedDate     string `json:"checkedDate"`
		TransferStatus  string `json:"transferStatus"`
	}

	Error struct {
		Timestamp string `json:"timestamp"`
		Status    int    `json:"status"`
		Error     string `json:"error"`
		Message   string `json:"message"`
		Path      string `json:"path"`
	}
)
