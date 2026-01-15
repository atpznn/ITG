package main

import (
	"net/http"
	"sort"
	"time"

	"ITG/services"
	dime_transaction "ITG/services/dime/transaction"
	dime_transaction_dividend "ITG/services/dime/transaction/dividend"
	dime_transaction_fee "ITG/services/dime/transaction/fee"
	dime_transaction_stock "ITG/services/dime/transaction/stock"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DimeBody struct {
	Text string `json:"text"`
	Sort bool   `json:"sort"`
}

func processing(texts []string) (any, error) {
	// parser, err := dime_transaction.NewDimeTransaction(text)
	// if err != nil {
	// return nil, err
	// }
	// result, err := parser.ToJson()
	// if err != nil {
	// return nil, err
	// }
	// return result, nil
	return nil, nil
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hi from itg")
}

type Param struct {
	c            echo.Context
	transactions []string
	processing   func(text []string) (*dime_transaction.DimeTransactionLogResponse, error)
	sort         bool
}

func sortResult(result []any) {
	sort.Slice(result, func(i, j int) bool {
		getDate := func(item any) time.Time {
			switch v := item.(type) {
			case *dime_transaction_dividend.DimeDividendTransaction:
				return v.ExecutedDate
			case *dime_transaction_fee.DimeTransactionFee:
				return v.ExecutedDate
			case *dime_transaction_stock.DimeTransactionStock:
				return v.ExecutedDate
			default:
				return time.Time{}
			}
		}
		return getDate(result[i]).Before(getDate(result[j]))
	})
}

func singleProcess(param Param) error {
	results := make([]any, len(param.transactions))
	result, err := processing(param.transactions)
	if err != nil {
		return param.c.String(http.StatusInternalServerError, err.Error())
	}
	if param.sort {
		sortResult(results)
	}
	return param.c.JSON(http.StatusOK, result)
}

func multitaskProcess(param Param) error {
	// numWorkers := runtime.NumCPU()
	// g, ctx := errgroup.WithContext(param.c.Request().Context())
	// resultChan := make(chan any, numWorkers)
	// for _, transaction := range param.transactions {
	// 	g.Go(func() error {
	// 		select {
	// 		case <-ctx.Done():
	// 			return ctx.Err()
	// 		default:
	// 		}
	// 		result, err := processing(transaction)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		resultChan <- result
	// 		return nil
	// 	})
	// }
	// go func() {
	// 	g.Wait()
	// 	close(resultChan)
	// }()

	// results := []any{}
	// for r := range resultChan {
	// 	results = append(results, r)
	// }

	// if err := g.Wait(); err != nil {
	// 	return param.c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	// }
	// if param.sort {
	// 	sortResult(results)
	// }
	return param.c.JSON(http.StatusOK, 1)
}

func compose(
	processing func(text []string) (*dime_transaction.DimeTransactionLogResponse, error),
) func(c echo.Context) error {
	return func(c echo.Context) error {
		u := new(DimeBody)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		transactions := services.SplitWithDate(u.Text)
		param := Param{transactions: transactions, processing: processing, c: c, sort: u.Sort}
		return singleProcess(param)
		// return req(param)
	}
}

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)
	e.POST("single/dime/text-process", compose(dime_transaction.NewDimeSingleTransactions))
	e.POST("multi/dime/text-process", compose(dime_transaction.NewDimeMultipleTransactions))
	e.Logger.Fatal(e.Start(":8081"))
}
