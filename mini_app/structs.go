package mini_app

type (
	UserInfo struct {
		IndividualId   int64  `json:"individualId"`   // Хэрэглэгчийн дугаар
		RegisterNumber string `json:"registerNumber"` // Регистрийн дугаар
		LastName       string `json:"lastName"`       // Овог
		FirstName      string `json:"firstName"`      // Нэр
		Account        string `json:"account"`        // Дансны дугаар
		MobileNumber   string `json:"mobileNumber"`   // Утасны дугаар
		Email          string `json:"email"`          // Имэйл
		ImgUrl         string `json:"imgUrl"`         // SocialPay дахь хэрэглэгчийн профайл зураг
	}
)
