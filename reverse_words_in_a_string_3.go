/*
Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.



Example 1:

Input: s = "Let's take LeetCode contest"
Output: "s'teL ekat edoCteeL tsetnoc"
Example 2:

Input: s = "God Ding"
Output: "doG gniD"


Constraints:

1 <= s.length <= 5 * 104
s contains printable ASCII characters.
s does not contain any leading or trailing spaces.
There is at least one word in s.
All the words in s are separated by a single space.
*/

package main

func reverseWords(s string) string {
	t := []byte(s)
	for i := 0; i < len(t); i++ {
		j := i
		for j < len(t) && t[j] != ' ' {
			j++
		}
		for st, ed := i, j-1; st < ed; st, ed = st+1, ed-1 {
			t[st], t[ed] = t[ed], t[st]
		}
		i = j
	}
	return string(t)
}
