package dime_transaction

import (
	"context"
	"strings"

	dime_transaction_dividend "ITG/services/dime/transaction/dividend"
	dime_transaction_fee "ITG/services/dime/transaction/fee"
	dime_transaction_stock "ITG/services/dime/transaction/stock"

	"golang.org/x/sync/errgroup"
)

type DimeTransactionLogResponse struct {
	fee          []dime_transaction_fee.DimeTransactionFee
	dividendLogs []dime_transaction_dividend.DimeTransactionDividend
	stockLogs    []dime_transaction_stock.DimeTransactionStock
}
type DimeTransactionRequest struct {
	fee          []dime_transaction_fee.DimeFeeTransaction
	dividendLogs []dime_transaction_dividend.DimeDividendTransaction
	stockLogs    []dime_transaction_stock.DimeStockTransaction
}

func ParseTextToTransactionReq(transactions []string) (*DimeTransactionRequest, error) {
	n := len(transactions)
	stockLogs := make([]dime_transaction_stock.DimeStockTransaction, 0, n)
	feeLogs := make([]dime_transaction_fee.DimeFeeTransaction, 0, n)
	dividendLogs := make([]dime_transaction_dividend.DimeDividendTransaction, 0, n)
	for _, t := range transactions {
		switch {
		case strings.Contains(t, "Sell"), strings.Contains(t, "Buy"):
			stockLogs = append(stockLogs, dime_transaction_stock.NewDimeTransactionStock(t))
		case strings.Contains(t, "TAF"):
			feeLogs = append(feeLogs, dime_transaction_fee.NewDimeTransactionFee(t))
		case strings.Contains(t, "Dividend"):
			dividendLogs = append(dividendLogs, dime_transaction_dividend.NewDimeTransactionDividend(t))
		}
	}
	return &DimeTransactionRequest{
		stockLogs:    stockLogs,
		fee:          feeLogs,
		dividendLogs: dividendLogs,
	}, nil
}

func NewDimeMultipleTransactions(ctx context.Context, req *DimeTransactionRequest) (*DimeTransactionLogResponse, error) {
	g, _ := errgroup.WithContext(ctx)

	var fees []dime_transaction_fee.DimeTransactionFee
	var dividends []dime_transaction_dividend.DimeTransactionDividend
	var stocks []dime_transaction_stock.DimeTransactionStock

	g.Go(func() error {
		fees = make([]dime_transaction_fee.DimeTransactionFee, 0, len(req.fee))
		for _, itm := range req.fee {
			if val, err := itm.ToJson(); err == nil {
				fees = append(fees, *val)
			}
		}
		return nil
	})

	g.Go(func() error {
		dividends = make([]dime_transaction_dividend.DimeTransactionDividend, 0, len(req.dividendLogs))
		for _, itm := range req.dividendLogs {
			if val, err := itm.ToJson(); err == nil {
				dividends = append(dividends, *val)
			}
		}
		return nil
	})

	g.Go(func() error {
		stocks = make([]dime_transaction_stock.DimeTransactionStock, 0, len(req.stockLogs))
		for _, itm := range req.stockLogs {
			if val, ok := itm.ToJson(); ok == nil {
				stocks = append(stocks, *val)
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &DimeTransactionLogResponse{
		fee:          fees,
		dividendLogs: dividends,
		stockLogs:    stocks,
	}, nil
}

func NewDimeSingleTransactions(transactions *DimeTransactionRequest) (*DimeTransactionLogResponse, error) {
	res := &DimeTransactionLogResponse{
		fee:          make([]dime_transaction_fee.DimeTransactionFee, len(transactions.fee)),
		dividendLogs: make([]dime_transaction_dividend.DimeTransactionDividend, len(transactions.dividendLogs)),
		stockLogs:    make([]dime_transaction_stock.DimeTransactionStock, len(transactions.stockLogs)),
	}

	for index, dl := range transactions.dividendLogs {
		result, err := dl.ToJson()
		if err != nil {
		}
		res.dividendLogs[index] = *result
	}

	for index, dl := range transactions.fee {
		result, err := dl.ToJson()
		if err != nil {
		}
		res.fee[index] = *result
	}

	for index, dl := range transactions.stockLogs {
		result, err := dl.ToJson()
		if err != nil {
		}
		res.stockLogs[index] = *result
	}

	return res, nil
}
