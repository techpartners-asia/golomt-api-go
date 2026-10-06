package model

// 6.1. Харилцагч шалгах
type (
	CustomerCheckReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	CustomerCheckResp struct {
		// Харилцагч эсэх. YES – тийм, NO - үгүй
		IsCustomer string `json:"isCustomer"`
	}
)

// 6.2. Харилцагчийн утасны дугаар мөн эсхийг шалгах
type (
	CustomerPhoneCheckReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Харилцагчийн утасны дугаар
		PhoneNum string `json:"phoneNum" validate:"required"`
	}
	CustomerPhoneCheckResp struct {
		// Харилцагчийн дугаар мөн эсэх. Y – тийм, N - үгүй
		IsCustomerPhoneCheck string `json:"isCustomerPhoneCheck"`
	}
)

// 6.3. Харилцагчийн мэдээлэл харах
type (
	CustomerInquireReq struct {
		// Банканд бүртгэлтэй харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	CustomerInquireResp struct {
		// Үндсэн мэдээлэл
		Customer CustomerInquireInfo `json:"customer"`
		// Холбоо барих мэдээлэл
		PhoneEmails []CustomerInquirePhoneEmail `json:"phoneEmails"`
	}
	CustomerInquireInfo struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo"`
		// Өөрийн нэр
		FirstName string `json:"firstName"`
		// Эцэг- эхийн нэр
		LastName string `json:"lastName"`
		// Бүтэн нэр
		Name string `json:"name"`
		// Овог
		FamilyName string `json:"familyName"`
		// Хүйс
		Salutation string `json:"salutation"`
		// Төрсөн огноо (yyyy-MM-dd)
		BirthDate string `json:"birthDate"`
		// Үндсэн салбарын дугаар
		BranchID string `json:"branchId"`
		// Интернэт банк ашигладаг эсэх
		IsInternetBankingEnabled bool `json:"isInternetBankingEnabled"`
	}
	CustomerInquirePhoneEmail struct {
		// PHONE - Утасны дугаар, EMAIL - и-мэйл хаяг
		Type string `json:"type"`
		// Утас, И-мэйл хаягийн төрөл. CELLPH, HOMEEML г.м
		SubType string `json:"subType"`
		// Улсын код
		CountryCode string `json:"countryCode"`
		// Type == EMAIL үед и-мэйл хаяг
		EmailID string `json:"emailId"`
		// Type == PHONE үед утасны дугаар
		PhoneNo string `json:"phoneNo"`
	}
)

// 6.4. Иргэн – Харилцагч бүртгэх
type (
	// Харилцагчийн зураг хуулах хүсэлт
	CustomerImageUploadReq struct {
		// Харилцагчийн овог нэр
		Name string `json:"name" validate:"required"`
		// Харилцагчийн регистрийн дугаар
		RegNo string `json:"regNo" validate:"required"`
		// Утасны дугаар
		PhoneNo string `json:"phoneNo,omitempty"`
		// И-мэйл хаяг
		Email string `json:"email,omitempty"`
		// Зурагны төрөл. IDFRONT, IDBACK, SELFIE
		ImgCat string `json:"imgCat" validate:"required"`
		// Хуулах зурагны нэр
		ImgName string `json:"imgName" validate:"required"`
		// Base64 руу хөрвүүлсэн зурагны файл
		ImgBase64 string `json:"imgBase64" validate:"required"`
	}
	CustomerImageUploadResp struct {
		// Лавлах дугаар
		Code int64 `json:"code"`
		// SUCCESS – Амжилттай, бусад үед амжилтгүй
		Status string `json:"status"`
	}

	// Иргэн харилцагч бүртгэх хүсэлт
	CustomerRetailAddReq struct {
		// Харилцагчийн эцэг – эхийн нэр
		LastName string `json:"lastName" validate:"required"`
		// Харилцагчийн өөрийн нэр
		FirstName string `json:"firstName" validate:"required"`
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Харилцагчийн хүйс. M - эрэгтэй, F - эмэгтэй
		CustGender string `json:"custGender" validate:"required"`
		// Харилцагчийн ажлын газрын нэр
		OrgName string `json:"orgName" validate:"required"`
		// Харилцагчийн ажлын газрын хаяг
		OrgAddress string `json:"orgAddress" validate:"required"`
		// Мэргэжил
		Occupation string `json:"occupation" validate:"required"`
		// Утас, и-мэйл
		PhoneEmails []CustomerPhoneEmail `json:"phoneEmails"`
		// Хаягийн мэдээлэл
		Address []CustomerRetailAddress `json:"address"`
	}
	CustomerPhoneEmail struct {
		// EMAIL эсвэл PHONE
		Type string `json:"type" validate:"required"`
		// Type == PHONE үед заавал
		Phone string `json:"phone,omitempty"`
		// Type == EMAIL үед заавал
		Email string `json:"email,omitempty"`
		// Улсын код
		CountryLocalCode string `json:"countryLocalCode" validate:"required"`
	}
	CustomerRetailAddress struct {
		// Улс орны мэдээлэл
		Country string `json:"country" validate:"required"`
		// Хот эсвэл аймаг
		State string `json:"state" validate:"required"`
		// Сум дүүргийн мэдээлэл
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
		// Хаягийн мэдээлэл – 1. Формат: хороо,гудамж,байр,тоот
		AddressLine1 string `json:"addressLine1" validate:"required"`
		// Хаягийн мэдээлэл – 2
		AddressLine2 string `json:"addressLine2,omitempty"`
		// Хаягийн мэдээлэл – 3
		AddressLine3 string `json:"addressLine3,omitempty"`
	}
	// Харилцагч бүртгэсэн хариу
	CustomerAddResp struct {
		// Лавлах дугаар
		ReferenceID string `json:"referenceId"`
		// Харилцах дансны дугаар
		AccountID string `json:"accountId"`
	}
)

