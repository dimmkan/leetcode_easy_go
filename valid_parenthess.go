/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
 

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

package main

func isValid(s string) bool {
	var brackets []rune
	for _, c := range s {
			if (c == '(' || c == '{' || c == '[') {
					brackets = append(brackets, c)
			} else if (c == ')' && len(brackets) != 0 && brackets[len(brackets) - 1] == '(') {
					brackets = brackets[:len(brackets) - 1]
			} else if (c == '}' && len(brackets) != 0 && brackets[len(brackets) - 1] == '{') {
					brackets = brackets[:len(brackets) - 1]
			} else if (c == ']' && len(brackets) != 0 && brackets[len(brackets) - 1] == '[') {
					brackets = brackets[:len(brackets) - 1]
			} else {
					return false;
			}


	}

	return len(brackets) == 0
}