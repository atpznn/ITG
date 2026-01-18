package dimets

import (
	"strings"
)

type DimeDividendTransaction interface {
	ToJson() (*DimeTransactionDividend, error)
}

type DimeTransactionDividend struct {
	BaseDimeTransactionLog
}

func NewDimeTransactionDividend(text string) DimeDividendTransaction {
	if strings.Contains(text, "Dividend Withholding Tax") {
		return DimeDividendTaxTransaction{Text: text}
	}
	return DimeDividendIncomeTransaction{
		Text: text,
	}
}
