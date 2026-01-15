package dime_transaction_dividend

import (
	"errors"
	"strings"

	dime_transaction_model "ITG/services/dime/transaction/model"
)

type DimeDividendTaxTransaction struct {
	Text string
}

func (b DimeDividendTaxTransaction) ToJson() (*DimeTransactionDividend, error) {
	startIndex := strings.Index(b.Text, "Dividend Withholding Tax")
	if startIndex == -1 {
		return nil, errors.New("invalid transaction format: 'Dividend Withholding Tax' not found")
	}
	texts := strings.Split(b.Text[startIndex:], "\n")
	if len(texts) < 2 {
		return nil, errors.New("invalid transaction format: insufficient lines")
	}
	return &DimeTransactionDividend{
		BaseDimeTransactionLog: dime_transaction_model.BaseDimeTransactionLog{
			Type: dime_transaction_model.DimeTaxDividendTransactionType,
			Kind: dime_transaction_model.DimeTransactionExpense,
		},
	}, nil
}
