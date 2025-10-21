package model

type (
	StateListReq struct {
		// Хот, аймагийн код: ALL
		StateCode string `json:"stateCode" validate:"required"`
	}
	StateListResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Хот, аймагийн код
		Code string `json:"code"`
		// Хот, аймагийн нэр
		Name string `json:"name"`
	}
	DistrictListReq struct {
		// STATEINQ хариу мэдэгдэл дээр ирсэн код байна
		StateCode string `json:"stateCode" validate:"required"`
	}
	DistrictListResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Сум, дүүргийн код
		Code string `json:"code"`
		// Сум, дүүргийн нэр
		Name string `json:"name"`
	}
	CategoryReq struct {
		// Лавлах төрөл: SECTOR_CODE
		Type string `json:"type" validate:"required"`
	}
	CategoryResp struct {
		// Төрөл
		Type string `json:"type"`
		// Код
		Code string `json:"code"`
		// Label
		Label string `json:"label"`
		// Тайлбар
		Description string `json:"description"`
	}

	BranchListReq struct {
		// Салбарын дугаар; ALL
		SolId string `json:"solId" validate:"required"`
	}
	BranchListResp struct {
		// Салбарын дугаар
		BranchID string `json:"branchId"`
		// Салбарын нэр
		BranchName string `json:"branchName"`
	}

	ProductListReq struct {
		// Category төрөл; ALL
		Type string `json:"type" validate:"required"`
	}
	ProductData struct {
		// Бүтээгдэхүүний төрөл
		ProductType string `json:"prodType"`
		// Бүтээгдэхүүний код
		Code string `json:"code"`
		// Тайлбар нэр
		Description string `json:"description"`
		// Нээх боломжтой харилцагчийн төрөл
		// R-энгийн харилцагч
		// C-Байгууллагын харилцагч
		// B-both
		CustormerType string `json:"custType"`
		//Бүтээгдэхүүн үүсгэх доод үлдэгдэл
		Minbalances []MinBalanceData `json:"minbalances"`
		// Бүтээгдэхүүн үүсгэх нөхцөл
		Interests []InterestData `json:"interests"`
	}

	MinBalanceData struct {
		// Валют
		Currency string `json:"currency"`
		// Хамгийн бага үлдэгдэл
		MinBalance float64 `json:"minBalance"`
	}

	InterestData struct {
		// Хугацаа
		Month int `json:"month"`
		// Хүүний хэмжээ
		Interest float64 `json:"interest"`
	}

	RateReq struct {
		// Валют
		Currency string `json:"currency"`
	}
	RateResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Өдрийн огноо
		Date string `json:"date"`
		// Дарааллын дугаар
		Sequence int `json:"sequence"`
		// Валютүүд
		Currencies []RateData `json:"currencies"`
	}
	RateData struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Валют код
		CurrencyCode string `json:"currencyCode"`
		// Валют нэр
		CurrencyName string `json:"currencyName"`
		// Бэлэн ханш авах
		CashValueSell string `json:"cashValueSell"`
		// Бэлэн ханш зарах
		CashValueBuy string `json:"cashValueBuy"`
		// Бэлэн бус ханш авах
		NonCashValueSell string `json:"nonCashValueSell"`
		// Бэлэн бус ханш зарах
		NonCashValueBuy string `json:"nonCashValueBuy"`
	}
)
