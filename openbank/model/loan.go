package model

import "encoding/json"

// 7.1, 7.2 Зээлийн дансны дэлгэрэнгүй, график
type (
	LoanDetailReq struct {
		// Зээлийн дансны дугаар
		AccountID string `json:"accountId" validate:"required"`
		// Данс эзэмшигчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	LoanDetailResp struct {
		// Бүтээгдэхүүний нэр
		AccountCategory string `json:"accountCategory"`
		// Дансны нэр
		AccountName string `json:"accountName"`
		// Дансны дугаар
		AccountNumber string `json:"accountNumber"`
		// Валют
		Currency string `json:"currency"`
		// Данс нээсэн огноо
		OpenDate string `json:"openDate"`
		// Данс хаагдах хугацаа
		MaturityDate string `json:"maturityDate"`
		// Хугацаа /өдрөөр/
		PeriodInDays json.Number `json:"periodInDays"`
		// Хугацаа /сараар/
		PeriodInMonths json.Number `json:"periodInMonths"`
		// Дараагийн өр төлөх хугацаа
		NextInstallmentDueDate string `json:"nextInstallmentDueDate"`
		// Төлөх бага дүн
		OverdueAmount json.Number `json:"overdueAmount"`
		// Нийт үлдэгдэл (SPEC дээрх нэршлээр: liabliltyAmount)
		LiabilityAmount json.Number `json:"liabliltyAmount"`
	}
	// SPEC: 7.2-ын хариу нь зөвхөн `account` талбартай гэж тодорхойлогдсон
	LoanScheduleResp struct {
		// Дансны дугаар
		Account string `json:"account"`
	}
)

