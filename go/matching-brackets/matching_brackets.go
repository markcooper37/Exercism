package brackets

func Bracket(input string) bool {
	brackets := []rune{}
	bracketsMap := map[rune]rune{'}': '{', ']': '[', ')': '('}
	for _, char := range input {
		if char == '{' || char == '[' || char == '(' {
			brackets = append(brackets, char)
		} else if char == '}' || char == ']' || char == ')' {
			if len(brackets) == 0 {
				return false
			} else if bracketsMap[char] != brackets[len(brackets) - 1] {
				return false
			}
			brackets = brackets[:len(brackets)-1]
		}
	}
	return len(brackets) == 0
}
