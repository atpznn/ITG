package dimets

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DimeBuyTransaction struct {
	Text string
}

func (b DimeBuyTransaction) ToJson() (*DimeTransactionStock, error) {
	startIndex := strings.Index(b.Text, "Buy")
	if startIndex == -1 {
		return nil, errors.New("invalid transaction format: 'Buy' not found")
	}
	texts := strings.Split(b.Text[startIndex:], "\n")
	if len(texts) < 2 {
		return nil, errors.New("invalid transaction format: insufficient lines")
	}
	symbolAndPrice := strings.Fields(strings.TrimSpace(strings.ReplaceAll(texts[0], "Buy", "")))
	if len(symbolAndPrice) < 2 {
		return nil, errors.New("invalid symbol or amount format")
	}
	symbol := symbolAndPrice[0]
	priceStr := symbolAndPrice[1]
	price, err := strconv.ParseFloat(priceStr, 32)
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
	amountStr := strings.Replace(texts[1], dateStr, "", 1)
	amountStr = strings.Replace(amountStr, "Executed Price", "", 1)
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return nil, fmt.Errorf("parse price failed: %w", err)
	}

	return &DimeTransactionStock{
		BaseDimeTransactionLog: BaseDimeTransactionLog{
			Type:         DimeBuyTransactionType,
			Symbol:       symbol,
			Kind:         DimeTransactionExpense,
			ExecutedDate: time,
			Amount:       amount,
		},
		Shares: amount / price,
		Price:  price,
	}, nil
}