// 7.3, 7.5, 7.6 Зээл / кредит картын хүсэлт
type (
	LoanRequestReq struct {
		// Зээлийн мэдээлэл
		LoanInfo LoanInfo `json:"loanInfo" validate:"required"`
		// Зээлдэгчийн мэдээлэл
		Customer LoanCustomer `json:"customer" validate:"required"`
		// Зээлдэгчийн нэмэлт мэдээлэл
		Demographic LoanDemographic `json:"demographic" validate:"required"`
		// Гэрээ болон бусад төрлийн холболтын нөхцөл зөвшөөрөх
		Agreements LoanAgreements `json:"agreements" validate:"required"`
		// Гуравдагч системийн мэдээлэл
		ThirdPartyInfo LoanThirdPartyInfo `json:"thirdPartyInfo" validate:"required"`
		// Зээлдэгчийн холбоо барих мэдээлэл
		Contact []LoanContact `json:"contact" validate:"required"`
		// Зээлдэгчийн хаягийн мэдээлэл
		Address []LoanAddress `json:"address" validate:"required"`
		// Зээлийн тооцоолол хийгдсэний дараа хариу илгээх URL
		CallbackUrl string `json:"callbackUrl,omitempty"`
	}
	LoanInfo struct {
		// Тэтгэврийн зээл: Y - scoring хүсэлт, N - зээлийн хүсэлт (зөвхөн 7.6)
		IsPensionScoringRequest string `json:"isPensionScoringRequest,omitempty"`
		// Авах зээлийн дүн
		Amount float64 `json:"amount" validate:"required"`
		// Урьдчилгаа төлбөрийн дүн
		PrepaidAmt float64 `json:"prepaidAmt"`
		// Зээл авах хугацаа /сараар/
		Period int `json:"period" validate:"required"`
		// Зээлийн төрөл. LOAN, PENSION
		ProdType string `json:"prodType" validate:"required"`
		// Зээлийн бүтээгдэхүүний код. Лавлах төрөл: ONLNLN
		ProdCode string `json:"prodCode" validate:"required"`
		// Харилцагчийн бүртгэл байрших салбарын дугаар
		BranchID string `json:"branchId" validate:"required"`
		// Зээл авах валют
		LoanCrn string `json:"loanCrn" validate:"required"`
		// Зээлийн scoring шимтгэл авах данс (тэтгэврийн зээлд байхгүй)
		ChrgAccountID string `json:"chrgAccountId,omitempty"`
		// Сард төлөлт хийх давтамж. M - сард нэг, T - сард хоёр
		InstallmentType string `json:"installmentType" validate:"required"`
	}
	LoanCustomer struct {
		// Зээлдэгчийн өөрийн нэр
		FirstName string `json:"firstName" validate:"required"`
		// Зээлдэгчийн эцэг-эхийн нэр
		LastName string `json:"lastName" validate:"required"`
		// Регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	LoanDemographic struct {
		// Албан тушаалын код. Лавлах төрөл: OCCUPATION
		Appointment string `json:"appointment" validate:"required"`
		// Секторын код. Лавлах төрөл: SECTOR_CODE
		Sector string `json:"sector" validate:"required"`
		// Дэд секторын код. Лавлах төрөл: SUB_SECTOR_CODE
		SubSector string `json:"subSector,omitempty"`
		// Сургуульд орсон огноо. yyyy-MM-dd
		EnrollmentDate string `json:"enrollmentDate,omitempty"`
		// Төгссөн сургууль
		SchoolName string `json:"schoolName,omitempty"`
		// Мэргэжлийн зэрэг. Лавлах төрөл: QUALIFICATION
		Degree string `json:"degree" validate:"required"`
		// Гэрлэлтийн төлөв. Лавлах төрөл: MARITAL_STATUS
		MaritalStatus string `json:"maritalStatus" validate:"required"`
		// Ажилдаа орсон огноо. yyyy-MM-dd
		StartDate string `json:"startDate" validate:"required"`
		// Хөдөлмөр эрхэлсэн нийт жил
		YearsWork int `json:"yearsWork" validate:"required"`
	}
	LoanAgreements struct {
		// Зээлийн үйлчилгээний нөхцөл зөвшөөрч буй эсэх. Y, N
		ZmsFlg string `json:"zmsFlg" validate:"required"`
		// ДАН системээс мэдээлэл татах эсэх. Y, N
		DanFlg string `json:"danFlg" validate:"required"`
	}
	LoanThirdPartyInfo struct {
		// Урьдчилан өгөгдсөн байгууллагын дугаар
		CorpID string `json:"corpId" validate:"required"`
		// Байгууллагын дансны дугаар
		CorpAccountID string `json:"corpAccountId" validate:"required"`
		// Дансны нэр
		CorpAccountName string `json:"corpAccountName,omitempty"`
	}
	LoanContact struct {
		// EMAIL, PHONE
		Type string `json:"type" validate:"required"`
		// CELLPH, HOMEEML, WORKEML
		SubType string `json:"subType,omitempty"`
		// type == PHONE үед заавал
		Phone string `json:"phone,omitempty"`
		// type == EMAIL үед заавал
		Email string `json:"email,omitempty"`
		// Улсын код
		CountryCode string `json:"countryCode" validate:"required"`
	}
	LoanAddress struct {
		// Улс
		Country string `json:"country" validate:"required"`
		// HOME, WORK, DELIVERY (зөвхөн кредит карт)
		Type string `json:"type" validate:"required"`
		// Хот эсвэл аймаг
		State string `json:"state" validate:"required"`
		// Сум дүүрэг
		City string `json:"city" validate:"required"`
		// Хорооны дугаар
		SubDistrict string `json:"subDistrict" validate:"required"`
		// Гудамжын нэр
		StreetName string `json:"streetName" validate:"required"`
		// Хотхоны нэр
		Town string `json:"town,omitempty"`
		// Байрын дугаар
		Apartment string `json:"apartment,omitempty"`
		// Орцын дугаар
		Entry string `json:"entry,omitempty"`
		// Тоотын дугаар
		DoorNo string `json:"doorNo,omitempty"`
		// Хаяг 1. Формат: хороо,гудамж,байр,тоот
		AddressLine1 string `json:"addressLine1" validate:"required"`
		// Хаяг 2
		AddressLine2 string `json:"addressLine2,omitempty"`
		// Хаяг 3
		AddressLine3 string `json:"addressLine3,omitempty"`
	}
	// 7.3, 7.5, 7.6, 7.13, 7.14 хариу
	LoanRequestResp struct {
		// FAILED, SUCCESS
		Status string `json:"status"`
		// Мессеж
		Message string `json:"message"`
		// Хүсэлтийн лавлах дугаар
		RequestID string `json:"requestId"`
	}

	LoanCardRequestReq struct {
		// Картын мэдээлэл
		CardInfo LoanCardInfo `json:"cardInfo" validate:"required"`
		// Зээлдэгчийн мэдээлэл
		Customer LoanCustomer `json:"customer" validate:"required"`
		// Зээлдэгчийн нэмэлт мэдээлэл
		Demographic LoanDemographic `json:"demographic" validate:"required"`
		// Гэрээ болон бусад төрлийн холболтын нөхцөл зөвшөөрөх
		Agreements LoanAgreements `json:"agreements" validate:"required"`
		// Гуравдагч системийн мэдээлэл
		ThirdPartyInfo LoanThirdPartyInfo `json:"thirdPartyInfo" validate:"required"`
		// Зээлдэгчийн холбоо барих мэдээлэл
		Contact []LoanContact `json:"contact" validate:"required"`
		// Хаягийн мэдээлэл. deliveryFlg = Y үед DELIVERY төрлийн хаяг нэмнэ
		Address []LoanAddress `json:"address" validate:"required"`
		// Зээлийн тооцоолол хийгдсэний дараа хариу илгээх URL
		CallbackUrl string `json:"callbackUrl,omitempty"`
	}
	LoanCardInfo struct {
		// Авах зээлийн дүн
		Amount float64 `json:"amount" validate:"required"`
		// Урьдчилгаа төлбөрийн дүн
		PrepaidAmt float64 `json:"prepaidAmt"`
		// Зээлийн төрөл. CARD
		ProdType string `json:"prodType" validate:"required"`
		// Бүтээгдэхүүний код
		ProdCode string `json:"prodCode" validate:"required"`
		// Салбарын дугаар
		BranchID string `json:"branchId" validate:"required"`
		// Картын валют
		Currency string `json:"currency" validate:"required"`
		// Карт дээр бичигдэх нэр
		EmbossName string `json:"embossName" validate:"required"`
		// Биет карт эсэх. Y, N
		PlasticFlg string `json:"plasticFlg" validate:"required"`
		// Хүргэлтээр авах эсэх. Y, N
		DeliveryFlg string `json:"deliveryFlg,omitempty"`
		// Авто төлөлт бүртгэх эсэх. Y, N
		AutoPayFlg string `json:"autoPayFlg" validate:"required"`
		// Авто төлөлт холбох данс. autoPayFlg = Y үед
		PayAccount string `json:"payAccount,omitempty"`
	}
)

