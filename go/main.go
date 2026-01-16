package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"runtime"
	"sync"

	"ITG/services"
	dime_transaction "ITG/services/dime/transaction"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/otiai10/gosseract"
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
	transactions *dime_transaction.DimeTransactionRequest
	processing   func(text *dime_transaction.DimeTransactionRequest) (*dime_transaction.DimeTransactionLogResponse, error)
	sort         bool
}

// func sortResult(result []any) {
// 	sort.Slice(result, func(i, j int) bool {
// 		getDate := func(item any) time.Time {
// 			switch v := item.(type) {
// 			case *dime_transaction_dividend.DimeDividendTransaction:
// 				return v.ExecutedDate
// 			case *dime_transaction_fee.DimeTransactionFee:
// 				return v.ExecutedDate
// 			case *dime_transaction_stock.DimeTransactionStock:
// 				return v.ExecutedDate
// 			default:
// 				return time.Time{}
// 			}
// 		}
// 		return getDate(result[i]).Before(getDate(result[j]))
// 	})
// }

func singleProcess(param Param) error {
	// results := make([]any, len(param.transactions))
	result, err := param.processing(param.transactions)
	if err != nil {
		return param.c.String(http.StatusInternalServerError, err.Error())
	}
	// if param.sort {
	// 	sortResult(results)
	// }
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
	processing func(transactions *dime_transaction.DimeTransactionRequest) (*dime_transaction.DimeTransactionLogResponse, error),
) func(c echo.Context) error {
	return func(c echo.Context) error {
		u := new(DimeBody)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		texts := services.SplitWithDate(u.Text)
		transactions, err := dime_transaction.ParseTextToTransactionReq(texts)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		param := Param{transactions: transactions, processing: processing, c: c, sort: u.Sort}
		return singleProcess(param)
		// return req(param)
	}
}

var (
	ocrClient *gosseract.Client
	mutex     sync.Mutex
)

func ocrHandler(c echo.Context) error {
	err := c.Request().ParseMultipartForm(32 << 20)
	if err != nil {
		return c.String(http.StatusBadRequest, "Parse form failed: "+err.Error())
	}
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusBadRequest, "Parse form failed")
	}

	files := form.File["images"]
	if len(files) == 0 {
		return c.String(http.StatusBadRequest, "No images uploaded")
	}
	results := make([]string, len(files))
	for index, fileHeader := range files {
		// client := gosseract.NewClient()
		// client.SetLanguage("eng")
		src, err := fileHeader.Open()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		tmpFile, err := os.CreateTemp("", "ocr-*.png")
		if err != nil {
			src.Close()
			return c.String(http.StatusInternalServerError, "Cannot create temp file")
		}

		if _, err := io.Copy(tmpFile, src); err != nil {
			tmpFile.Close()
			src.Close()
			return err
		}
		ocrClient.SetImage(tmpFile.Name())
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		src.Close()
		text, err := ocrClient.Text()
		// client.Close()
		gosseract.ClearPersistentCache()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		results[index] = text
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"results": results,
	})
}

func ocrHandlerSafe(c echo.Context) error {
	mutex.Lock()
	defer mutex.Unlock()
	// err := c.Request().ParseMultipartForm(32 << 20)
	// if err != nil {
	// 	return c.String(http.StatusBadRequest, "Parse form failed: "+err.Error())
	// }
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println(err.Error())
		return c.String(http.StatusBadRequest, "Parse form failed")
	}

	files := form.File["images"]
	if len(files) == 0 {
		return c.String(http.StatusBadRequest, "No images uploaded")
	}
	results := make([]string, len(files))
	for index, fileHeader := range files {
		// client := gosseract.NewClient()
		// client.SetLanguage("eng")
		src, err := fileHeader.Open()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		tmpFile, err := os.CreateTemp("", "ocr-*.png")
		if err != nil {
			src.Close()
			return c.String(http.StatusInternalServerError, "Cannot create temp file")
		}

		if _, err := io.Copy(tmpFile, src); err != nil {
			tmpFile.Close()
			src.Close()
			return err
		}
		ocrClient.SetImage(tmpFile.Name())
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		src.Close()
		text, err := ocrClient.Text()
		// client.Close()
		gosseract.ClearPersistentCache()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		results[index] = text
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"results": results,
	})
}

func ocrHandlerMulti(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.String(http.StatusBadRequest, "Parse form failed")
	}

	files := form.File["images"]
	if len(files) == 0 {
		return c.String(http.StatusBadRequest, "No images uploaded")
	}

	numFiles := len(files)
	jobs := make(chan job, numFiles)
	resultsChan := make(chan result, numFiles)
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, resultsChan)
		}()
	}

	for i, f := range files {
		jobs <- job{index: i, file: f}
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	finalResults := make([]string, numFiles)
	for res := range resultsChan {
		finalResults[res.index] = res.text
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"workers_used": numWorkers,
		"total_files":  numFiles,
		"results":      finalResults,
	})
}

type job struct {
	index int
	file  *multipart.FileHeader
}

type result struct {
	index int
	text  string
}

func worker(jobs <-chan job, results chan<- result) {
	// สร้าง Client หนึ่งตัวต่อหนึ่ง Worker เพื่อประหยัด Resource
	client := gosseract.NewClient()
	client.SetLanguage("eng")
	defer client.Close()

	for j := range jobs {
		f, _ := j.file.Open()
		imgBytes, _ := io.ReadAll(f)
		f.Close()

		client.SetImageFromBytes(imgBytes)
		text, _ := client.Text()

		results <- result{index: j.index, text: text}
	}
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
	e.POST("single/dime/text-process", compose(dime_transaction.NewDimeSingleTransactions))
	e.POST("ocr-single", ocrHandler)
	e.POST("ocr-single-safe", ocrHandlerSafe)
	e.POST("ocr-multi", ocrHandlerMulti)
	// e.POST("multi/dime/text-process", compose(dime_transaction.NewDimeMultipleTransactions))
	e.Logger.Fatal(e.Start(":8081"))
}
