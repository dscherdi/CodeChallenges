package main

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/valid-parentheses/?envType=study-plan-v2&envId=top-interview-150
func isValid(s string) bool {
	var stack []rune

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (c == ')' && top != '(') || (c == ']' && top != '[') || (c == '}' && top != '{') {
				return false
			}
		}
	}

	return len(stack) == 0
}

// https://leetcode.com/problems/simplify-path/?envType=study-plan-v2&envId=top-interview-150
func simplifyPath(path string) string {
	var stack = []string{}

	dirs := strings.Split(path, "/")

	for _, s := range dirs {
		if s == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if s != "/" && s != ".." && s != "." && s != "" {
			stack = append(stack, s)
		}
	}
	correctPath := "/" + strings.Join(stack, "/")

	return correctPath
}

// https://leetcode.com/problems/evaluate-reverse-polish-notation/?envType=study-plan-v2&envId=top-interview-150
func evalRPN(tokens []string) int {
	var stack []int

	performOperation := func(op1 int, op2 int, operator string) int {
		switch operator {
		case "*":
			return op1 * op2
		case "/":
			return op1 / op2
		case "+":
			return op1 + op2
		case "-":
			return op1 - op2
		default:
			panic("Bad operator")
		}
	}

	for _, t := range tokens {
		switch t {
		case "*", "/", "+", "-":
			operand1 := stack[len(stack)-2]
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, performOperation(operand1, operand2, t))
		default:
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}
