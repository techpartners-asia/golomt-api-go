package openbank

type AccountTypeEnum string

const (
	DepoAccountType AccountTypeEnum = "DEPO"
	OperAccountType AccountTypeEnum = "OPER"
	LoanAccountType AccountTypeEnum = "LOAN"
)
