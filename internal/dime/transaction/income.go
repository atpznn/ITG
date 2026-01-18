package dimets

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DimeDividendIncomeTransaction struct {
	Text string
}

func (b DimeDividendIncomeTransaction) ToJson() (*DimeTransactionDividend, error) {
	startIndex := strings.Index(b.Text, "Dividend")
	if startIndex == -1 {
		return nil, errors.New("invalid transaction format: 'Dividend' not found")
	}
	texts := strings.Split(b.Text[startIndex:], "\n")
	if len(texts) < 2 {
		return nil, errors.New("invalid transaction format: insufficient lines")
	}
	amountAndSymbolStr := strings.ReplaceAll(strings.ReplaceAll(texts[0], "Dividend", ""), "USD", "")
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
			Type:         DimeDividendTransactionType,
			Kind:         DimeTransactionIncome,
			Symbol:       symbol,
			Amount:       amount,
			ExecutedDate: time,
		},
	}, nil
}
