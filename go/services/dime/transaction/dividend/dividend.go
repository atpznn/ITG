package dime_transaction_dividend

import (
	"strings"

	dime_transaction_model "ITG/services/dime/transaction/model"
)

type DimeDividendTransaction interface {
	ToJson() (*DimeTransactionDividend, error)
}

type DimeTransactionDividend struct {
	dime_transaction_model.BaseDimeTransactionLog
}

func NewDimeTransactionDividend(text string) DimeDividendTransaction {
	if strings.Contains(text, "Dividend Withholding Tax") {
		return DimeDividendTaxTransaction{Text: text}
	}
	return DimeDividendIncomeTransaction{
		Text: text,
	}
}
