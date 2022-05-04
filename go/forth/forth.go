package forth

import (
	"errors"
	"strconv"
	"strings"
)

type Stack []int

func Forth(input []string) ([]int, error) {
	stack := Stack{}
	userDefined := map[string][]string{}
	for _, line := range input {
		splitOps := strings.Split(strings.ToLower(line), " ")
		if splitOps[0] == ":" && splitOps[len(splitOps)-1] == ";" {
			if _, err := strconv.Atoi(splitOps[1]); err == nil {
				return []int{}, errors.New("can't redefine number")
			}
			userDefined[splitOps[1]] = GetBasicOps(splitOps[2 : len(splitOps)-1], userDefined)
		} else {
			splitOps = GetBasicOps(splitOps, userDefined)
			for _, op := range splitOps {
				if number, err := strconv.Atoi(op); err == nil {
					stack = append(stack, number)
				} else {
					if err := stack.PerformOperation(op); err != nil {
						return []int{}, err
					}
				}
			}
		}
	}
	return stack, nil
}

func GetBasicOps(splitOps []string, userDefined map[string][]string) []string {
	for index, op := range splitOps {
		if userDefinedOperation, ok := userDefined[op]; ok {
			splitOps = append(splitOps[:index], append(userDefinedOperation, splitOps[index+1:]...)...)
		}
	}
	return splitOps
}

func (s *Stack) PerformOperation(op string) error {
	if len(*s) == 0 {
		return errors.New("cannot perform operation: no elements in stack")
	} else if len(*s) == 1 && (op == "+" || op == "-" || op == "*" || op == "/" || op == "swap" || op == "over") {
		return errors.New("not enough elements in stack to perform operation: " + op)
	}
	switch op {
	case "+":
		*s = append((*s)[:len(*s)-2], (*s)[len(*s)-2]+(*s)[len(*s)-1])
	case "-":
		*s = append((*s)[:len(*s)-2], (*s)[len(*s)-2]-(*s)[len(*s)-1])
	case "*":
		*s = append((*s)[:len(*s)-2], (*s)[len(*s)-2]*(*s)[len(*s)-1])
	case "/":
		if (*s)[1] == 0 {
			return errors.New("can't divide by zero")
		}
		*s = append((*s)[:len(*s)-2], (*s)[len(*s)-2]/(*s)[len(*s)-1])
	case "dup":
		*s = append(*s, (*s)[len(*s)-1])
	case "drop":
		*s = (*s)[:len(*s)-1]
	case "swap":
		(*s)[len(*s)-2], (*s)[len(*s)-1] = (*s)[len(*s)-1], (*s)[len(*s)-2]
	case "over":
		*s = append(*s, (*s)[len(*s)-2])
	default:
		return errors.New("invalid operation")
	}
	return nil
}
