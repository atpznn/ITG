package dimets

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
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
	amountAndSymbolStr := strings.ReplaceAll(strings.ReplaceAll(texts[0], "Dividend Withholding Tax", ""), "USD", "")
	amountAndSymbol := strings.Fields(amountAndSymbolStr)
	if len(amountAndSymbol) < 2 {
		return nil, errors.New("invalid transaction format: insufficient lines")
	}
	symbol := amountAndSymbol[0]
	amount, err := strconv.ParseFloat(amountAndSymbol[1], 32)
	if err != nil {
		return nil, errors.New("can't parse amout to float")
	}
	dateStr := re.FindString(texts[1])
	if dateStr == "" {
		return nil, errors.New("timestamp not found in second line")
	}
	layout := "2 Jan 2006 - 03:04:05 PM"
	time, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, fmt.Errorf("parse time failed: %w", err)
	}
	return &DimeTransactionDividend{
		BaseDimeTransactionLog: BaseDimeTransactionLog{
			Type:         DimeTaxDividendTransactionType,
			Kind:         DimeTransactionExpense,
			Symbol:       symbol,
			Amount:       math.Abs(amount),
			ExecutedDate: time,
		},
	}, nil
}
