package markdown

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	paragraphs := strings.Split(markdown, "\n")
	output := ""
	list := false
	for index, paragraph := range paragraphs {
		paragraph = strings.Replace(paragraph, "__", "<strong>", 1)
		paragraph = strings.Replace(paragraph, "__", "</strong>", 1)
		paragraph = strings.Replace(paragraph, "_", "<em>", 1)
		paragraph = strings.Replace(paragraph, "_", "</em>", 1)
		header := 0
		char := paragraph[0]
		if char == '#' {
			pos := 0
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			output += fmt.Sprintf("<h%d>%s</h%d>", header, paragraph[pos+1:], header)
		} else if char == '*' {
			newParagraph := "<li>" + paragraph[2:] + "</li>"
			if !list {
				newParagraph = "<ul>" + newParagraph
			}
			list = true
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