// 6.5. Байгууллага - Харилцагч бүртгэх
type (
	CustomerCorpAddReq struct {
		// Байгууллагын нэр. Зөвхөн үсэг
		CorporateName string `json:"corporateName" validate:"required"`
		// Байгууллагын регистрийн дугаар
		RegisterNumber string `json:"registerNumber,omitempty"`
		// Байгууллага дээр данс үүсгэх эсэх. Y, N
		CreateBankAccount string `json:"createBankAccount" validate:"required"`
		// Байгууллагын хаяг
		CorporateAddress []CustomerCorpAddress `json:"corporateAddress"`
		// Утас, и-мэйл
		PhoneEmails []CustomerPhoneEmail `json:"phoneEmails"`
		// Эзэмшигчийн мэдээлэл
		Retail CustomerCorpRetail `json:"retail" validate:"required"`
	}
	CustomerCorpAddress struct {
		// Хаягийн мэдээлэл
		AddressLine1 string `json:"addressLine1" validate:"required"`
		// Улс орны мэдээлэл
		Country string `json:"country" validate:"required"`
		// Хот эсвэл аймаг
		State string `json:"state" validate:"required"`
		// Сум дүүргийн мэдээлэл
		City string `json:"city" validate:"required"`
	}
	CustomerCorpRetail struct {
		// Иргэний регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Ургийн овог
		FamilyName string `json:"familyName" validate:"required"`
		// Эцэг эхийн нэр
		LastName string `json:"lastName" validate:"required"`
		// Эзэмшигчийн нэр
		FirstName string `json:"firstName" validate:"required"`
		// Хүйс. MR., MS.
		Salutation string `json:"salutation" validate:"required"`
		// Төрсөн огноо (yyyy-MM-dd)
		BirthDate string `json:"birthDate" validate:"required"`
		// Утас, и-мэйл
		PhoneEmails []CustomerPhoneEmail `json:"phoneEmails"`
		// Хаягийн мэдээлэл
		Address []CustomerCorpAddress `json:"address"`
	}
)

