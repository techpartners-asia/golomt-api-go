package socialpay

import "github.com/techpartners-asia/golomt-api-go/utils"

func mapToCommonResponse(resp map[string]interface{}) (response CommonResponse) {
	response = CommonResponse{
		Description: utils.GetValidString(resp["desc"]),
		Status:      utils.GetValidString(resp["status"]),
	}
	return
}

func mapToInvoiceResponse(resp map[string]interface{}) (response InvoiceResponse) {
	response = InvoiceResponse{
		ApprovalCode:        utils.GetValidString(resp["approval_code"]),
		Amount:              utils.GetValidFloat(resp["amount"]),
		CardNumber:          utils.GetValidString(resp["card_number"]),
		ResponseDescription: utils.GetValidString(resp["resp_desc"]),
		ResponseCode:        utils.GetValidString(resp["resp_code"]),
		Terminal:            utils.GetValidString(resp["terminal"]),
		Invoice:             utils.GetValidString(resp["invoice"]),
		Checksum:            utils.GetValidString(resp["checksum"]),
	}
	return
}

func mapToSettlementResponse(resp map[string]interface{}) (response SettlementResponse) {
	response = SettlementResponse{
		Amount: utils.GetValidFloat(resp["amount"]),
		Count:  resp["count"].(int),
		Status: utils.GetValidString(resp["status"]),
	}
	return
}

func mapToErrorResponse(resp map[string]interface{}) (response ErrorResponse) {
	return ErrorResponse{
		ErrorDescription: utils.GetValidString(resp["errorDesc"]),
		ErrorType:        utils.GetValidString(resp["errorType"]),
	}
}