// 7.4 Хамтран нэмж зээл судлуулах
type (
	LoanJointAddReq struct {
		// Өмнөх зээлийн хүсэлтийн дугаар
		RefID string `json:"refId" validate:"required"`
		// Үндсэн зээл хүсэгчийн регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Хамтран зээл хүсэгчийн регистрийн дугаар
		JointRegNumber string `json:"jointRegNumber" validate:"required"`
		// Хамтран зээлдэгчийн и-мэйл
		Email string `json:"email" validate:"required"`
		// Хамтран зээлдэгчийн утас
		Phone string `json:"phone" validate:"required"`
		// Хамаарал. Лавлах төрөл: RELATION
		RelType string `json:"relType" validate:"required"`
	}
	LoanJointAddResp struct {
		// FAILED, SUCCESS
		Status string `json:"status"`
		// ДАН руу орж баталгаажуулах линк
		Message string `json:"message"`
	}
)

// 7.7, 7.8 Байгууллагын зээл
type (
	LoanCorporateClosureReq struct {
		// Байгууллагын регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Зээлийн дансны дугаар
		LoanAccount string `json:"loanAccount" validate:"required"`
		// Зээлийн дүн суутгагдах харилцах данс
		OperAccount string `json:"operAccount" validate:"required"`
	}
	LoanCorporateClosureResp struct {
		// Төлөв
		Status string `json:"status"`
		// Тайлбар
		Message string `json:"message"`
	}
	LoanCorporateReq struct {
		// SL - supplier led, BL - buyer led
		Type string `json:"type" validate:"required"`
		// Нэхэмжлэх дугаар
		InvoiceID string `json:"invoiceId" validate:"required"`
		// Зээлийн мэдээлэл
		LoanDetails LoanCorporateDetails `json:"loanDetails" validate:"required"`
		// Нийлүүлэгчийн мэдээлэл
		Supplier LoanCorporateParty `json:"supplier" validate:"required"`
		// Худалдан авагчийн мэдээлэл
		Buyer LoanCorporateParty `json:"buyer" validate:"required"`
	}
	LoanCorporateDetails struct {
		// Зээлийн дүн
		Amount float64 `json:"amount" validate:"required"`
		// Шимтгэлийн дүн
		FeeAmount float64 `json:"feeAmount" validate:"required"`
		// Зээлийн хугацаа
		Period int `json:"period" validate:"required"`
		// CLA - байгууллагын зээл
		ProdType string `json:"prodType" validate:"required"`
		// Бүтээгдэхүүн
		ProdCode string `json:"prodCode" validate:"required"`
		// Салбарын дугаар
		Branch string `json:"branch" validate:"required"`
		// Зээлийн валют
		Currency string `json:"currency" validate:"required"`
		// Төлөх давтамж. M - сар бүр
		InstallmentType string `json:"installmentType" validate:"required"`
		// Олголтын шимтгэл төлөх данс
		ChrgAccountID string `json:"chrgAccountId" validate:"required"`
		// Сарын хүү
		Interest       float64 `json:"interest" validate:"required"`
		GroupCode      string  `json:"groupCode,omitempty"`
		PurposeCode    string  `json:"purposeCode,omitempty"`
		SubPurposeCode string  `json:"subPurposeCode,omitempty"`
		// Зээл төлөх өдрүүд
		Paydays []int `json:"paydays" validate:"required"`
	}
	LoanCorporateParty struct {
		// Байгууллагын регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Нэр
		Name string `json:"name" validate:"required"`
		// Харилцах данс
		AccountNumber string `json:"accountNumber" validate:"required"`
		// Зээлийн шугамын дугаар
		LineID string `json:"lineId" validate:"required"`
	}
	LoanCorporateResp struct {
		// Зээлийн дансны дугаар
		Account string `json:"account"`
		// Төлөв
		Status string `json:"status"`
		// Тайлбар
		Message string `json:"message"`
	}
)

