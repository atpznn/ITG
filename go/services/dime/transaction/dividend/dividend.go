package dime_transaction_dividend

import (
	dime_transaction_model "ITG/services/dime/transaction/model"
	"strings"
)
type DimeDividendTransaction struct {
	dime_transaction_model.DimeTransactionLog
}
func NewDimeTransactionDividend(text string) dime_transaction_model.DimeTransaction{
	if strings.Contains(text,"Dividend Withholding Tax") {
		return DimeDividendTaxTransaction{Text: text}
	}
	return DimeDividendIncomeTransaction{
		Text: text,
	}
}