package dime_transaction_model

import (
	"time"
)

type BaseDimeTransactionLog struct {
	Type         DimeTransactionType
	Symbol       string
	Amount       float64
	ExecutedDate time.Time
	Kind         DimeTransactionKine
}
type DimeTransactionKine string

const (
	DimeTransactionExpense DimeTransactionKine = "expense"
	DimeTransactionIncome  DimeTransactionKine = "income"
)

type DimeTransactionType string

const (
	DimeDividendTransactionType    DimeTransactionType = "dividend"
	DimeTaxDividendTransactionType DimeTransactionType = "tax-dividend"
	DimeBuyTransactionType         DimeTransactionType = "buy"
	DimeSellTransactionType        DimeTransactionType = "sell"
	DimeTafTransactionType         DimeTransactionType = "tax-taf"
)
