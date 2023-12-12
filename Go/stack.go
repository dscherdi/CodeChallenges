package main

import "strings"

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
