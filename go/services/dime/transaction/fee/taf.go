package dime_transaction_fee

import dime_transaction_model "ITG/services/dime/transaction/model"

type DimeTafTransaction struct {
	Text string
}

func (c DimeTafTransaction) ToJson() (any, error) {
	return &DimeTransactionFee{
		DimeTransactionLog: dime_transaction_model.DimeTransactionLog{
			Type: dime_transaction_model.DimeTafTransactionType,
			Kind: dime_transaction_model.DimeTransactionExpense,
		},
	}, nil
}