// 7.9 Харилцагч авах зээлийн дүнг баталгаажуулах
type (
	LoanConfirmReq struct {
		// Баталгаажуулах зээлийн хэмжээ
		ApprovedAmt float64 `json:"approvedAmt" validate:"required"`
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId" validate:"required"`
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Зээл олгох дансны шинэ данс эсэх. Y, N (хэрэглээний зээл)
		NewAcctFlg string `json:"newAcctFlg,omitempty"`
		// Зээл хүлээж авах данс (хэрэглээний зээл)
		CreditAccount string `json:"creditAccount,omitempty"`
		// Зээл хүлээж авах дансны нэр (хэрэглээний зээл)
		CreditAccountName string `json:"creditAccountName,omitempty"`
		// Зээлийн төлбөр төлөх өдөр. M үед [5], T үед [5, 20]. Лавлах төрөл: PAYDAY
		Payday []int `json:"payday,omitempty"`
		// Харилцагч бүртгэх эсэх. Y, N
		CustRegFlg string `json:"custRegFlg,omitempty"`
		// Шинэ харилцагчийн мэдээлэл
		CustomerInfo *LoanConfirmCustomerInfo `json:"customerInfo,omitempty"`
	}
	LoanConfirmCustomerInfo struct {
		// Иргэншил
		Nationality string `json:"nationality" validate:"required"`
		// Төрсөн улс
		PlaceOfBirth string `json:"placeOfBirth" validate:"required"`
		// Сууц өмчлөлийн хэлбэр. Лавлах төрөл: GLM_OWNERTYPE
		ResidenceType string `json:"residenceType" validate:"required"`
		// Сууцны төрөл. Лавлах төрөл: RESIDENCE_STATUS
		ResidenceStatus string `json:"residenceStatus" validate:"required"`
		// Ажлын газрын нэр
		OrganizationName string `json:"organizationName" validate:"required"`
		// Сарын орлого
		MonthlyIncome float64 `json:"monthlyIncome" validate:"required"`
		// Орлогын эх үүсвэр. Лавлах төрөл: GLM_SOURCE_INCOME
		IncomeSource string `json:"incomeSource" validate:"required"`
		// Төгссөн сургуулийн нэр
		SchoolName string `json:"schoolName,omitempty"`
		// Ажил эрхлэлт. Лавлах төрөл: EMPLOYMENT_STATUS
		EmploymentStatus string `json:"employmentStatus" validate:"required"`
		// Ам бүлийн тоо
		NumFamily int `json:"numFamily" validate:"required"`
		// Орлоготой гишүүдийн тоо
		NumFamilyIncome int `json:"numFamilyIncome" validate:"required"`
		// Ургийн овог
		FamilyName string `json:"familyName" validate:"required"`
		// АНУ ногоон карттай эсэх. Y, N
		GreenCardFlg string `json:"greenCardFlg" validate:"required"`
		// Яаралтай үед холбоо барих хүний утас
		OthersPhone string `json:"othersPhone" validate:"required"`
		// Яаралтай үед холбоо барих хүний нэр
		OthersLName string `json:"othersLName" validate:"required"`
		// Яаралтай үед холбоо барих хүний овог
		OthersFName string `json:"othersFName" validate:"required"`
		// Таны хэн болох
		OthersRelation string `json:"othersRelation" validate:"required"`
	}
	LoanConfirmResp struct {
		// FAILED, SUCCESS
		Status string `json:"status"`
		// Мессеж
		Message string `json:"message"`
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Зээл бол зээлийн дансны дугаар, кредит карт бол картын токен
		Account string `json:"account"`
		// Өмнөх зээлийн хүүг шинэ хүүгээр update хийсэн эсэх. Y, N
		InterestUpdateFlg string `json:"interestUpdateFlg"`
	}
)

