package dimets

import (
	"regexp"
	"slices"
	"strings"
	"time"
)

type DimeTransactionLogResponse struct {
	Fee          []DimeTransactionFee
	DividendLogs []DimeTransactionDividend
	StockLogs    []DimeTransactionStock
}
type DimeTransactionRequest struct {
	fee          []DimeFeeTransaction
	dividendLogs []DimeDividendTransaction
	stockLogs    []DimeStockTransaction
}

var (
	pattern = `\d{1,2}\s[A-Z][a-z]{2}\s\d{4}\s-\s\d{2}:\d{2}:\d{2}\s(AM|PM)`
	re      = regexp.MustCompile(pattern)
)

func splitWithDate(text string) []string {
	locs := re.FindAllStringIndex(text, -1)
	if len(locs) == 0 {
		return []string{text}
	}

	transactions := make([]string, len(locs))
	lastPos := 0

	for i, loc := range locs {
		transactions[i] = text[lastPos:loc[1]]
		lastPos = loc[1]
	}

	return transactions
}

type DimeTransaction interface {
	GetExecutedDate() time.Time
}

type LogT interface {
	DimeTransaction
	DimeTransactionFee | DimeTransactionDividend | DimeTransactionStock
}

func (f DimeTransactionFee) GetExecutedDate() time.Time      { return f.ExecutedDate }
func (d DimeTransactionDividend) GetExecutedDate() time.Time { return d.ExecutedDate }
func (s DimeTransactionStock) GetExecutedDate() time.Time    { return s.ExecutedDate }

func sortLogsDescending[T LogT](logs []T) {
	slices.SortFunc(logs, func(a, b T) int {
		return b.GetExecutedDate().Compare(a.GetExecutedDate())
	})
}
func ReadToJson(texts []string) (any, error) {
	results := DimeTransactionLogResponse{Fee: []DimeTransactionFee{}, DividendLogs: []DimeTransactionDividend{}, StockLogs: []DimeTransactionStock{}}
	for _, text := range texts {
		transactionStr := splitWithDate(text)
		transactions, err := parseTextToTransactionReq(transactionStr)
		if err != nil {
			return nil, err
		}
		json, err := newDimeSingleTransactions(transactions)
		if err != nil {
			return nil, err
		}
		results.Fee = append(results.Fee, json.Fee...)
		results.DividendLogs = append(results.DividendLogs, json.DividendLogs...)
		results.StockLogs = append(results.StockLogs, json.StockLogs...)
	}
	sortLogsDescending(results.StockLogs)
	sortLogsDescending(results.Fee)
	sortLogsDescending(results.DividendLogs)
	return &results, nil
}
func parseTextToTransactionReq(transactions []string) (*DimeTransactionRequest, error) {
	n := len(transactions)
	stockLogs := make([]DimeStockTransaction, 0, n)
	feeLogs := make([]DimeFeeTransaction, 0, n)
	dividendLogs := make([]DimeDividendTransaction, 0, n)
	for _, t := range transactions {
		switch {
		case strings.Contains(t, "Sell"), strings.Contains(t, "Buy"):
			stockLogs = append(stockLogs, NewDimeTransactionStock(t))
		case strings.Contains(t, "TAF"):
			feeLogs = append(feeLogs, NewDimeTransactionFee(t))
		case strings.Contains(t, "Dividend"):
			dividendLogs = append(dividendLogs, NewDimeTransactionDividend(t))
		}
	}
	return &DimeTransactionRequest{
		stockLogs:    stockLogs,
		fee:          feeLogs,
		dividendLogs: dividendLogs,
	}, nil
}

func newDimeSingleTransactions(transactions *DimeTransactionRequest) (*DimeTransactionLogResponse, error) {
	res := &DimeTransactionLogResponse{
		Fee:          make([]DimeTransactionFee, len(transactions.fee)),
		DividendLogs: make([]DimeTransactionDividend, len(transactions.dividendLogs)),
		StockLogs:    make([]DimeTransactionStock, len(transactions.stockLogs)),
	}

	for index, dl := range transactions.dividendLogs {
		result, err := dl.ToJson()
		if err != nil {
		}
		res.DividendLogs[index] = *result
	}

	for index, dl := range transactions.fee {
		result, err := dl.ToJson()
		if err != nil {
		}
		res.Fee[index] = *result
	}

	for index, dl := range transactions.stockLogs {
		result, err := dl.ToJson()
		if err != nil {
		}
		res.StockLogs[index] = *result
	}

	return res, nil
}
