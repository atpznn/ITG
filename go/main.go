package main

/*
#include <malloc.h>
*/
import "C"

import (
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"

	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/otiai10/gosseract"
)

type DimeBody struct {
	Text string `json:"text"`
	Sort bool   `json:"sort"`
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hi from itg")
}

var (
	ocrClient *gosseract.Client
	mutex     sync.Mutex
)
var sem = make(chan struct{}, 2)
var max_queue = 6
var waitQueue = make(chan struct{}, max_queue)

func ocrHandlerSafe(c echo.Context) error {
	select {
	case waitQueue <- struct{}{}:
		defer func() {
			<-waitQueue
		}()
	default:
		return c.JSON(http.StatusTooManyRequests, map[string]string{
			"error": fmt.Sprintf("คิวเต็ม %d / %d กรุณาลองใหม่ในอีก 30 วินาที", len(waitQueue), max_queue),
		})
	}
	mutex.Lock()
	defer mutex.Unlock()
	reader, err := c.Request().MultipartReader()
	if err != nil {
		return err
	}

	var results []string

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if part.FormName() != "images" {
			continue
		}

		tmpFile, err := os.CreateTemp("", "ocr-*.png")
		if err != nil {
			return err
		}
		tmpPath := tmpFile.Name()
		_, err = io.Copy(tmpFile, part)
		tmpFile.Close()
		ocrClient.SetImage(tmpPath)
		text, _ := ocrClient.Text()
		results = append(results, text)
		os.Remove(tmpFile.Name())
	}
	// ocrClient.Close()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"results": results,
	})
}

func ocrHandler(c echo.Context) error {
	mutex.Lock()
	defer mutex.Unlock()
	reader, err := c.Request().MultipartReader()
	if err != nil {
		return err
	}

	var results []string

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if part.FormName() != "images" {
			continue
		}

		tmpFile, err := os.CreateTemp("", "ocr-*.png")
		if err != nil {
			return err
		}
		tmpPath := tmpFile.Name()
		_, err = io.Copy(tmpFile, part)
		tmpFile.Close()
		ocrClient.SetImage(tmpPath)
		text, _ := ocrClient.Text()
		results = append(results, text)
		os.Remove(tmpFile.Name())
	}
	// ocrClient.Close()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"results": results,
	})
}
func main() {
	func() {
		ocrClient = gosseract.NewClient()
		ocrClient.SetLanguage("eng")
	}()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Recover())
	e.GET("/", hello)
	e.POST("ocr-single-safe", ocrHandlerSafe)
	e.POST("ocr-single", ocrHandler)
	pprof.Register(e)
	e.Logger.Fatal(e.Start(":8081"))
}
