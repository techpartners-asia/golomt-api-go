package openbank

type AccountTypeEnum string

const (
	DepoAccountType AccountTypeEnum = "DEPO" // Хадгаламжийн данс
	OperAccountType AccountTypeEnum = "OPER" // Харилцах данс
	LoanAccountType AccountTypeEnum = "LOAN" // Зээлийн данс
)
