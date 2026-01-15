package service

import (
	"testing"
)

func TestShouldbeReturnCorrectText(t *testing.T) {
	imgBytes, err := GetImageBytes("../../../imageTest/1000000973.jpg")
	if err != nil {
		return
	}
	result, err := ParseImageToText(imgBytes)
	if err != nil {
		return
	}
	if result != "Sell JPM = NYSE\n0.2465087 Shares\n\nExecuted Price 199.486 USD\n\nTotal Credit\n\nStock Amount 49.18 USD\nCommission Fee -0.07 USD\nVAT 7% -0.0052 USD\nReserved Fee ©\n\nSEC Fee -0.01 USD\nTAF Fee -0.01 USD\nSubmission Date 28 Mar 2024 - 22:13\nCompletion date 28 Mar 2024 - 22:13\nOrder Type 3? Market Order" {
		return
	}
}
