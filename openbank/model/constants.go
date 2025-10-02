package model

type AccountTypeEnum string

const (
	DepoAccountType AccountTypeEnum = "DEPO" // Хадгаламжийн данс
	OperAccountType AccountTypeEnum = "OPER" // Харилцах данс
	LoanAccountType AccountTypeEnum = "LOAN" // Зээлийн данс
)

func (a AccountTypeEnum) String() string {
	return string(a)
}

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

func (a AccountStatusEnum) String() string {
	return string(a)
}

type CardMerchantStatementTypeEnum string

const (
	CardMerchantStatementTypePOSF    CardMerchantStatementTypeEnum = "POSF"    // POSF
	CardMerchantStatementTypePOSS    CardMerchantStatementTypeEnum = "POSS"    // POSS
	CardMerchantStatementTypeFAILED  CardMerchantStatementTypeEnum = "FAILED"  // FAILED
	CardMerchantStatementTypeSUCCESS CardMerchantStatementTypeEnum = "SUCCESS" // SUCCESS
	CardMerchantStatementTypeTNKSUC  CardMerchantStatementTypeEnum = "TNKSUC"  // TNKSUC
	CardMerchantStatementTypeTNKFAL  CardMerchantStatementTypeEnum = "TNKFAL"  // TNKFAL
	CardMerchantStatementTypeSTTL    CardMerchantStatementTypeEnum = "STTL"    // STTL
)

func (c CardMerchantStatementTypeEnum) String() string {
	return string(c)
}
