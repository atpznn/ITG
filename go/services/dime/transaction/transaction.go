package dime_transaction

import (
	dime_transaction_dividend "ITG/services/dime/transaction/dividend"
	dime_transaction_fee "ITG/services/dime/transaction/fee"
	dime_transaction_model "ITG/services/dime/transaction/model"
	dime_transaction_stock "ITG/services/dime/transaction/stock"
	"errors"
	"strings"
)




func NewDimeTransaction(text string) (dime_transaction_model.DimeTransaction,error) {
	if strings.Contains(text, "Sell") || strings.Contains(text,"Buy"){
		return dime_transaction_stock.NewDimeTransactionStock(text),nil
	}
	if strings.Contains(text, "TAF") {
		return dime_transaction_fee.NewDimeTransactionFee(text),nil
	}
	if strings.Contains(text, "Dividend") {
		return dime_transaction_dividend.NewDimeTransactionDividend(text),nil
	}
	return nil,errors.New("Not Found Transaction Type")
}
