package model

import "encoding/json"

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
		CorporateRegisterNo string `json:"corpRegisterNo" validate:"required"`
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
		// Картын дугаар /маскласан/
		CardNumber string `json:"cardNumber"`
		// Картын нэр
		EmbossName string `json:"embossName"`
		// Картын төлөв
		Status string `json:"status"`
		// Картын дуусах хугацаа (yyyyMM)
		ExpiryDate string `json:"expiryDate"`
		// Картын пластикийн тайлбар
		ProductGroupDescription string `json:"productGroupDescription"`
		// Сүүлд төлбөр төлсөн огноо
		LastDueDate string `json:"lastDueDate"`
		// Сүүлд төлсөн мөнгөн дүн
		LastDueAmount json.Number `json:"lastDueAmount"`
		// Төлөлт хийх боломжтой бага дүн
		MinimumPaymentDueAmount json.Number `json:"minimumPaymentDueAmount"`
		// Картын боломжит үлдэгдэл
		AccountAvailableLimit json.Number `json:"accountAvailableLimit"`
		// Картын зарцуулалт
		AccountOutstandingBalance json.Number `json:"accountOutstandingBalance"`
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

	CardListReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	CardListData struct {
		// Картын токен. Бусад хүсэлт дээр энэхүү токенийг ашиглана
		CardToken string `json:"cardToken"`
		// Картын дугаар /маскласан/
		CardNumber string `json:"cardNumber"`
		// Картын пластикийн төрөл
		Brand string `json:"brand"`
		// Картын төрөл. DEBIT, CREDIT
		Type string `json:"type"`
	}

	UnionPayQRReq struct {
		// Wallet id
		WalletID string `json:"walletId" validate:"required"`
		// Төхөөрөмжийн Id
		DeviceID string `json:"deviceID" validate:"required"`
		// Токен
		Token string `json:"token" validate:"required"`
	}
	UnionPayQRResp struct {
		// Бар код
		BarcodeCpqrcPayload string `json:"barcodeCpqrcPayload"`
		// QR код
		EmvCpqrcPayload string `json:"emvCpqrcPayload"`
		// Хариу код. "00" амжилттай
		ResponseCode string `json:"responseCode"`
		// Хариу мессеж
		ResponseMsg string `json:"responseMsg"`
	}

	UnionPayTokenCreateReq struct {
		// Wallet id
		WalletID string `json:"walletId" validate:"required"`
		// Төхөөрөмжийн Id
		DeviceID string `json:"deviceID" validate:"required"`
		// Голомт банкны картын токен
		Token string `json:"token" validate:"required"`
		// Утасны дугаар
		MobileNumber string `json:"mobileNumber" validate:"required"`
	}
	UnionPayTokenCreateResp struct {
		// Тодорхойлолт
		Par string `json:"par"`
		// UPI токен нууцалсан талбар
		MaskedToken string `json:"maskedToken"`
		// Картын дугаар
		MaskedPan string `json:"maskedPan"`
		// UPI токены хугацаа
		TokenExpiry string `json:"tokenExpiry"`
		// UPI токены статус
		TokenState string `json:"tokenState"`
		// Төхөөрөмжийн Id
		DeviceID string `json:"deviceId"`
		// UPI токен нууцлалтгүй талбар
		Token string `json:"token"`
		// Хариу код. "00" амжилттай
		ResponseCode string `json:"responseCode"`
		// Хариу мессеж. "Approved"
		ResponseMsg    string `json:"responseMsg"`
		Signature      string `json:"signature"`
		UmpsSignCertID string `json:"umpsSignCertId"`
		// Хүсэлтийн төрөл
		MsgType string `json:"msgType"`
	}

	UnionPayTokenUpdateReq struct {
		// Wallet id
		WalletID string `json:"walletId" validate:"required"`
		// Төхөөрөмжийн Id
		DeviceID string `json:"deviceID" validate:"required"`
		// UPI токен
		UpiToken string `json:"upiToken" validate:"required"`
		// UPI токены статус. Жишээ: ACTIVE
		TokenAction string `json:"tokenAction,omitempty"`
		// Утасны дугаар
		MobileNumber string `json:"mobileNumber" validate:"required"`
		// Голомт банкны картын токен
		CardToken string `json:"cardToken" validate:"required"`
	}
	UnionPayTokenUpdateResp struct {
		// UPI токены статус
		TokenState string `json:"tokenState"`
		// Төхөөрөмжийн Id
		DeviceID string `json:"deviceId"`
		// Хариу код. "00" амжилттай
		ResponseCode string `json:"responseCode"`
		// Хариу мессеж
		ResponseMsg string `json:"responseMsg"`
		// UPI токен
		Token          string `json:"token"`
		Signature      string `json:"signature"`
		UmpsSignCertID string `json:"umpsSignCertId"`
		// Хүсэлтийн төрөл
		MsgType string `json:"msgType"`
	}
)

