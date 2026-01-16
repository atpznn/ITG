package main

import (
	"fmt"
	"sync"
	"time"

	"ITG/services"
	dime_transaction "ITG/services/dime/transaction"
)

func doSomething(text string) *dime_transaction.DimeTransactionLogResponse {
	start := time.Now()
	texts := services.SplitWithDate(text)
	if transactions, err := dime_transaction.ParseTextToTransactionReq(texts); err == nil {
		if res, err := dime_transaction.NewDimeSingleTransactions(transactions); err == nil {
			elapsed := time.Since(start)
			fmt.Printf("งานเสร็จสิ้น per func: %s\n", elapsed)
			return res
		}
	}
	return nil
}

func main() {
	// text, err := service.ParseImagePathToText("../imageTest/dime/en/transactions/1.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// pattern := `\d{1,2}\s[A-Z][a-z]{2}\s\d{4}\s-\s\d{2}:\d{2}:\d{2}\s(AM|PM)`
	// re := regexp.MustCompile(pattern)
	// parts := re.Split(text, -1)
	// timestamps := re.FindAllString(text, -1)
	// transactions := make([]string, len(timestamps))
	// for index, timestamp := range timestamps {
	// 	transactions[index] = fmt.Sprintf("%s%s", parts[index], timestamp)
	// }
	texts := []string{
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
		"Activity f Schedule >\nInvestment Cash\n| @ Thai stock @ us options\nQ Enter a stock symbol to search you.. 25\nDone\nNovember 2025\nDividend Withholding Tax SGOV -0.25 USD\nDeduct from Dime! ปริ 6 Nov 2025 - 08:02:18 PM\nDividend SGOV 1.72 USD\nDeposit to Dime! UsD 6 Nov 2025 - 08:02:17 PM\nOctober 2025\nBuy SGOV 496.75 USD\nExecuted Price 100.45 8 Oct 2025 - 06:45:17 PM\nShares 4.9452463\nSeptember 2025\nDividend GOOGL 0.25 USD\nDeposit to Dime! UsD 16 Sep 2025 - 06:06:38 AM\nDividend Withholding Tax GOOGL -0.03 USD\nDeduct from Dime! UsD 16 Sep 2025 - 06:06:04 AM\"",
	}
	start1 := time.Now()
	for _, t := range texts {
		doSomething(t)
	}
	elapsed1 := time.Since(start1)

	fmt.Printf("งานเสร็จสิ้น ใช้เวลาไป: %s\n", elapsed1)
	start1 = time.Now()
	var wg sync.WaitGroup
	for _, t := range texts {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			doSomething(t)
		}(t)
	}

	wg.Wait()
	elapsed1 = time.Since(start1)

	fmt.Printf("งานเสร็จสิ้น ใช้เวลาไป: %s\n", elapsed1)
	// for _, transaction := range transactions {
	// 	parser, err := dimeTransaction.NewDimeSingleTransactions(transaction)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	fmt.Println(parser.ToJson())
	// }
}
