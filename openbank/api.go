package openbank

import "net/http"

var (
	AuthApi = API{
		Url:         "/v1/auth/login",
		Method:      http.MethodPost,
		ServiceName: "LGIN",
	}

	StatementApi = API{
		Url:         "/v1/account/operative/statement/}",
		Method:      http.MethodPost,
		ServiceName: "OPERACCTSTA",
	}
	AccountListApi = API{
		Url:         "/v1/account/list",
		Method:      http.MethodPost,
		ServiceName: "ACCTLST",
	}
	AccountTypeInq = API{
		Url:         "/v1/account/type/inq",
		Method:      http.MethodPost,
		ServiceName: "ACCTTYPEINQ",
	}

	AccountBalcApi = API{
		Url:         "/v1/account/balance/inq",
		Method:      http.MethodPost,
		ServiceName: "ACCTBALINQ",
	}

	UtilityRateAPI = API{
		Url:         "/v1/utility/rate/inq",
		Method:      http.MethodPost,
		ServiceName: "RATEINQ",
	}
)
