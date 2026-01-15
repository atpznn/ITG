package dime_transaction_fee

import (
	dime_transaction_model "ITG/services/dime/transaction/model"
)

type DimeTransactionFee struct {
	dime_transaction_model.BaseDimeTransactionLog
}
type DimeFeeTransaction interface {
	ToJson() (*DimeTransactionFee, error)
}

func NewDimeTransactionFee(text string) DimeFeeTransaction {
	return DimeTafTransaction{Text: text}
}
