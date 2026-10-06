package model

import "encoding/json"

type (
	TransactionReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Шилжүүлэг хийх банкны код.
		// Лавлах төрөл: BANK
		// 8.2 банк хоорондын гүйлгээнд заавал
		BankCode string `json:"bankCode,omitempty"`
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

	TransactionSelfReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Гүйгээний утга
		Remarks string `json:"remarks" validate:"required"`
		// Гүйлгээний төрөл. Үүнд: - TSF (шилжүүлэг) - PMT (төлбөр) - PRC (худалдаа)
		Type string `json:"type" validate:"required"`
		// Тухайн холбогдож буй систем (гуравдагч систем) дээр хийгдсэн гүйлгээний дугаар
		RefCode string `json:"refCode" validate:"required"`
		// Шилжүүлэгчийн мэдээлэл агуулсан объект
		Initiator TransactionObject `json:"initiator" validate:"required"`
		// Хүлээн авагч талын мэдээлэл агуулсан массив утга
		Receives []TransactionObject `json:"receives" validate:"required"`
	}

	TransactionSelfResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Амжилттай гүйлгээний тоо
		SuccessCount int `json:"successCount"`
		// Амжилтгүй гүйлгээний тоо
		FailedCount int `json:"failedCount"`
		// Нийт гүйлгээний тоо
		TotalCount int `json:"totalCount"`
		// хариу мэдэгдлийн жагсаалт
		Part []TransactionPart `json:"part"`
	}
	TransactionPart struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Гүйлгээ хийгдсэн дансны дугаар
		AccountID string `json:"accountId"`
		// Гүйлгээний дугаар
		TransactionID string `json:"tranId"`
		// Гүйлгээний дугаар. Банк хоорондын гүйлгээ дээр хоосон биш байна
		OrderID string `json:"orderId"`
		// Гүйлгээ хийгдсэн огноо. YYYY-MM-DD HH:MM:SS
		TransactionDate string `json:"tranDate"`
		// Тухайн гүйлгээний төлөв. SUCCESS - амжилттай, FAILED - амжилтгүй
		Status string `json:"status"`
	}

	TransactionRefundReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Голомт банк дахь гүйлгээний дугаар. Анх гүйлгээ хийх үед хариу дээр ирж байна.
		TransactionCode string `json:"tranCode" validate:"required"`
		// Гүйлгээ хийсэн огноо. ISO DATE (yyyy-MM-dd) форматтай байна.
		TransactionDate string `json:"tranDate" validate:"required"`
		// Гүйлгээг буцаасан шалтгаан.
		Reason string `json:"reason" validate:"required"`
	}

	TransactionRefundResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Төлөв. SUCCESS - амжилттай, FAILED - амжилтгүй
		Status string `json:"status"`
		// Тайлбар
		Message string `json:"message"`
	}

	TransactionCheckReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNo string `json:"registerNo" validate:"required"`
		// Гуравдагч систем дээрх гүйлгээний дугаар. Анх гүйлгээ хийх үед Голомт Банк руу илгээсэн байх.
		TransactionCode string `json:"tranCode" validate:"required"`
		// Гүйлгээ хийсэн огноо. ISO DATE (yyyy-MM-dd) форматтай байна.
		TransactionDate string `json:"tranDate" validate:"required"`
		// Арилжааны банкны код
		BankCode string `json:"bankCode" validate:"required"`
	}

	TransactionCheckResp struct {
		// Хүсэлтийн дугаар
		RequestID string `json:"requestId"`
		// Голомт банк дахь гүйлгээний дугаар
		TransactionID string `json:"tranId"`
		// Гүйлгээ хийгдсэн огноо.
		TransactionDate string `json:"tranDate"`
		// Гүйлгээний төлөв. Жишээ: POS
		TransactionStatus string `json:"tranStatus"`
	}

	TransactionBatchReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Багц гүйлгээний утга
		Remarks string `json:"remarks" validate:"required"`
		// гүйлгээний мэдээллийг агуулсан массив байна
		Transactions []TransactionSelfReq `json:"transactions" validate:"required"`
	}
	TransactionBatchResp struct {
		// Багц гүйлгээний лавлах дугаар
		RequestID string `json:"requestId"`
		// Гүйлгээг хүлээж авсан мэдээлэл
		Message string `json:"message"`
	}

	TransactionBatchCheckReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Багц гүйлгээний лавлах дугаар
		RequestID string `json:"requestId" validate:"required"`
	}

	TransactionBatchCheckResp struct {
		// Хүсэлтийн лавлах дугаар
		RequestID string `json:"requestId"`
		// Багц гүйлгээний мэдээлэл агуулсан массив байна
		List       []TransactionBatchCheckData `json:"list"`
		Pagination struct {
			// Нийт хуудасны тоо
			TotalPages int `json:"totalPages"`
			// Нийт бичлэгийн тоо
			TotalElements int `json:"totalElements"`
			// Хэд дэх бичлэгийг буцаасан тоо
			NumberOfElements int `json:"numberOfElements"`
		} `json:"pagination"`
	}

	TransactionBatchCheckData struct {
		// Байгууллагын регистрийн дугаар
		RegisterNumber string `json:"regNum"`
		// Шилжүүлэгчийн дансны дугаар
		IniAccountNo string `json:"iniAcctNo"`
		// Шилжүүлэгчийн дансны нэр
		IniAccountName string `json:"iniAcctName"`
		// Хүлээн авагч талын дансны дугаар
		RecAccountNo string `json:"recAcctNo"`
		// Хүлээн авагч талын дансны нэр
		RecAccountName string `json:"recAcctName"`
		// Гүйлгээний дүн
		TransactionAmount float64 `json:"tranAmt"`
		// Гүйлгээний валют
		TransactionCurrency string `json:"tranCrn"`
		// Гүйлгээний утга
		Particular string `json:"particular"`
		// Багц гүйлгээний утга
		Remarks string `json:"remarks"`
		// Гүйлгээний төлөв
		TransactionStatus string `json:"tranStatus"`
		// Үндсэн систем дээрх гүйлгээний дугаар
		CoreTransactionID string `json:"coreTranId"`
		// Үндсэн систем дээр гүйлгээ хийгдсэн огноо
		CoreTransactionDate string `json:"coreTranDate"`
		// Хэрвээ гүйлгээ нь амжилтгүй болсон үед шалтгаан нь байна
		CoreReason string `json:"coreReason"`
		// Амжилттай болсон огноо
		PostedDate string `json:"postedDate"`
		// Шилжүүлэг хийгдсэн банкны код
		BankCode string `json:"bankCode"`
		// Багц гүйлгээний дугаар
		TransactionRefNumber string `json:"tranRefNum"`
	}

	TransactionConfirmReq struct {
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

	TransactionConfirmResp struct {
		// Гүйлгээний дугаар
		TransactionID string `json:"tranId"`
		// Гүйлгээ гарсан огноо
		TransactionDate string `json:"tranDate"`
		// Гүйлгээний төлөв. Жишээ: POS
		TransactionStatus string `json:"tranStatus"`
	}

	// Багц гүйлгээ файлаар хийх үеийн хүсэлт
	// TransactionBatchFileReq struct {
	// 	// Харилцагчийн регистрийн дугаар
	// 	RegisterNumber string `json:"registerNumber" validate:"required"`
	// 	// Багц гүйлгээний зориулалт. SAL, PKG
	// 	FileCode string `json:"fileCode" validate:"required"`
	// 	// 30 тэмдэгтээс ихгүй гүйлгээ хийх утга
	// 	Remarks string `json:"remarks" validate:"required"`
	// 	// Гүйлгээний мэдээллийг агуулсан файл
	// 	File string `json:"file" validate:"required"`
	// }

	// Багц гүйлгээ файлаар хийх үеийн оролт
	TransactionBatchFileInput struct {
		RegisterNo string `json:"registerNo" validate:"required"`
		// Багц гүйлгээний зориулалт. SAL, PKG
		FileCode string `json:"fileCode" validate:"required"`
		// 30 тэмдэгтээс ихгүй гүйлгээ хийх утга
		Remarks string `json:"remarks" validate:"required"`
		// Гүйлгээний мэдээллийг агуулсан файл
		File []TransactionBatchDataInput `json:"file" validate:"required"`
	}

	// Багц гүйлгээ файлаар хийх үеийн хариу
	TransactionBatchFileResp struct {
		// Хүсэлтийн лавлах дугаар
		RequestID string `json:"requestId"`
		// Файл хуулагдсан тухай мэдээлэл харуулна
		Message string `json:"message"`
	}

	// Багц гүйлгээний файл /JSON/
	TransactionBatchDataInput struct {
		// Голомт банкны дансны дугаар.
		Foracid string `json:"foracid"`
		// Гүйлгээ хийх валют код. MNT, USD etc
		TranCrncyCode string `json:"tran_crncy_code"`
		// Гэрээ байгуулсан салбар эсвэл 870
		SolId string `json:"sol_id"`
		// C – гүйлгээ хүлээн авагч тал
		//
		// D – гүйлгээ илгээгч тал
		PartTranType string `json:"part_tran_type"`
		// Гүйлгээний мөнгөн дүн. 1000.00, 7200.00 etc
		TranAmt string `json:"tran_amt"`
		// Гүйлгээний утга
		TransactionParticulars string `json:"transaction_particulars"`
		// MNT руу хөрвүүлсэн дүн байх ба төгрөгөөр гүйлгээ хийсэн үед TRAN_AMT - тай ижил байна
		RefAmt string `json:"ref_amt"`
		// Тогтмол MNT байна
		RefCcy string `json:"ref_ccy"`
		// Арилжаатай гүйлгээ хийх үед банкнаас өгсөн код байна. MNT бол BOMR тогтмол байна
		RateCode string `json:"rate_code"`
		// Ханшны дүн
		Rate string `json:"rate"`
		// Гүйлгээ хийх огноо. Формат: DD/MM/YYYY
		ValueDate string `json:"value_date"`
		// Гүйлгээ хийгдсэн огноо. Формат: DD/MM/YYYY
		GlDate string `json:"gl_date"`
	}

	// 8.4. Гаалийн гүйлгээ хийх
	CustomsPayReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// “Нэхэмжлэх / Цахим төлбөрийн даалгавар”-ын дугаар
		InvoiceNo string `json:"invoiceNo" validate:"required"`
		// Шилжүүлэгчийн нэр
		AccountName string `json:"acctName" validate:"required"`
		// Шилжүүлэгчийн дансны дугаар
		AccountNo string `json:"acctNo" validate:"required"`
		// Гүйлгээний дүн агуулсан объект
		Amount AmountDetail `json:"amount" validate:"required"`
	}
	CustomsPayResp struct {
		// Шилжүүлгийн Id
		Status string `json:"status"`
		// Мессеж
		Message string `json:"message"`
	}

	// 8.5. Татварын гүйлгээ хийх
	TaxPayReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Гүйлгээний утга
		Remarks string `json:"remarks" validate:"required"`
		// “Нэхэмжлэх / Цахим төлбөрийн даалгавар”-ын дугаар
		InvoiceNo string `json:"invoiceNo" validate:"required"`
		// Татварын төрлийн код
		TaxTypeCode string `json:"taxTypeCode" validate:"required"`
		// Шилжүүлэгчийн мэдээлэл агуулсан объект
		Initiator TransactionObject `json:"initiator" validate:"required"`
	}
	TaxPayResp struct {
		// Шилжүүлгийн Id
		TransactionID string `json:"tranId"`
		// Шилжүүлсэн огноо
		TransactionDate string `json:"tranDate"`
		// Шилжүүлгийн статус
		TransactionStatus string `json:"tranStatus"`
	}

	// Нэхэмжлэхийн хөрөнгийн дэлгэрэнгүй
	TaxAsset struct {
		// Хөрөнгийн төрөл. 1=Галт зэвсэг, 2=Үл хөдлөх эд хөрөнгө, 3=Эд хөрөнгө,
		// 4=Мал, амьтан, 5=Тээврийн хэрэгсэл, 6=Эдийн бус хөрөнгө
		ObjType json.Number `json:"objType"`
		// Галт зэвсгийн замагны дугаар
		LockNo string `json:"lockNo"`
		// Галт зэвсгийн гол төмрийн дугаар
		BarrelNo string `json:"barrelNo"`
		// Тээврийн хэрэгслийн улсын бүртгэлийн дугаар
		RegNumber string `json:"regNumber"`
		// Тээврийн хэрэгслийн арлын дугаар
		ChassisNo string `json:"chassisNo"`
		// Тээврийн хэрэгслийн загвар
		Model string `json:"model"`
		// Тээврийн хэрэгслийн марк
		Mark string `json:"mark"`
		// Үл хөдлөх эд хөрөнгийн улсын бүртгэлийн гэрчилгээний дугаар
		CertificateNo string `json:"certificateNo"`
	}

	// 8.12. Татварын төлбөрийн жагсаалт харах TIN
	TaxTINInqReq struct {
		// Татвар төлөгч бүрт олгогдсон дугаар
		TIN string `json:"tin" validate:"required"`
	}
	TaxTINInqData struct {
		// “Нэхэмжлэх / Цахим төлбөрийн даалгавар”-ын дугаар
		InvoiceNo string `json:"invoiceNo"`
		// Татварын төрлийн код
		TaxTypeCode string `json:"taxTypeCode"`
		// Нэхэмжлэхийн үлдэгдэл дүн
		Amount float64 `json:"amount"`
		// Нэхэмжлэхийн дэлгэрэнгүй
		Asset *TaxAsset `json:"asset"`
	}

	// 8.13. Татварын төлбөрийн жагсаалт харах PIN
	TaxPINInqReq struct {
		// УБЕГ-аас олгосон иргэний регистрийн дугаар болон хуулийн этгээдийн 7 оронтой улсын бүртгэлийн дугаар
		PIN string `json:"pin" validate:"required"`
	}
	TaxPINInqData struct {
		// Татвар төлөгчийн дугаар
		TIN string `json:"tin"`
		// Татвар төлөгчийн нэр
		Name string `json:"name"`
	}

	// 8.14. Татварын нэхэмжлэхийн лавлагаа
	TaxInvoiceInqReq struct {
		// “Нэхэмжлэх / Цахим төлбөрийн даалгавар”-ын дугаар
		InvoiceNo string `json:"invoiceNo" validate:"required"`
	}
	TaxInvoiceInqResp struct {
		// “Нэхэмжлэх / Цахим төлбөрийн даалгавар”-ын дугаар
		InvoiceNo string `json:"invoiceNo"`
		// Төлбөрийн дүн
		Amount float64 `json:"amount"`
		// Регистрийн дугаар
		PIN string `json:"pin"`
		// Татвар төлөгчийн дугаар
		TIN string `json:"tin"`
		// Татвар төлөгчийн нэр
		TaxPayerName string `json:"taxPayerName"`
		// Тайлбар
		Description string `json:"description"`
		// Төрийн сан дахь татварын орлогын дансны дугаар
		StateAccount string `json:"stateAccount"`
		// Төрийн сан дахь татварын орлогын дансны нэр
		StateAccountName string `json:"stateAccountName"`
		// IBAN дансны дугаар
		IBAN string `json:"iban"`
		// 1=Нэхэмжлэх, 2=Төлбөрийн даалгавар, 3=Урьдчилгаа төлбөрийн нэхэмжлэх
		InvoiceType json.Number `json:"invoiceType"`
		// Татварын төрлийн код
		TaxTypeCode string `json:"taxTypeCode"`
		// Татварын төрлийн нэр
		TaxTypeName string `json:"taxTypeName"`
		// Татварын албаны код
		BranchCode string `json:"branchCode"`
		// Татварын албаны нэр
		BranchName string `json:"branchName"`
		// Татварын дэд албаны код
		SubBranchCode string `json:"subBranchCode"`
		// Татварын дэд албаны нэр
		SubBranchName string `json:"subBranchName"`
		// Тайлант жил
		Year json.Number `json:"year"`
		// Тайлант үечлэл
		PeriodType string `json:"periodType"`
		// Тайлангийн хугацаа
		Period json.Number `json:"period"`
		// 1=Бүтэн төлнө, 0=Дутуу төлж болно
		PayFull json.Number `json:"payFull"`
		// Илүү төлж болох эсэх
		PayMore json.Number `json:"payMore"`
		// Нэхэмжлэхийн дэлгэрэнгүй
		Asset *TaxAsset `json:"asset"`
		// Шилжүүлгийн статус
		PayStatus string `json:"payStatus"`
	}

	// 8.15. Гаалийн нэхэмжлэхийн лавлагаа
	CustomsInvoiceInqReq struct {
		// “Нэхэмжлэх / Цахим төлбөрийн даалгавар”-ын дугаар
		Invoice string `json:"invoice,omitempty"`
		// Улсын бүртгэлийн дугаар
		Register string `json:"register,omitempty"`
		// “Нэхэмжлэх / Цахим төлбөрийн даалгавар”-ын дугаар
		Barcode string `json:"barcode,omitempty"`
	}
	CustomsInvoiceInqResp struct {
		// “Нэхэмжлэх / Цахим төлбөрийн даалгавар”-ын дугаар
		InvoiceNo string `json:"invoiceNo"`
		// Салбарын дугаар
		BranchID string `json:"branchId"`
		// Салбарын нэр
		BranchName string `json:"branchName"`
		DecDate    string `json:"decDate"`
		// Регистрийн дугаар
		NatID    string  `json:"natId"`
		PymtCode string  `json:"pymtCode"`
		PymtAmt  float64 `json:"pymtAmt"`
		// Гүйлгээний мөнгөн дүн
		TotPymtAmt float64 `json:"totPymtAmt"`
		// Дансны дугаар
		CrAcct string `json:"crAcct"`
		// Дансны нэр
		CrAcctName string `json:"crAcctName"`
		// Нэхэмжлэлийн код
		InvoiceCode string `json:"invoiceCode"`
		// Татвар төлөгчийн нэр
		TaxpayerName string `json:"taxpayerName"`
	}

	// 8.16. Файлаар хийсэн багц гүйлгээний дэлгэрэнгүй татах
	TransactionBatchFileInqReq struct {
		// Харилцагчийн регистрийн дугаар
		RegisterNumber string `json:"registerNumber" validate:"required"`
		// Багц гүйлгээний зориулалт. SAL, PKG
		FileCode string `json:"fileCode" validate:"required"`
		// Багц гүйлгээний лавлах дугаар
		ReferenceNumber string `json:"referenceNumber" validate:"required"`
	}
	TransactionBatchFileInqResp struct {
		// Хүсэлтийн лавлах дугаар
		RequestID string `json:"requestId"`
		// Файлын зориулалт
		FileCode string `json:"fileCode"`
		// Файлын нэр
		FileName string `json:"fileName"`
		// Файлын төрөл
		FileType string `json:"fileType"`
		// Гүйлгээний дугаар. Бүх гүйлгээ нэг л гүйлгээний дугаартай гарна
		TransactionID string `json:"tranId"`
		// Гүйлгээний төлөв
		Status string `json:"status"`
		// Утга
		Remarks string `json:"remarks"`
		// Амжилтгүй болсон шалтгаан
		Reason string `json:"reason"`
		// Гүйлгээний огноо. YYYY-MM-DD
		ValueDate string `json:"valueDate"`
		// Бүртгэсэн огноо. yyyy-MM-ddThh:mm:ss
		CreDate string `json:"creDate"`
	}
)