// 7.10 - 7.15 Зээлийн хүсэлтийн лавлагаа, баталгаажуулалт
type (
	LoanListReq struct {
		// Данс эзэмшигчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	LoanListData struct {
		// LOAN, CARD
		ProdType string `json:"prodType"`
		// Зээлийн хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Зээл хүссэн дүн
		Amount float64 `json:"amount"`
		// Урьдчилгаа төлбөр
		PrepaidAmt float64 `json:"prepaidAmt"`
		// Scoring дууссан огноо
		ScoringDate string `json:"scoringDate"`
		// Олгож болох зээлийн дээд хэмжээ
		ScoringAmt float64 `json:"scoringAmt"`
		// Сард төлөх дүн
		MonthlyAmt float64 `json:"monthlyAmt"`
		// Зээлийн хугацаа
		Period int `json:"period"`
		// Шимтгэлийн дүн
		ChrgAmt float64 `json:"chrgAmt"`
		// Зээлийн хүү
		Interest float64 `json:"interest"`
		// Салбарын дугаар
		BranchID string `json:"branchId"`
		// Валют
		CrnCode string `json:"crnCode"`
		// Сард төлөлт хийх давтамж. M, T
		InstType string `json:"instType"`
		// Бүтээгдэхүүний код
		ProdCode string `json:"prodCode"`
		// Карт дээр бичигдэх нэр
		EmbossName string `json:"embossName"`
		// Зээлийн төлвийн мэдээлэл
		LoanSteps []LoanStep `json:"loanSteps"`
		// Зээлдэгчийн мэдээлэл
		Customer LoanCustomer `json:"customer"`
	}
	LoanStep struct {
		// Зээлийн алхам. LOAN_REQ г.м
		StepCd string `json:"stepCd"`
		// Алхамын тайлбар
		StepDesc string `json:"stepDesc"`
		// Алхамын төлөв. SENT г.м
		StatusCd string `json:"statusCd"`
		// Алхамын төлвийн тайлбар
		StatusDesc string `json:"statusDesc"`
	}
	// 7.11, 7.12, 7.13 хүсэлт
	LoanAppReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId" validate:"required"`
	}
	LoanStatusResp struct {
		// Зээл хүссэн дүн
		Amount float64 `json:"amount"`
		// Урьдчилгаа төлбөр
		PrepaidAmt float64 `json:"prepaidAmt"`
		// Scoring дууссан огноо
		ScoringDate string `json:"scoringDate"`
		// Олгож болох зээлийн дээд хэмжээ
		ScoringAmt float64 `json:"scoringAmt"`
		// Сард төлөх дүн
		MonthlyAmt float64 `json:"monthlyAmt"`
		// Зээлийн хугацаа
		Period int `json:"period"`
		// Шимтгэлийн дүн
		ChrgAmt float64 `json:"chrgAmt"`
		// Зээлийн хүү
		Interest float64 `json:"interest"`
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Голомт банкны харилцагч эсэх. Y, N
		IsCustomer string `json:"isCustomer"`
		// Зээлийн мастер гэрээтэй эсэх. Y, N
		MasterContrFlg string `json:"masterContrFlg"`
		// Баталгаажсан зээлийн данс / картын дугаар
		NewAccount string `json:"newAccount"`
		// Зээлийн төлвийн мэдээлэл
		LoanSteps []LoanStep `json:"loanSteps"`
	}
	LoanContractResp struct {
		// application/pdf төрөлтэй byte array (base64)
		File []byte `json:"file"`
	}
	LoanCancelReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId" validate:"required"`
		// Цуцлах болсон шалтгаан
		Reason string `json:"reason,omitempty"`
	}
	LoanCheckRecordReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	LoanCheckRecordResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Идэвхитэй хүсэлт байгаа эсэх. Y, N
		IsActiveExists string `json:"isActiveExists"`
		// Зээлийн хүсэлтүүд
		AppRecords []LoanAppRecord `json:"appRecords"`
	}
	LoanAppRecord struct {
		// PENDING, THIRD_PARTY_APPROVAL, DISBURSEMENT
		Step string `json:"step"`
		// Хүсэлт үүсгэсэн гуравдагч систем
		CreatedBy string `json:"createdBy"`
		// Хүсэлт үүсгэсэн огноо
		CreatedDate string `json:"createdDate"`
	}
)

