package transaction_transaction_stock

import (
	dime_transaction_model "ITG/services/dime/transaction/model"
	"time"
)

type DimeSellTransaction struct {
	Text string
}

func (c DimeSellTransaction) ToJson() (any, error) {
	return &DimeTransactionStock{
		DimeTransactionLog: dime_transaction_model.DimeTransactionLog{
			Type: dime_transaction_model.DimeSellTransactionType,
			Symbol: "",
			Amount: 12,
			Kind: dime_transaction_model.DimeTransactionIncome,
			ExecutedDate: time.Now(),
		},
	}, nil
}