// 6.6. Харилцагчийн мэдээлэл бүртгэх
type (
	// Насанд хүрсэн болон 0-18 насны иргэний аль алинд ашиглана
	CustomerSaveInfoReq struct {
		// Харилцагчийн мэдээлэл
		Customer CustomerSaveInfo `json:"customer" validate:"required"`
		// Харилцагчийн нэмэлт мэдээлэл (насанд хүрсэн иргэн)
		Demographic *CustomerDemographic `json:"demographic,omitempty"`
		// Харилцагчийн оршин суугаа хаягийн мэдээлэл
		Address CustomerInfoAddress `json:"address" validate:"required"`
		// Асран хамгаалагчийн мэдээлэл (0-18 насны иргэн)
		Relatives []CustomerRelative `json:"relatives,omitempty"`
		// Харилцагчийн холбогдох хүмүүсийн мэдээлэл
		ContactInformationList []CustomerContactInformation `json:"contactInformationList,omitempty"`
	}
	CustomerSaveInfo struct {
		// Харилцагчийн өөрийн нэр
		FirstName string `json:"firstName" validate:"required"`
		// Харилцагчийн эцэг-эхийн нэр
		LastName string `json:"lastName" validate:"required"`
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Иргэний бүртгэлийн дугаар
		CivilID string `json:"civilId" validate:"required"`
		// Харилцагчийн ургийн овог
		FamilyName string `json:"familyName" validate:"required"`
		// Төрсөн өдөр (yyyy-MM-dd)
		Birthdate string `json:"birthdate" validate:"required"`
		// Харилцагчийн холбогдох mail хаяг
		Email string `json:"email" validate:"required"`
		// Харилцагчийн холбогдох утасны дугаар
		Phone string `json:"phone" validate:"required"`
		// Ажлын газрын нэр (насанд хүрсэн иргэн)
		EmployerName string `json:"employerName,omitempty"`
		// Иргэншил
		Nationality string `json:"nationality" validate:"required"`
		// Мэдээллийн эзний зөвшөөрөл. Y, N
		MarketingConsent string `json:"marketingConsent,omitempty"`
		// Америкийн green card-тай эсэх. Y, N
		GreenCardUSFlg string `json:"greenCardUSFlg" validate:"required"`
	}
	CustomerDemographic struct {
		// Ажил эрхлэлт. Лавлах: EMPLOYMENT_STATUS
		EmploymentStatus string `json:"employmentStatus,omitempty"`
		// Үйл ажиллагааны чиглэл. Лавлах: INDUSTRY_TYPE
		IndustryType string `json:"industryType,omitempty"`
		// Мэргэжил. Лавлах: APPOINTMENT
		Occupation string `json:"occupation,omitempty"`
		// Албан тушаалын код. Лавлах: OCCUPATION
		Appointment string `json:"appointment,omitempty"`
	}
	CustomerInfoAddress struct {
		// Улс орны мэдээлэл
		Country string `json:"country" validate:"required"`
		// Хот эсвэл аймаг
		State string `json:"state" validate:"required"`
		// Сум дүүргийн мэдээлэл
		City string `json:"city" validate:"required"`
		// Хаягийн төрөл. HOME, WORK
		Type string `json:"type" validate:"required"`
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
		// Хаягийн мэдээлэл – 1. Формат: хороо,гудамж,байр,тоот
		AddressLine1 string `json:"addressLine1" validate:"required"`
	}
	CustomerRelative struct {
		// Асран хамгаалагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Харилцагчийн хэн болох. Лавлах: GUARDIAN_TYPE
		RelativeType string `json:"relativeType" validate:"required"`
		// Асран хамгаалагчийн утасны дугаар
		RelativePhone string `json:"relativePhone" validate:"required"`
	}
	CustomerContactInformation struct {
		// Холбогдох хүний нэр
		FirstName string `json:"firstName" validate:"required"`
		// Холбогдох хүний овог
		LastName string `json:"lastName" validate:"required"`
		// Харилцагчийн хэн болох. Лавлах: RELATION
		Type string `json:"type" validate:"required"`
		// Холбогдох хүний утасны дугаар
		PhoneNo string `json:"phoneNo" validate:"required"`
	}
	CustomerSaveInfoResp struct {
		// Бүртгэлийн лавлах дугаар
		InfoID string `json:"infoId"`
	}
)

