package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	paragraphs := strings.Split(markdown, "\n")
	output := ""
	list := false
	for index, paragraph := range paragraphs {
		re := regexp.MustCompile(`__(.*)__`)
		paragraph = re.ReplaceAllString(paragraph, "<strong>$1</strong>")
		re = regexp.MustCompile(`_(.*)_`)
		paragraph = re.ReplaceAllString(paragraph, "<em>$1</em>")
		char := paragraph[0]
		if char == '#' {
			header, pos := 0, 0
			for char == '#' {
				header++
				pos++
				char = paragraph[pos]
			}
			output += fmt.Sprintf("<h%d>%s</h%d>", header, paragraph[pos+1:], header)
		} else if char == '*' {
			newParagraph := "<li>" + paragraph[2:] + "</li>"
			if !list {
				newParagraph = "<ul>" + newParagraph
				list = true
			}
			if index == len(paragraphs)-1 || paragraphs[index+1][0] != '*' {
				newParagraph += "</ul>"
				list = false
			}
			output += newParagraph
		} else {
			output += "<p>" + paragraph + "</p>"
		}
	}
	return output
}
