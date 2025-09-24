package model

type (
	ErrorResp struct {
		Status       string     `json:"status"`
		Timestamp    string     `json:"timestamp"`
		Message      string     `json:"message"`
		DebugMessage string     `json:"debugMessage"`
		SubErrors    []SubError `json:"subErrors"`
	}
	SubError struct {
		Type string `json:"type"`
		Desc string `json:"desc"`
		Code string `json:"code"`
	}

	OAuthResp struct {
		ClientID     string `json:"clientId"`
		ResponseType string `json:"responseType"`
		RedirectUri  string `json:"redirectUri"`
		Scope        string `json:"scope"`
		State        string `json:"state"`
		RequestID    string `json:"requestId"`
		Url          string `json:"url"`
		XsrvID       string `json:"xsrvId"`
	}

	OpenbankInput struct {
		OrganizationName string `json:"organizationName"` // Бүртгүүлсэн байгууллагын нэр
		Username         string `json:"username"`         // Хэрэглэгчийн нэр
		Password         string `json:"password"`         // Нууц үг
		IvKey            string `json:"ivKey"`            // IV key
		SessionKey       string `json:"sessionKey"`       // Session key
		Url              string `json:"url"`              // URL
		RegisterNo       string `json:"registerNo"`       // Бүртгүүлсэн байгууллагын дугаар
		ClientID         string `json:"clientId"`         // Хэрэглэгчийн ID
	}
)