// 6.7. Бүртгэгдсэн харилцагчийн мэдээлэл татах
type (
	CustomerGetInfoReq struct {
		// Бүртгэлийн лавлах дугаар
		InfoID string `json:"infoId" validate:"required"`
	}
	CustomerGetInfoResp struct {
		// Харилцагчийн мэдээлэл
		CustomerInfo CustomerGetInfo `json:"customerInfo"`
		// Харилцагчийн нэмэлт мэдээлэл
		Demographic CustomerGetInfoDemographic `json:"demographic"`
		// Оршин суугаа хаягийн мэдээлэл
		Address CustomerInfoAddress `json:"address"`
		// Асран хамгаалагчийн мэдээлэл
		RelativesList []CustomerRelative `json:"relativesList"`
		// Холбогдох хүмүүсийн мэдээлэл
		ContactInfoList []CustomerContactInfo `json:"contactInfoList"`
	}
	CustomerGetInfo struct {
		// Харилцагчийн өөрийн нэр
		FirstName string `json:"firstName"`
		// Харилцагчийн эцэг-эхийн нэр
		LastName string `json:"lastName"`
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo"`
		// Харилцагчийн ургийн овог
		FamilyName string `json:"familyName"`
		// Иргэний бүртгэлийн дугаар
		CivilID string `json:"civilId"`
		// Төрсөн өдөр
		Birthdate string `json:"birthdate"`
		// Холбогдох mail хаяг
		Email string `json:"email"`
		// Холбогдох утасны дугаар
		Phone string `json:"phone"`
		// Иргэншил
		Nationality string `json:"nationality"`
		// Ажлын газрын нэр
		EmployerName string `json:"employerName"`
		// Мэдээллийн эзний зөвшөөрөл. Y, N
		MarketingConsent string `json:"marketingConsent"`
		// Хүйс. MS., MR.
		Gender string `json:"gender"`
		// Өрхийн орлоготой гишүүдийн тоо
		FamilyMembersWithIncome string `json:"familyMembersWithIncome"`
		// Ам бүлийн тоо
		NumOfFamily string `json:"numOfFamily"`
		// Сарын орлого
		MonthlyIncome string `json:"monthlyIncome"`
		// Орлогын байдал
		IncomeStatus string `json:"incomeStatus"`
		// Орлогын давтамж
		IncomeFrequency string `json:"incomeFrequency"`
		// Орлогын хэлбэр
		IncomeType string `json:"incomeType"`
		// Хөрөнгө орлогын эх үүсвэр
		IncomeSource string `json:"incomeSource"`
		// Америкийн ногоон карттай эсэх. Y, N
		GreenCardUSFlg string `json:"greenCardUSFlg"`
		// Иргэний үнэмлэхний нүүр талын зураг upload хийсэн эсэх
		HasIDFrontPic string `json:"hasIdFrontPic"`
		// Иргэний үнэмлэхний ар талын зураг upload хийсэн эсэх
		HasIDBackPic string `json:"hasIDBackPic"`
		// Selfie зураг upload хийсэн эсэх
		HasSelfiePic string `json:"hasSelfiePic"`
		// Төрсний гэрчилгээний зураг upload хийсэн эсэх
		HasCertifPic string `json:"hasCertifPic"`
	}
	CustomerGetInfoDemographic struct {
		// Албан тушаалын код
		Appointment string `json:"appointment"`
		// Мэргэжил
		Occupation string `json:"occupation"`
		// Үйл ажиллагааны чиглэл
		IndustryType string `json:"industryType"`
		// Ажил эрхлэлт
		EmploymentStatus string `json:"employmentStatus"`
		// Сургуульд орсон огноо (yyyy-MM-dd)
		EnrollmentDate string `json:"enrollmentDate"`
		// Төгссөн сургууль
		SchoolName string `json:"schoolName"`
		// Мэргэжлийн зэрэг
		Degree string `json:"degree"`
		// Гэрлэлтийн төлөв
		MaritalStatus string `json:"maritalStatus"`
		// Ажилдаа орсон огноо (yyyy-MM-dd)
		StartDate string `json:"startDate"`
		// Хөдөлмөр эрхэлсэн нийт жил
		YearsWork string `json:"yearsWork"`
	}
	CustomerContactInfo struct {
		// Холбогдох хүний овог нэр
		Name string `json:"name"`
		// Холбогдох хүний утасны дугаар
		PhoneNo string `json:"phoneNo"`
		// Харилцагчийн хэн болох
		Type string `json:"type"`
	}
)

// 6.8. - 6.9. Харилцагчийн иргэний үнэмлэхний зураг upload хийх
type (
	CustomerNewImageUploadReq struct {
		// Бүртгэлийн лавлах дугаар
		InfoID string `json:"infoId" validate:"required"`
		// Зургийн төрөл. IDFRONT, IDBACK, SELFIE, CERTIF
		ImgCat string `json:"imgCat" validate:"required"`
		// Зургийн нэр
		ImgName string `json:"imgName" validate:"required"`
		// Зургийн файл /base64/
		ImgBase64 string `json:"imgBase64" validate:"required"`
	}
	CustomerNewImageFormUploadInput struct {
		// Бүртгэлийн лавлах дугаар
		InfoID string `json:"infoId" validate:"required"`
		// Зургийн төрөл. IDFRONT, IDBACK, SELFIE, CERTIF
		ImgCat string `json:"imgCat" validate:"required"`
		// Зургийн нэр
		ImgName string `json:"imgName" validate:"required"`
		// Зургийн файлын нэр (multipart filename)
		FileName string `json:"fileName" validate:"required"`
		// Зургийн файл /MultipartFile/
		ImgFile []byte `json:"imgFile" validate:"required"`
	}
	CustomerNewImageUploadResp struct {
		// Статус
		Status string `json:"status"`
		// Upload хийгдсэн огноо
		Date string `json:"date"`
		// Зургийн id
		Reference string `json:"reference"`
	}
)

