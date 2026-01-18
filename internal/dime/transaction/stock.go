package dimets

import (
	"strings"
)

type DimeTransactionStock struct {
	BaseDimeTransactionLog
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
