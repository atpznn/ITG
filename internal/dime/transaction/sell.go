package dimets

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DimeSellTransaction struct {
	Text string
}

func (c DimeSellTransaction) ToJson() (*DimeTransactionStock, error) {
	startIndex := strings.Index(c.Text, "Sell")
	if startIndex == -1 {
		return nil, errors.New("invalid transaction format: 'Buy' not found")
	}
	texts := strings.Split(c.Text[startIndex:], "\n")
	if len(texts) < 2 {
		return nil, errors.New("invalid transaction format: insufficient lines")
	}
	symbolAndAmount := strings.Fields(strings.TrimSpace(strings.ReplaceAll(texts[0], "Sell", "")))
	if len(symbolAndAmount) < 2 {
		return nil, errors.New("invalid symbol or amount format")
	}
	symbol := symbolAndAmount[0]
	amountStr := symbolAndAmount[1]
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
	priceStr := strings.Replace(texts[1], dateStr, "", 1)
	priceStr = strings.Replace(priceStr, "Executed Price", "", 1)
	priceStr = strings.TrimSpace(priceStr)
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return nil, fmt.Errorf("parse price failed: %w", err)
	}

	return &DimeTransactionStock{
		BaseDimeTransactionLog: BaseDimeTransactionLog{
			Type:         DimeSellTransactionType,
			Symbol:       symbol,
			Amount:       amount * price,
			Kind:         DimeTransactionIncome,
			ExecutedDate: time,
		},
		Shares: amount,
		Price:  price,
	}, nil
}
