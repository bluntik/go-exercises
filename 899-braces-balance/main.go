package main

import (
	"fmt"
)

func main() {
	tests := []struct {
		payload  []byte
		expected bool
	}{
		{[]byte(`([{}])`), true},
		{[]byte(`(){}(())[]`), true},
		{[]byte(`(`), false},
		{[]byte(`((((`), false},
		{[]byte(`([)]`), false},
		{[]byte(`]`), false},
	}

	for i, test := range tests {
		if validate(test.payload) != test.expected {
			panic(fmt.Sprintf("bad test case validate %d. Payload: `%s`", i, string(test.payload)))
		}

		if validateType(test.payload) != test.expected {
			panic(fmt.Sprintf("bad test case validateType %d. Payload: `%s`", i, string(test.payload)))
		}
	}

	fmt.Printf("DONE")
}

type Stack []rune

func (s *Stack) push(r rune) {
	*s = append(*s, r)
}

func (s Stack) isEmpty() bool {
	return len(s) == 0
}

func (s Stack) top() rune {
	return s[len(s)-1]
}

func (s *Stack) pop() {
	n := len(*s) - 1

	*s = (*s)[:n]
}

func validateType(payload []byte) bool {
	stack := Stack{}

	for _, sym := range payload {
		s := rune(sym)

		if s == '(' || s == '{' || s == '[' {
			stack.push(s)
			continue
		}

		if stack.isEmpty() {
			return false
		}

		if (s == ')' && stack.top() != '(') || (s == ']' && stack.top() != '[') || (s == '}' && stack.top() != '{') {
			return false
		}

		stack.pop()
	}

	return stack.isEmpty()
}

func validate(payload []byte) bool {
	var stack []rune

	for _, sym := range payload {
		s := rune(sym)

		if s == '(' || s == '{' || s == '[' {
			stack = append(stack, s)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		n := len(stack) - 1
		top := stack[n]

		if (s == ')' && top != '(') || (s == ']' && top != '[') || (s == '}' && top != '{') {
			return false
		}

		stack = stack[:n]
	}

	return len(stack) == 0
}
