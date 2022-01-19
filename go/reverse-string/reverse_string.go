package reverse

import "fmt"

func Reverse(input string) string {
	output := ""
    for _, char := range input {
        output = fmt.Sprintf("%s%s", string(char), output)
    }
	return output
}