// Картын захиалгад ашиглагдах нийтлэг бүтэц
type (
	CardPhoneEmail struct {
		// Төрөл. PHONE – утасны дугаар, EMAIL – и-мэйл хаяг
		Type string `json:"type" validate:"required"`
		// Дэд төрөл. EMAIL бол HOMEEML, PHONE бол CELLPH
		SubType string `json:"subType" validate:"required"`
		// И-мэйл хаяг. type == EMAIL үед заавал
		Email string `json:"email,omitempty"`
		// Утасны дугаар. type == PHONE үед заавал
		Phone string `json:"phone,omitempty"`
		// Улсын код. Жишээ: 976
		CountryCode string `json:"countryCode" validate:"required"`
	}
	CardAddress struct {
		// Хаягийн төрөл. HOME, WORK, DELIVERY
		Type string `json:"type" validate:"required"`
		// Улс. Жишээ: MN
		Country string `json:"country" validate:"required"`
		// Хот, аймаг
		State string `json:"state" validate:"required"`
		// Сум, дүүрэг
		City string `json:"city" validate:"required"`
		// Хорооны дугаар
		SubDistrict string `json:"subDistrict" validate:"required"`
		// Гудамжны нэр
		StreetName string `json:"streetName" validate:"required"`
		// Хотхоны нэр
		Town string `json:"town,omitempty"`
		// Байрны дугаар
		Apartment string `json:"apartment,omitempty"`
		// Орцны дугаар
		Entry string `json:"entry,omitempty"`
		// Тоотын дугаар
		DoorNo string `json:"doorNo,omitempty"`
		// Хаягийн мэдээлэл 1. Формат: хороо,гудамж,байр,тоот
		AddressLine1 string `json:"addressLine1" validate:"required"`
		// Хаягийн мэдээлэл 2
		AddressLine2 string `json:"addressLine2,omitempty"`
		// Хаягийн мэдээлэл 3
		AddressLine3 string `json:"addressLine3,omitempty"`
	}
	CardPostRequest struct {
		// Хүргэлтээр авах эсэх. Y, N
		IsByPost string `json:"isByPost" validate:"required"`
		// Хүргэлтийн төлбөр төлсөн эсэх. Y, N
		IsByPostalFee string `json:"isByPostalFee" validate:"required"`
		// Хүргэлтийн шимтгэл
		PostalFee float64 `json:"postalFee"`
	}
	// Нийтлэг message хариу
	CardMessageResp struct {
		// Хариу мессеж. Жишээ: SUCCESS: 4628********2570 status changed successfully
		Message string `json:"message"`
	}
	// Нийтлэг status, message хариу
	CardStatusResp struct {
		// Төлөв. SUCCESS – амжилттай
		Status string `json:"status"`
		// Тайлбар
		Message string `json:"message"`
	}
)