// 6.10. Харилцагчийн иргэний үнэмлэхний зураг татах
type (
	CustomerImageDownloadReq struct {
		// Хэрэглэгчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Зургийн төрөл. IDFRONT, IDBACK, SELFIE, CERTIF
		ImageType string `json:"imageType" validate:"required"`
	}
	CustomerImageDownloadResp struct {
		// base64 зурган файл
		Img string `json:"img"`
	}
)

// 6.11. Байгууллага - Харилцагчийн мэдээлэл татах
type (
	CustomerCorpDetailReq struct {
		// Байгууллагын регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	CustomerCorpDetailResp struct {
		// Sub sector код. Лавлах: SUB_SECTOR_CODE
		SubSector string `json:"subSector"`
		// Sector код. Лавлах: SECTOR_CODE
		Sector string `json:"sector"`
		// Industry код. Лавлах: INDUSTRY_TYPE
		IndustryCode string `json:"industryCode"`
		// Хаягийн мэдээлэл
		Address []CustomerCorpDetailAddress `json:"address"`
		// Холбоо барих мэдээлэл
		Contact []CustomerCorpDetailContact `json:"contact"`
	}
	CustomerCorpDetailAddress struct {
		// Улс орны мэдээлэл
		Country string `json:"country"`
		// Хаягийн төрөл. HOME, WORK
		Type string `json:"type"`
		// Хот эсвэл аймаг
		State string `json:"state"`
		// Сум дүүргийн мэдээлэл
		City string `json:"city"`
		// Хорооны дугаар
		SubDistrict string `json:"subDistrict"`
		// Гудамжын нэр
		StreetName string `json:"streetName"`
		// Хотхоны нэр
		Town string `json:"town"`
		// Байрын дугаар
		Apartment string `json:"apartment"`
		// Орцын дугаар
		Entry string `json:"entry"`
		// Тоотын дугаар
		DoorNo string `json:"doorNo"`
		// Хаягийн мэдээлэл – 1
		AddressLine1 string `json:"addressLine1"`
		// Хаягийн мэдээлэл – 2
		AddressLine2 string `json:"addressLine2"`
		// Хаягийн мэдээлэл – 3
		AddressLine3 string `json:"addressLine3"`
	}
	CustomerCorpDetailContact struct {
		// EMAIL – и-мэйл, PHONE – утас
		Type string `json:"type"`
		// CELLPH, HOMEEML, WORKEML
		SubType string `json:"subType"`
		// Утас
		Phone string `json:"phone"`
		// И-мэйл
		Email string `json:"email"`
		// Улсын код
		CountryCode string `json:"countryCode"`
	}
)

// 6.12. Байгууллага - Харилцагчийн limit-н мэдээлэл татах
type (
	CustomerCorpLimitDetailReq struct {
		// Байгууллагын регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
	}
	CustomerCorpLimitDetailResp struct {
		// Limit-н мэдээллийн жагсаалт
		CorporateCustomerLimitDetails []CustomerCorpLimitDetail `json:"corporateCustomerLimitDetails"`
	}
	CustomerCorpLimitDetail struct {
		// Дугаар
		LimitB2KID string `json:"limitB2KID"`
		// Ерөнхий тайлбар
		LimitDesc string `json:"limitDesc"`
		// Шугмын нийт дүн
		OrigSanctLim float64 `json:"origSanctLim"`
		// Нийт дүн
		SanctLim float64 `json:"sanctLim"`
		// Drawing power
		DrwngPower float64 `json:"drwngPower"`
		// Liability
		Liab float64 `json:"liab"`
		// Валют
		Crncy string `json:"crncy"`
		// Зарцуулалт
		ContingentLiab float64 `json:"contingentLiab"`
		// Үүсгэсэн огноо
		LimSanctDate string `json:"limSanctDate"`
		// Дуусах огноо
		LimExpDate string `json:"limExpDate"`
		// Шугамын зээл авах үед ашиглах дугаар
		LimPrefix string `json:"limPrefix"`
		// Төрөл
		LimType string `json:"limType"`
		// Үлдэгдэл дүн
		AvailLim float64 `json:"availLim"`
	}
)
