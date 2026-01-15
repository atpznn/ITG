package dime_transaction_stock

import (
	dime_transaction_model "ITG/services/dime/transaction/model"
	"strings"
)


type DimeTransactionStock struct {
	dime_transaction_model.DimeTransactionLog
	Shares       float64
	Price        float64
}
func NewDimeTransactionStock(text string) dime_transaction_model.DimeTransaction {
	if strings.Contains(text,"Buy"){
		return DimeBuyTransaction{Text: text}
	}
	return DimeSellTransaction{Text: text}
}