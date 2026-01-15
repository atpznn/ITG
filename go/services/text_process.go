package services

import (
	"regexp"
)

var (
	pattern = `\d{1,2}\s[A-Z][a-z]{2}\s\d{4}\s-\s\d{2}:\d{2}:\d{2}\s(AM|PM)`
	re      = regexp.MustCompile(pattern)
)

func SplitWithDate(text string) []string {
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
