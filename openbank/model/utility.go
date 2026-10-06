package model

import "encoding/json"

type (
	StateListReq struct {
		// Хот, аймагийн код: ALL
		StateCode string `json:"stateCode" validate:"required"`
	}
	StateListResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Хот, аймагийн код
		StateCode string `json:"stateCode"`
		// Хот, аймагийн нэр
		StateName string `json:"stateName"`
	}
	DistrictListReq struct {
		// STATEINQ хариу мэдэгдэл дээр ирсэн код байна
		StateCode string `json:"stateCode" validate:"required"`
	}
	DistrictListResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Сум, дүүргийн код
		CityCode string `json:"cityCode"`
		// Сум, дүүргийн нэр
		CityName string `json:"cityName"`
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
		Minbalances []MinBalanceData `json:"minBalances"`
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
		Currency string `json:"currency" validate:"required"`
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
		CashValueSell json.Number `json:"cashValueSell"`
		// Бэлэн ханш зарах
		CashValueBuy json.Number `json:"cashValueBuy"`
		// Бэлэн бус ханш авах
		NonCashValueSell json.Number `json:"nonCashValueSell"`
		// Бэлэн бус ханш зарах
		NonCashValueBuy json.Number `json:"nonCashValueBuy"`
	}

	// 10.7. RateCode лавлах
	RateCodeReq struct {
		// Жишээ: EUR
		RefCurrency string `json:"refCurrency" validate:"required"`
		// Жишээ: USD
		AccCurrency string `json:"accCurrency" validate:"required"`
	}
	RateCodeResp struct {
		RequestID string `json:"requestId"`
		// CNN6S / CNN6B (BUY/SELL)
		RateCode string `json:"rateCode"`
	}

	// 10.8. Exchange Rate лавлах
	ExchangeRateReq struct {
		// Хөрвүүлэх валют
		FromCurrency string `json:"fromCurrency" validate:"required"`
		// Хөрвөсөн валют
		ToCurrency string `json:"toCurrency"`
		// CNN6S / CNN6B (BUY/SELL)
		RateCode string `json:"rateCode"`
	}
	ExchangeRateResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Хөрвүүлэх валют
		FixedCurrCode string `json:"fixedCurrCode"`
		// Тоо ширхэг
		FixedCurrUnits float64 `json:"fixedCurrUnits"`
		IsRateLatest   string  `json:"isRateLatest"`
		// Хөрвөсөн валют
		VarCurrCode string `json:"varCurrCode"`
		// Тооцоолол
		VarCurrUnits float64 `json:"varCurrUnits"`
	}
)
