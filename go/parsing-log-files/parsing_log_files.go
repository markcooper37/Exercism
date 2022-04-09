package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\])|(\[DBG\])|(\[INF\])|(\[WRN\])|(\[ERR\])|(\[FTL\])`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	return []string{}
}

func CountQuotedPasswords(lines []string) int {
	return 0
}

func RemoveEndOfLineText(text string) string {
	return ""
}

func TagWithUserName(lines []string) []string {
	return []string{}
}
