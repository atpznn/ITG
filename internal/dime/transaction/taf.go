package dimets

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type DimeTafTransaction struct {
	Text string
}

func (c DimeTafTransaction) ToJson() (*DimeTransactionFee, error) {
	startIndex := strings.Index(c.Text, "TAF")
	if startIndex == -1 {
		return nil, errors.New("invalid transaction format: 'Buy' not found")
	}
	texts := strings.Split(c.Text[startIndex:], "\n")
	if len(texts) < 2 {
		return nil, errors.New("invalid transaction format: insufficient lines")
	}
	amountStr := strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(texts[0], "TAF Fee", ""), "USD", ""))
	amount, err := strconv.ParseFloat(amountStr, 32)
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
	return &DimeTransactionFee{
		BaseDimeTransactionLog: BaseDimeTransactionLog{
			Type:         DimeTafTransactionType,
			Kind:         DimeTransactionExpense,
			ExecutedDate: time,
			Symbol:       "-",
			Amount:       math.Abs(amount),
		},
	}, nil
}
