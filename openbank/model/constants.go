package model

type AccountTypeEnum string

const (
	DepoAccountType AccountTypeEnum = "DEPO" // Хадгаламжийн данс
	OperAccountType AccountTypeEnum = "OPER" // Харилцах данс
	LoanAccountType AccountTypeEnum = "LOAN" // Зээлийн данс
)

type FreezeStatusCodeEnum string

const (
	FreezeStatusCodeTotal  FreezeStatusCodeEnum = "T" // Нийт
	FreezeStatusCodeCredit FreezeStatusCodeEnum = "C" // Кредит
	FreezeStatusCodeDebit  FreezeStatusCodeEnum = "D" // Дебит
)

type AccountStatusEnum string

const (
	AccountStatusActive   AccountStatusEnum = "A" // Active
	AccountStatusInactive AccountStatusEnum = "I" // Inactive
	AccountStatusDormant  AccountStatusEnum = "D" // Dormant
)
