package main

import (
	"ITG/services"
	dime_transaction "ITG/services/dime/transaction"
	dime_transaction_dividend "ITG/services/dime/transaction/dividend"
	dime_transaction_fee "ITG/services/dime/transaction/fee"
	dime_transaction_stock "ITG/services/dime/transaction/stock"
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/sync/errgroup"
)
type DimeBody struct {
    Text string `json:"text"`
}
func handler(c echo.Context) error {
    u := new(DimeBody)
    if err := c.Bind(u); err != nil {
        return err
    }

    transactions := services.SplitWithDate(u.Text)    
    g, ctx := errgroup.WithContext(c.Request().Context())
    resultChan := make(chan any, len(transactions))

    for _, n := range transactions {
        txText := n 
        
        g.Go(func() error {
            select {
            case <-ctx.Done():
                return ctx.Err()
            default:
            }

            parser, err := dime_transaction.NewDimeTransaction(txText)
            if err != nil {
                return err 
            }

            result, err := parser.ToJson()
            if err != nil {
                return err
            }

            resultChan <- result
            return nil
        })
    }

    if err := g.Wait(); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Process stopped due to: " + err.Error(),
        })
    }

    close(resultChan)

    var finalResults []any
    for r := range resultChan {
        finalResults = append(finalResults, r)
    }
	sort.Slice(finalResults, func(i, j int) bool { 
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
		}}
    return getDate(finalResults[i]).Before(getDate(finalResults[j]))
})
    return c.JSON(http.StatusOK, finalResults)
}
func hello(c echo.Context)error{
	return c.String(http.StatusOK,"hi from itg")

}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/",hello)
	e.POST("/dime/text-process",handler)
	e.Logger.Fatal(e.Start(":8081"))
}