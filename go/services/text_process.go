package services

import (
	"fmt"
	"regexp"
)

var pattern = `\d{1,2}\s[A-Z][a-z]{2}\s\d{4}\s-\s\d{2}:\d{2}:\d{2}\s(AM|PM)`
func SplitWithDate(text string) []string{
	re := regexp.MustCompile(pattern)
	parts := re.Split(text, -1)
	timestamps := re.FindAllString(text, -1)
	transactions := make([]string, len(timestamps))
	for index, timestamp := range timestamps {
		transactions[index] = fmt.Sprintf("%s%s", parts[index], timestamp)
	}
	return transactions
}