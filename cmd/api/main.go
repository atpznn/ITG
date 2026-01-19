package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"runtime/debug"

	dimets "ITG/internal/dime/transaction"
	"ITG/internal/ocr"

	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DimeBody struct {
	Text string `json:"text"`
	Sort bool   `json:"sort"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hi from itg")
}
func MaxQueueMiddleware(queue chan struct{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			select {
			case queue <- struct{}{}:
				defer func() { <-queue }()
				return next(c)
			default:
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					"message": "Server busy, please try again later",
				})
			}
		}
	}
}
func main() {
	ocrService := ocr.NewOCRService(6)
	ocrHandler := ocr.NewOCRHandler(ocrService)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Recover())
	ocrQueue := make(chan struct{}, 5)
	ocrGroup := e.Group("/dime")
	ocrGroup.Use(MaxQueueMiddleware(ocrQueue))
	e.GET("/", hello)
	ocrGroup.POST("", ocrHandler.HandleUpload(dimets.ReadToJson))
	e.GET("clear", func(c echo.Context) error {
		runtime.GC()
		debug.FreeOSMemory()
		return c.String(http.StatusOK, "cleared")
	})
	pprof.Register(e)
	e.Logger.Fatal(e.Start(":8081"))
}