// 7.16, 7.17 Голомт банкнаас гуравдагч систем рүү илгээх хүсэлтүүд
type (
	// 7.16 Зээлийн онлайн тооцооллын callback хүсэлт
	LoanCallbackReq struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo"`
		// Зээл хүссэн дүн
		Amount float64 `json:"amount"`
		// Урьдчилгаа төлбөр
		PrepaidAmt float64 `json:"prepaidAmt"`
		// Scoring дууссан огноо
		ScoringDate string `json:"scoringDate"`
		// Олгож болох зээлийн дээд хэмжээ
		ScoringAmt float64 `json:"scoringAmt"`
		// Сард төлөх дүн
		MonthlyAmt float64 `json:"monthlyAmt"`
		// Зээлийн хугацаа
		Period int `json:"period"`
		// Шимтгэлийн дүн
		ChrgAmt float64 `json:"chrgAmt"`
		// Зээлийн хүү
		Interest float64 `json:"interest"`
		// Голомт банкны харилцагч эсэх. Y, N
		IsCustomer string `json:"isCustomer"`
		// Зээлийн мастер гэрээтэй эсэх. Y, N
		MasterContrFlg string `json:"masterContrFlg"`
		// Зээлийн төлвийн мэдээлэл
		LoanSteps []LoanStep `json:"loanSteps"`
	}
	// 7.16 callback-д гуравдагч системийн буцаах хариу
	LoanCallbackResp struct {
		// FAILED, SUCCESS
		Status string `json:"status"`
		// Мессеж
		Message string `json:"message"`
	}
	// 7.17 Ногоон бүтээгдэхүүний жагсаалт татах хүсэлт
	LoanGreenProductReq struct {
		// Захиалгын дугаар
		BillID string `json:"billId"`
	}
	// 7.17 гуравдагч системийн буцаах хариу
	LoanGreenProductResp struct {
		// Харилцагчийн дугаар
		UserID string `json:"userId"`
		// Бүтээгдэхүүний жагсаалт
		Products []LoanGreenProduct `json:"products"`
	}
	LoanGreenProduct struct {
		// Барааны код
		Code string `json:"code"`
		// Барааны нэр
		Name string `json:"name"`
		// Бренд
		Brand string `json:"brand"`
		// Барааны мөнгөн дүн
		Amount float64 `json:"amount"`
		// Барааны ангилал
		Category string `json:"category"`
	}
)
