package dime_transaction_fee

import (
	dime_transaction_model "ITG/services/dime/transaction/model"
)

type DimeTransactionFee struct {
	dime_transaction_model.DimeTransactionLog
}

func NewDimeTransactionFee(text string) dime_transaction_model.DimeTransaction{
	return DimeTafTransaction{Text: text}
}