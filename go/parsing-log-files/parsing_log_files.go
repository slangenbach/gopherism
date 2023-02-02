package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	if re.MatchString(text) {
		return true
	} else {
		return false
	}

}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\<(~*|\**|=*|-*|-*)*\>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re := regexp.MustCompile(`(?i)"p(\w*\s*)*d"`) // yes, that's a hack

	for _, line := range lines {
		if re.MatchString(line) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")

}

func TagWithUserName(lines []string) []string {
	var taggedLines []string
	var newLine string
	re := regexp.MustCompile(`User{1}\s{1,}(\w+)`)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			newLine = fmt.Sprintf("[USR] %s %s", match[1], line)
		} else {
			newLine = line
		}
		taggedLines = append(taggedLines, newLine)
	}
	return taggedLines
}