type (
	// 9.4. Кредит карт захиалах (Pre-approved virtual credit card)
	CardCreditOrderReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Бүтээгдэхүүний код
		ProdCode string `json:"prodCode" validate:"required"`
		// Хүсэж буй картын эрх
		Amount float64 `json:"amount" validate:"required"`
		// Өөрийн нэр
		FirstName string `json:"firstName" validate:"required"`
		// Эцэг/эхийн нэр
		LastName string `json:"lastName" validate:"required"`
		// Төрсөн огноо. yyyy-MM-dd
		Birthdate string `json:"birthdate" validate:"required"`
		// Холбоо барих мэдээлэл
		PhoneEmail []CardPhoneEmail `json:"phoneEmail" validate:"required"`
	}
	CardOrderResp struct {
		// Картын токен
		CardToken string `json:"cardToken"`
	}

	// 9.5. Дебит карт захиалах
	CardDebitOrderReq struct {
		// Данс нээлгэхийг зөвшөөрсөн эсэх. Y, N
		AcctRegFlg string `json:"acctRegFlg" validate:"required"`
		// Үйлчилгээний нөхцөл зөвшөөрсөн эсэх. Y, N
		ContFlg string `json:"contFlg" validate:"required"`
		// Бүтээгдэхүүний код
		Product string `json:"product" validate:"required"`
		// Салбарын дугаар
		BranchID string `json:"branchId" validate:"required"`
		// Харилцагчийн регистрийн дугаар
		RegNo string `json:"regNo" validate:"required"`
		// Дансны дугаар
		AccountNumber string `json:"accountNumber,omitempty"`
		// Шимтгэл төлөх данс
		FeeAccount string `json:"feeAccount,omitempty"`
		// Гуравдагч системийн токен
		Token string `json:"token,omitempty"`
		// Карт хүргэлтийн мэдээлэл
		PostRequest *CardPostRequest `json:"postRequest,omitempty"`
	}
	CardDebitOrderResp struct {
		// Үр дүнгийн мессеж (лавлах дугаартай)
		ResultMsg string `json:"resultMsg"`
		// Амжилттай бол SUCCESS
		Status string `json:"status"`
		// Картын токен
		CardToken string `json:"cardToken"`
		// Шинээр нээсэн харилцах данс
		AccountID string `json:"accountId"`
	}

	// 9.6. Их сургуулийн дебит карт захиалах
	CardUniversityDebitOrderReq struct {
		// Данс нээлгэхийг зөвшөөрсөн эсэх. Y, N
		AcctRegFlg string `json:"acctRegFlg" validate:"required"`
		// Үйлчилгээний нөхцөл зөвшөөрсөн эсэх. Y, N
		ContFlg string `json:"contFlg" validate:"required"`
		// Бүтээгдэхүүний код
		Product string `json:"product" validate:"required"`
		// Салбарын дугаар
		BranchID string `json:"branchId" validate:"required"`
		// Харилцагчийн регистрийн дугаар
		RegNo string `json:"regNo" validate:"required"`
		// Дансны дугаар
		AccountNumber string `json:"accountNumber,omitempty"`
		// Шимтгэл төлөх данс
		FeeAccount string `json:"feeAccount,omitempty"`
		// Гуравдагч системийн токен
		Token string `json:"token,omitempty"`
		// Их сургуулийн нэр
		UniversityName string `json:"universityName" validate:"required"`
		// Мэргэжил
		Profession string `json:"profession" validate:"required"`
		// Оюутны код
		StudentID string `json:"studentId" validate:"required"`
		// Base64 цээж зураг
		ProfilePhoto string `json:"profilePhoto" validate:"required"`
		// Карт хүргэлтийн мэдээлэл
		PostRequest *CardPostRequest `json:"postRequest,omitempty"`
	}
	CardUniversityDebitOrderResp struct {
		// Үр дүнгийн мессеж (лавлах дугаартай)
		ResultMsg string `json:"resultMsg"`
		// Амжилттай бол SUCCESS
		Status string `json:"status"`
		// Картын токен
		Token string `json:"token"`
		// Шинээр нээсэн харилцах данс
		AccountID string `json:"accountId"`
		// Маскласан картын дугаар
		CardNumber string `json:"cardNumber"`
	}

	// 9.7. Prepaid карт захиалах
	CardPrepaidOrderReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Өөрийн нэр
		FirstName string `json:"firstName" validate:"required"`
		// Эцэг/эхийн нэр
		LastName string `json:"lastName" validate:"required"`
		// Төрсөн огноо. yyyy-MM-dd
		Birthdate string `json:"birthdate" validate:"required"`
		// Гуравдагчийн unique id (Голомт банкны хэрэглэгч биш үед)
		CorporateID string `json:"corporateId,omitempty"`
		// Хэрэглэгчийн unique id (Голомт банкны хэрэглэгч биш үед)
		UserID string `json:"userId,omitempty"`
		// Картын мөнгөн дүн
		CreditLimit float64 `json:"creditLimit" validate:"required"`
		// Картын харьяалагдах салбар
		BranchID string `json:"branchId,omitempty"`
		// Холбоо барих мэдээлэл
		PhoneEmail []CardPhoneEmail `json:"phoneEmail" validate:"required"`
		// Хаягийн мэдээлэл
		Address []CardAddress `json:"address" validate:"required"`
	}

	// 9.8. Prepaid карт цэнэглэлт
	CardPrepaidTopupReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Цэнэглэлтийн гүйлгээ гаргах данс
		AccountNo string `json:"accountNo" validate:"required"`
		// Цэнэглэх дүн
		Amount float64 `json:"amount" validate:"required"`
		// Валют
		Currency string `json:"currency" validate:"required"`
		// Гүйлгээний утга
		Remarks string `json:"remarks" validate:"required"`
		// SecretKey-ээр үүсгэсэн 6 оронтой код
		ApproveCode string `json:"approveCode" validate:"required"`
	}
	CardPrepaidTopupResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// SUCCESS, FAILED
		Status string `json:"status"`
		// Гүйлгээний лавлах дугаартай тайлбар
		Message string `json:"message"`
	}

	// 9.10. Хүүхдийн карт захиалах
	CardChildOrderReq struct {
		// Хүүхдийн эцэг/эхийн нэр
		LastName string `json:"lastName" validate:"required"`
		// Хүүхдийн өөрийн нэр
		FirstName string `json:"firstName" validate:"required"`
		// Хүүхдийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Хүйс. M, F
		CustGender string `json:"custGender" validate:"required"`
		// Сургуулийн нэр
		OrgName string `json:"orgName" validate:"required"`
		// Сургуулийн хаяг
		OrgAddress string `json:"orgAddress" validate:"required"`
		// Ажил эрхлэлт. Хүүхэд бол СУРАГЧ
		Occupation string `json:"occupation" validate:"required"`
		// Үйлчилгээний нөхцөл зөвшөөрсөн эсэх. Y, N
		ContFlg string `json:"contFlg" validate:"required"`
		// Захиалах картын мэдээлэл
		Card CardChildOrderCard `json:"card" validate:"required"`
		// Эцэг/эх, асран хамгаалагчийн мэдээлэл
		Relationships []CardChildRelationship `json:"relationships" validate:"required"`
		// Холбоо барих мэдээлэл
		PhoneEmails []CardChildPhoneEmail `json:"phoneEmails,omitempty"`
		// Хаягийн мэдээлэл. deliveryFlg = Y бол DELIVERY төрөлтэй хаяг заавал
		Address []CardAddress `json:"address,omitempty"`
	}
	CardChildOrderCard struct {
		// Картын бүтээгдэхүүн
		Product string `json:"product" validate:"required"`
		// Салбарын код
		Branch string `json:"branch" validate:"required"`
		// Шимтгэл төлөх данс
		FeeAcctNo string `json:"feeAcctNo,omitempty"`
		// Картын валют. MNT
		Currency string `json:"currency" validate:"required"`
		// Карт дээр бичигдэх нэр
		EmbossName string `json:"embossName" validate:"required"`
		// Биет карт эсэх. Y, N
		PlasticFlg string `json:"plasticFlg" validate:"required"`
		// Хүргэлтээр авах эсэх. Y, N
		DeliveryFlg string `json:"deliveryFlg" validate:"required"`
	}
	CardChildRelationship struct {
		// Асран хамгаалагчийн эцэг/эхийн нэр
		LastName string `json:"lastName" validate:"required"`
		// Асран хамгаалагчийн өөрийн нэр
		FirstName string `json:"firstName" validate:"required"`
		// Асран хамгаалагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Асран хамгаалагчийн утас
		PrimaryPhone string `json:"primaryPhone" validate:"required"`
		// Асран хамгаалагчийн и-мэйл
		PrimaryEmail string `json:"primaryEmail" validate:"required"`
	}
	CardChildPhoneEmail struct {
		// EMAIL, PHONE
		Type string `json:"type" validate:"required"`
		// type == PHONE үед заавал
		Phone string `json:"phone,omitempty"`
		// type == EMAIL үед заавал
		Email string `json:"email,omitempty"`
		// Улсын код. Жишээ: 976
		CountryLocalCode string `json:"countryLocalCode" validate:"required"`
	}
	CardChildOrderResp struct {
		// Лавлах дугаар
		ReferenceID string `json:"referenceId"`
	}

	// 9.12. Кредит карт төлөв солих
	CardCreditStatusChangeReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Төлөв. A – active, L – lost
		Status string `json:"status" validate:"required"`
		// Шалтгаан
		Reason string `json:"reason,omitempty"`
	}

	// 9.13. Карт идэвхжүүлэх
	CardActivateReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
	}

	// 9.14. Кредит картын орлого хийх
	CardCreditPaymentReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Гүйлгээний дүн
		Amount float64 `json:"amount" validate:"required"`
		// Валют
		Currency string `json:"currency" validate:"required"`
		// Зарлага гаргах данс
		AccountNo string `json:"accountNo" validate:"required"`
		// Гүйлгээний утга
		Remarks string `json:"remarks" validate:"required"`
		// SecretKey-ээр үүсгэсэн 6 оронтой код
		ApproveCode string `json:"approveCode" validate:"required"`
	}

	// 9.16. Void transaction, 9.17. Байгууллагын картын гүйлгээ буцаах
	CardVoidReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Гүйлгээний дүн
		Amount float64 `json:"amount" validate:"required"`
		// Валют
		CurrencyCode string `json:"crncyCode" validate:"required"`
		// Терминал дугаар
		TerminalID string `json:"terminalId" validate:"required"`
		// Purchase хариуны approval code
		TransactionApprovalCode string `json:"tranApprovalCode" validate:"required"`
		// Purchase хариуны reference code
		TransactionReferenceCode string `json:"tranReferenceCode" validate:"required"`
	}
	CardVoidResp struct {
		// Гүйлгээний огноо
		PurchaseDate string `json:"prchDate"`
		// Гүйлгээний төлөв
		PurchaseStatus string `json:"prchStatus"`
	}
	CardCorpVoidResp struct {
		// Гүйлгээний төлөв. SUCCESS
		VoidStatus string `json:"voidStatus"`
	}

	// 9.18. Нэхэмжлэх төлөх
	CardInvoicePayReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Гүйлгээний дүн
		Amount float64 `json:"amount" validate:"required"`
		// Group id
		GroupID string `json:"groupId" validate:"required"`
		// Терминал дугаар
		TerminalID string `json:"terminalId" validate:"required"`
		// Мерчантын гүйлгээний дугаар
		TranID string `json:"tranId" validate:"required"`
		// Мерчантын дугаар
		PID string `json:"pid" validate:"required"`
		// Утасны дугаар
		Phone string `json:"phone" validate:"required"`
	}
	CardInvoicePayResp struct {
		// Гүйлгээний лавлах дугаар
		TranRef string `json:"tranRef"`
		// Гүйлгээний дүн
		TxnAmount json.Number `json:"txnAmount"`
		// Тогтмол ""
		Url string `json:"url"`
		// Тогтмол "0"
		GetOrPost string `json:"getOrPost"`
		// Маскласан картын дугаар
		CardNumber string `json:"cardNumber"`
		// Урамшууллын дугаар
		Lottery string `json:"lottery"`
		// SocialPay QR
		QrData string `json:"qrData"`
		// Coffee shop терминал эсэх
		CoffeeFlg string `json:"coffeeFlg"`
	}

	// 9.19. Карт токенжуулах
	CardTokenizeCustomerReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	CardTokenizeCustomerResp struct {
		// Байгууллагын client id
		ClientID string `json:"clientId"`
		// Grant code
		ResponseType string `json:"responseType"`
		// Карт токенжуулалт хийх URL
		RedirectUri string `json:"redirectUri"`
		// Encrypt хийсэн хүсэлт
		Scope string `json:"scope"`
		// Хүсэлтийн давтагдашгүй дугаар
		State string `json:"state"`
	}

	// 9.20. Мерчант бүртгэх
	CardMerchantAddReq struct {
		// Мерчантын регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Төлөлт хийх банкны код
		PaymentBankCode string `json:"paymentBankCode" validate:"required"`
		// Хуулга үүсгэх төрөл. H, S, B
		StmtGenType string `json:"stmtGenType" validate:"required"`
		// Хуулга илгээх и-мэйл
		StmtEmail string `json:"stmtEmail" validate:"required"`
		// Industry code
		IndustryCode string `json:"industryCode" validate:"required"`
		// Category code
		CategoryCode string `json:"categoryCode" validate:"required"`
		// Мерчантын төрөл. I – Individual, G – group
		Type string `json:"type" validate:"required"`
		// Байгууллагын харилцах данс
		AccountID string `json:"accountId" validate:"required"`
		// MDR group. Лавлах төрөл: MDR_GROUP
		MerchantMDRGroupID string `json:"merchantMDRgroupId" validate:"required"`
		// Терминалууд
		Terminals []CardMerchantTerminal `json:"terminals,omitempty"`
		// Холбоо барих мэдээлэл
		Contacts []CardMerchantContact `json:"contacts" validate:"required"`
	}
	CardMerchantTerminal struct {
		// Terminal id
		ID string `json:"id" validate:"required"`
		// Терминалын төрөл
		Type string `json:"type" validate:"required"`
	}
	CardMerchantContact struct {
		// Contact төрөл. MAIN
		Type string `json:"type" validate:"required"`
		// Регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Хэл
		Language string `json:"language,omitempty"`
		// Хүйс. F, M
		Gender string `json:"gender,omitempty"`
		// Нэр
		Name string `json:"name,omitempty"`
		// Байршлын мэдээлэл
		Addresses []CardMerchantAddress `json:"addresses" validate:"required"`
	}
	CardMerchantAddress struct {
		// Улс
		Country string `json:"country" validate:"required"`
		// Хот, аймаг
		State string `json:"state" validate:"required"`
		// Сум, дүүрэг
		City string `json:"city" validate:"required"`
		// Баг, хороо
		SubDistrict string `json:"subDistrict" validate:"required"`
		// Хаяг
		AddressLine1 string `json:"addressLine1" validate:"required"`
		// Байршлын төрөл. HOME, WORK, DELIVERY, MAIN, TEMP
		Type string `json:"type" validate:"required"`
	}
	CardMerchantAddResp struct {
		// Мерчантын дугаар
		MerchantID string `json:"merchantId"`
	}

	// 9.21. Терминал бүртгэх
	CardTerminalAddReq struct {
		// Мерчантын регистрийн дугаар
		MerchantRegisterNo string `json:"merchantRegisterNo" validate:"required"`
		// Мерчантын ID
		MerchantID string `json:"merchantId" validate:"required"`
		// Терминалын serial no
		SerialNo string `json:"serialNo,omitempty"`
		// Терминалын байршил
		Location string `json:"location,omitempty"`
	}
	CardTerminalAddResp struct {
		// Терминал дугаар
		TerminalID string `json:"terminalId"`
		// Мессеж код
		MsgCode json.Number `json:"msgCode"`
		// Мессеж тайлбар
		MsgDesc string `json:"msgDesc"`
	}

	// 9.22. Reversal гүйлгээ
	CardReversalReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын токен
		Token string `json:"token" validate:"required"`
		// Гүйлгээний дүн
		Amount float64 `json:"amount" validate:"required"`
		// Мерчантын гүйлгээний дугаар
		TransactionID string `json:"transactionId" validate:"required"`
		// Хэл. MN, EN
		Lang string `json:"lang" validate:"required"`
		// Гүйлгээний төрөл. REVERSAL
		TransactionType string `json:"transactionType" validate:"required"`
	}
	CardReversalResp struct {
		// Process code
		ProcessCode string `json:"processCode"`
		// Trace number
		TraceNumber string `json:"traceNumber"`
		// Гүйлгээний цаг. HHMMSS
		TransactionTime string `json:"transactionTime"`
		// Гүйлгээний огноо. MMDD
		TransactionDate string `json:"transactionDate"`
		// NII
		NII string `json:"NII"`
		// Reference number
		SystemReferenceNumber string `json:"systemReferenceNumber"`
		// Auth code
		ApprovalCode string `json:"approvalCode"`
		// Хариу код
		TransactionRespondCode string `json:"transactionRespondCode"`
		// Терминал дугаар
		TerminalID string `json:"terminalId"`
	}

	// 9.24. Терминалын жагсаалт авах
	CardTerminalListReq struct {
		// Байгууллагын регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Мерчантын дугаар
		MerchantID string `json:"merchantId" validate:"required"`
	}
	CardTerminalListResp struct {
		// Терминалын жагсаалт
		TerminalDetails []CardTerminalDetail `json:"terminalDetails"`
	}
	CardTerminalDetail struct {
		// Терминал дугаар
		TerminalID string `json:"terminalId"`
		// Serial дугаар
		SerialNo string `json:"serialNo"`
		// Brand
		Brand string `json:"brand"`
		// Model
		Model string `json:"model"`
		// Салбар
		Branch string `json:"branch"`
		// Master key
		MasterKey string `json:"masterKey"`
		// Derivation key
		DerivationKey string `json:"derivationKey"`
		// Байршил
		Location string `json:"location"`
		// Улсын код
		CountryCode string `json:"countryCode"`
		// Валют
		Currency string `json:"currency"`
		// Risk profile
		RiskProfile string `json:"riskProfile"`
		// Төлөв
		Status string `json:"status"`
		// Эхлэх огноо
		StartDate string `json:"startDate"`
		// Дуусах огноо
		EndDate string `json:"endDate"`
		// Master key check value
		MkeyCheckVal string `json:"mkeyCheckVal"`
		// Derivation key check value
		DkeyCheckVal string `json:"dkeyCheckVal"`
	}

	// 9.25. Картын мэдээлэл Openbank web дээр харах
	CardWebViewReq struct {
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Хэрэглэгчийн IP хаяг
		CustomerIpAddress string `json:"customerIpAddress" validate:"required"`
	}
	CardWebViewResp struct {
		// Картын мэдээлэл харах URL
		Url string `json:"url"`
	}

	// 9.26. Харилцагч дээр тухайн product-тай карт байгаа эсэх
	CardProductCheckReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Картын бүтээгдэхүүний код
		ProdCode string `json:"prodCode" validate:"required"`
	}
	CardProductCheckResp struct {
		// Карт байгаа эсэх. Y, N
		HasCard string `json:"hasCard"`
	}

	// 9.27. Картын token replace хийх
	CardTokenReplaceReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Одоо ашиглаж буй токен
		OldToken string `json:"oldToken" validate:"required"`
		// Шинэ картын токен
		NewToken string `json:"newToken" validate:"required"`
	}

	// 9.28. Байгууллага токентэй гүйлгээ хийх
	CardCorpPurchaseReq struct {
		// Картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Гүйлгээний дүн
		Amount float64 `json:"amount" validate:"required"`
		// Валют
		CurrencyCode string `json:"crncyCode" validate:"required"`
		// Терминал дугаар
		TerminalID string `json:"terminalId" validate:"required"`
		// SecretKey-ээр үүсгэсэн 6 оронтой код
		ApproveCode string `json:"approveCode" validate:"required"`
	}
	// X-DEVICE-* header-ийн мэдээлэл
	CardDeviceInfo struct {
		// Төхөөрөмжийн дугаар (X-DEVICE-ID)
		ID string
		// Төхөөрөмжийн нэр (X-DEVICE-NAME)
		Name string
		// Төхөөрөмжийн IP хаяг (X-DEVICE-IP)
		IP string
	}

	// 9.29. Картын шилжүүлэг хийх
	CardTransferReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Гүйлгээний төрөл. OSB, INB, CRD
		TxnType string `json:"txnType" validate:"required"`
		// Шилжүүлэг хийх картын токен
		CardToken string `json:"cardToken" validate:"required"`
		// Хүлээн авагч. OSB/INB – дансны дугаар, CRD – картын токен
		Receiver string `json:"receiver" validate:"required"`
		// Хүлээн авагч банкны код (OSB үед)
		BankCode string `json:"bankCode,omitempty"`
		// Гүйлгээний дүн
		Amount CardAmount `json:"amount" validate:"required"`
		// Гүйлгээний утга
		Remarks string `json:"remarks" validate:"required"`
	}
	CardAmount struct {
		// Мөнгөн дүн
		Value float64 `json:"value" validate:"required"`
		// Валют
		Currency string `json:"currency" validate:"required"`
	}
	CardTransferResp struct {
		// Төлөв. SUCCESS – амжилттай
		Status string `json:"status"`
		// Лавлах дугаар
		RRN string `json:"rrn"`
		// Trace number
		TraceNumber string `json:"traceNumber"`
	}

	// 9.30. Карт руу орлого оруулах
	CardDepositReq struct {
		// Мерчантын дугаар
		MerchantID string `json:"merchantId" validate:"required"`
		// Гүйлгээний дүн
		Amount float64 `json:"amount" validate:"required"`
		// Шимтгэлийн дүн
		ChrgAmount float64 `json:"chrgAmount"`
		// Валют
		Currency string `json:"currency" validate:"required"`
		// Терминал дугаар
		TerminalID string `json:"terminalId" validate:"required"`
		// Картын токен
		Token string `json:"token" validate:"required"`
		// Гуравдагч системийн гүйлгээний дугаар (давтагдашгүй)
		TranID string `json:"tranId" validate:"required"`
	}
	CardDepositResp struct {
		// Гүйлгээний огноо
		PurchaseDate string `json:"prchDate"`
		// Гүйлгээний төлөв
		PurchaseStatus string `json:"prchStatus"`
	}

	// 9.31. Картын орлогын гүйлгээний төлөв шалгах
	CardDepositCheckReq struct {
		// CDINUIP-д илгээсэн гүйлгээний дугаар
		TranID string `json:"tranId" validate:"required"`
	}
	CardDepositCheckResp struct {
		// Төлөв. PENDING, SUCCESS, FAILED
		PurchaseStatus string `json:"prchStatus"`
	}
)
