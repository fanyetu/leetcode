package main

import "fmt"

func main() {
	//s := "()[]{}"
	//s := "([)]"
	s := "{[]}"
	flag := isValid(s)
	fmt.Print(flag)
}

func isValid(s string) bool {
	var (
		stack []string = make([]string, 0)
	)

	for i := 0; i < len(s); i++ {
		if i == 0 {
			stack = append(stack, s[0:1])
			continue
		}

		if len(stack) < 1 {
			stack = append(stack, s[i:i+1])
			continue
		}

		last := stack[len(stack)-1]
		if match(last, s[i:i+1]) {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, s[i:i+1])
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}

func match(left string, right string) bool {
	switch left {
	case "(":
		if right == ")" {
			return true
		}
	case "{":
		if right == "}" {
			return true
		}
	case "[":
		if right == "]" {
			return true
		}
	default:
		return false
	}
	return false
}
