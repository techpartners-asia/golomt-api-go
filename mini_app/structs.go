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

	SendNotificationInput struct {
		IndividualId string `json:"individualId"`
		MessageTitle string `json:"messageTitle"`
		MessageText  string `json:"messageText"`
		VIEW_ID      string `json:"VIEW_ID"`
		MINIAPP_URL  string `json:"MINIAPP_URL"`
	}

	SendNotificationResponse struct {
		Message string `json:"message"`
	}
)
