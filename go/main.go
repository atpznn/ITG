package main

import (
	"ITG/services"
	dime_transaction "ITG/services/dime/transaction"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
type DimeBody struct {
    Text string `json:"text"`
}
func handler(c echo.Context) error{
	
	u := new(DimeBody)
    // Bind จะอ่าน Body และใส่ค่าลงใน u
    if err := c.Bind(u); err != nil {
        return err
    }
	
	transactions:= services.SplitWithDate(u.Text)
	results:= make([]any,len(transactions))
	for index, transaction := range transactions {
		parser, err := dime_transaction.NewDimeTransaction(transaction)
		if err != nil {
			return c.String(http.StatusBadRequest,err.Error())
		}
		result,err := parser.ToJson()
		if err != nil {
			return c.String(http.StatusBadRequest,err.Error())
		}
		results[index] = result
	}
		return c.JSON(http.StatusOK,results)
}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/dime/text-process",handler)
	e.Logger.Fatal(e.Start(":8081"))
}