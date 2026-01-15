package dime_transaction_dividend

import (
	dime_transaction_model "ITG/services/dime/transaction/model"
	"errors"
	"strings"
)


type DimeDividendIncomeTransaction struct {
	Text string
}
func (b DimeDividendIncomeTransaction) ToJson() (any, error) {
	startIndex := strings.Index(b.Text, "Dividend")
	if startIndex == -1 {
		return nil, errors.New("invalid transaction format: 'Dividend' not found")
	}
	texts := strings.Split(b.Text[startIndex:], "\n")
	if len(texts) < 2 {
		return nil, errors.New("invalid transaction format: insufficient lines")
	}
	return &DimeDividendTransaction{
		DimeTransactionLog: dime_transaction_model.DimeTransactionLog{
			Type: dime_transaction_model.DimeDividendTransactionType,
			Kind: dime_transaction_model.DimeTransactionIncome,
		},

		}, nil
}