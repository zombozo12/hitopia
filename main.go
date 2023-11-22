package main

import "fmt"

func main() {
	// WeightedStrings
	sWeightedString := "abbcccd"
	queries := []int{1, 3, 9, 8}
	fmt.Println(WeightedStrings(sWeightedString, queries)) // Output: ['Yes', 'Yes', 'Yes', 'No']

	// HighestPalindrome
	sPalindrome := "3943"
	k := 1
	fmt.Println(HighestPalindrome(sPalindrome, k)) // Output: '3993'

	// Brackets
	sBrackets := "{ [ ( ] ) }"
	fmt.Println(IsBalanced(sBrackets)) // Output: 'NO'
}

func WeightedStrings(s string, queries []int) []string {
	results := make([]string, len(queries))
	for i := 0; i < len(s); i++ {
		weight := 0
		for j := i; j < len(s); j++ {
			weight += int(s[j]) - 96
			if j-i+1 == len(s) || s[j] != s[j+1] {
				if weight == queries[i] {
					results[i] = "Yes"
				} else {
					results[i] = "No"
				}
			}
		}
	}
	return results
}

func HighestPalindrome(s string, k int) string {
	if !isPalindrome(s) {
		return "-1"
	}
	max := s
	for i := 0; i < len(s) && k > 0; i++ {
		if s[i] < '9' {
			s = s[:i] + "9" + s[i+1:]
			k--
			if isPalindrome(s) && s > max {
				max = s
			}
			s = s[:i] + "0" + s[i+1:]
		}
	}
	return max
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func IsBalanced(s string) string {
	stack := make([]rune, 0)
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 {
				return "NO"
			}
			top := stack[len(stack)-1]
			if (c == ')' && top != '(') || (c == ']' && top != '[') || (c == '}' && top != '{') {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return "NO"
	}
	return "YES"
}
