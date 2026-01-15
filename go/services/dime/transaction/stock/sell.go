package dime_transaction_stock

import (
	"time"

	dime_transaction_model "ITG/services/dime/transaction/model"
)

type DimeSellTransaction struct {
	Text string
}

func (c DimeSellTransaction) ToJson() (*DimeTransactionStock, error) {
	return &DimeTransactionStock{
		BaseDimeTransactionLog: dime_transaction_model.BaseDimeTransactionLog{
			Type:         dime_transaction_model.DimeSellTransactionType,
			Symbol:       "",
			Amount:       12,
			Kind:         dime_transaction_model.DimeTransactionIncome,
			ExecutedDate: time.Now(),
		},
	}, nil
}
