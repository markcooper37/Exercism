package parsinglogfiles

import (
    "regexp"
    "strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
    re := regexp.MustCompile(`\<[~*=-]*\>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    re := regexp.MustCompile(`(?i)".*password.*"`)
    count := 0
    for _, line := range lines {
        if re.MatchString(line) {
            count++
        }
    }
	return count
}

func RemoveEndOfLineText(text string) string {
    re := regexp.MustCompile(`end-of-line[0-9]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User +\S+`)
    for i, line := range lines {
        user := re.FindString(line)
        if user != "" {
            username := strings.Fields(user)[1]
            lines[i] = "[USR] " + username + " " + lines[i]
        }
    }
	return lines
}
