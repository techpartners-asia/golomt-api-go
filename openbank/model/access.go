package model

type (
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
	ServiceListReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Сервисүүдийн жагсаалтыг агуулсан массив байна
		Services []string `json:"services" validate:"required"`
		// Тухайн нэг сервисийн X-Golomt-Service дугаар байна.
		Code string `json:"code" validate:"required"`
	}

	ServiceListResp struct {
		// Тус байгууллагыг таних зорилготой дахин давдагдашгүй дугаар
		ClientID string `json:"clientId"`
		// “Access Grant Response” буцаалтаар ирсэн code параметрийн утга (grant code).
		ResponseType string `json:"responseType"`
		// Тус байгууллагын Буцах URL байна.
		// Энэхүү URL –ийн банканд өгч бүртгүүлэх шаардлагатай.
		RedirectUri string `json:"redirectUri"`
		// Тус хүсэлтийг encrypt хийсэн утга
		Scope string `json:"scope"`
		// Дахин давтагдашгүй дугаар бөгөөд тус хүсэлтийг илтгэнэ.
		State string `json:"state"`
	}

	GetPhoneResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// OTP код илгээх харилцагчийн дугаар
		MaskedPhone string `json:"maskedPhone"`
	}

	OTPSendResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Хүсэлтийн төлөв.
		//
		// SUCCESS эсвэл FAILED хариу буцаана
		Status string `json:"status"`
		// OTP илгээсэн харилцагчийн дугаар
		MaskedPhone string `json:"maskedPhone"`
		// Хүсэлтийн тайлбар
		Message string `json:"message"`
	}

	OTPVerifyReq struct {
		// Тус байгууллагыг таних зорилготой дахин давдагдашгүй дугаар
		ClientID string `json:"clientId" validate:"required"`
		// Харилцагчийн дугаар дээр очсон код
		OTP string `json:"otp" validate:"required"`
		// Тус байгууллагын Буцах URL байна.
		//
		// Энэхүү URL –ийн банканд өгч бүртгүүлэх шаардлагатай.
		RedirectUri string `json:"redirectUri" validate:"required"`
		// Тус хүсэлтийг encrypt хийсэн утга
		Scope string `json:"scope" validate:"required"`
		// Дахин давтагдашгүй дугаар бөгөөд тус хүсэлтийг илтгэнэ.
		State string `json:"state" validate:"required"`
	}

	OTPXypVerifyReq struct {
		// Тус байгууллагыг таних зорилготой дахин давдагдашгүй дугаар
		ClientID string `json:"clientId" validate:"required"`
		// Харилцагчийн дугаар дээр очсон код
		OTP string `json:"otp" validate:"required"`
		// Тус хүсэлтийг encrypt хийсэн утга
		Scope string `json:"scope" validate:"required"`
		// Дахин давтагдашгүй дугаар бөгөөд тус хүсэлтийг илтгэнэ.
		State string `json:"state" validate:"required"`
	}

	OTPVerifyResp struct {
		// Зөв эсхээс хамаарч SUCCESS эсвэл FAILED хариу буцаана
		Status string `json:"status"`
	}

	DigitalSignatureResp struct {
		// Зөв эсхээс хамаарч SUCCESS эсвэл FAILED хариу буцаана
		Status string `json:"status"`
		// Хүсэлтийн тайлбар
		Desc string `json:"desc"`
	}
)
