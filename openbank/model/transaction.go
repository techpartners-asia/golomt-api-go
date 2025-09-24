package model

type (
	TransactionReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Шилжүүлэг хийх банкны код.
		// Лавлах төрөл: BANK
		BankCode string `json:"bankCode"`
		// Гүйгээний утга
		Remarks string `json:"remarks" validate:"required"`
		// Шилжүүлэгчийн мэдээлэл агуулсан объект
		Initiator TransactionObject `json:"initiator" validate:"required"`
		// Хүлээн авагч талын мэдээлэл агуулсан массив утга
		Receives []TransactionObject `json:"receives" validate:"required"`
	}
	TransactionObject struct {
		// Шилжүүлэгчийн нэр
		AccountName string `json:"acctName" validate:"required"`
		// Шилжүүлэгчийн дансны дугаар
		AccountNo string `json:"acctNo" validate:"required"`
		// Гүйлгээний утга
		Particulars string `json:"particulars" validate:"required"`
		// Гүйлгээний дүн агуулсан объект
		Amount AmountDetail `json:"amount" validate:"required"`
	}

	TransactionResp struct {
		// Баталгаажуулалт хийх URL
		URL string `json:"url"`
	}

	TransactionCheckReq struct {
		// Мерчантад зориулагдаж тусдаа үүсгэсэн код байна
		ClientID string `json:"clientId" validate:"required"`
		// Дахин давтагдахгүйгээр үүсгэсэн нэг удаагийн токен
		State string `json:"state" validate:"required"`
		// Гүйлгээний хүсэлтийг тусгай алгоритмтаар нууцалсан утга
		Scope string `json:"scope" validate:"required"`
		// Гүйлгээний төрөл.
		//
		// INB – голомт банк данс хоорондын гүйлгээ;
		// OSB – Банк хоорондын гүйлгээ
		Type string `json:"type"`
	}

	// TODO: check response
	TransactionCheckResp struct {
		// Гүйлгээний дугаар
		TransactionID string `json:"tranId"`
		// Гүйлгээ гарсан огноо
		TransactionDate string `json:"tranDate"`
		// Гүйлгээний төлөв
		TransactionStatus string `json:"tranStatus"`
	}
)
