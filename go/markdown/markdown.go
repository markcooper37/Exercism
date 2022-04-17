package markdown

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	header := 0
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	list := false
	html := ""
	for pos := 0; pos < len(markdown); pos++ {
		char := markdown[pos]
		if char == '#' {
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			html += fmt.Sprintf("<h%d>", header)
		} else if char == '*' {
			if !list {
				html += "<ul>"
				list = true
			}
			html += "<li>"
			pos++
		} else if char == '\n' {
			if list {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		} else {
			html += string(char)
		}
	}
	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	} else if list {
		return html + "</li></ul>"
	} else {
		return "<p>" + html + "</p>"
	}
}
