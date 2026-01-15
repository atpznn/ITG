package dime_transaction_fee

import dime_transaction_model "ITG/services/dime/transaction/model"

type DimeTafTransaction struct {
	Text string
}

func (c DimeTafTransaction) ToJson() (*DimeTransactionFee, error) {
	return &DimeTransactionFee{
		BaseDimeTransactionLog: dime_transaction_model.BaseDimeTransactionLog{
			Type: dime_transaction_model.DimeTafTransactionType,
			Kind: dime_transaction_model.DimeTransactionExpense,
		},
	}, nil
}
