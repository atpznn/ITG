package transaction

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type DimeTransaction interface {
	ToJson() (*DimeTransactionLog, error)
}
type DimeTransactionLog struct {
	Type         string
	Symbol       string
	Amount       float64
	ExecutedDate time.Time
	Shares       float64
	Price        float64
}

func NewDimeTransaction(text string) (DimeTransaction, error) {
	fmt.Println(text)
	if strings.Contains(text, "Sell") {
		return DimeSellTransaction{Text: text}, nil
	}
	if strings.Contains(text, "Buy") {
		return DimeBuyTransaction{Text: text}, nil
	}
	if strings.Contains(text, "TAF") {
		return DimeTafTransaction{Text: text}, nil
	}
	if strings.Contains(text, "Dividend") {
		return DimeDividendTransaction{Text: text}, nil
	}
	return nil, errors.New("Not Found Transaction Type")
}
