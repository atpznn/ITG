package dime_transaction_stock

import (
	"strings"

	dime_transaction_model "ITG/services/dime/transaction/model"
)

type DimeTransactionStock struct {
	dime_transaction_model.BaseDimeTransactionLog
	Shares float64
	Price  float64
}
type DimeStockTransaction interface {
	ToJson() (*DimeTransactionStock, error)
}

func NewDimeTransactionStock(text string) DimeStockTransaction {
	if strings.Contains(text, "Buy") {
		return DimeBuyTransaction{Text: text}
	}
	return DimeSellTransaction{Text: text}
}
