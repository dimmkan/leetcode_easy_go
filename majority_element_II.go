/*
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Example 1:

Input: nums = [3,2,3]
Output: [3]
Example 2:

Input: nums = [1]
Output: [1]
Example 3:

Input: nums = [1,2]
Output: [1,2]
 

Constraints:

1 <= nums.length <= 5 * 104
-109 <= nums[i] <= 109
 

Follow up: Could you solve the problem in linear time and in O(1) space?
*/

func majorityElement(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	e1, e2, c1, c2 := 0, 1, 0, 0
	for _, e := range a {
		switch {
		case e == e1:
			c1++
		case e == e2:
			c2++
		case c1 == 0:
			e1 = e
			c1 = 1
		case c2 == 0:
			e2 = e
			c2 = 1
		default:
			c1--
			c2--
		}
	}

	res := []int{}

	if maj(a, e2) {
		res = append(res, e2)
	}

	if maj(a, e1) {
		res = append(res, e1)
	}

	return res
}

func maj(a []int, n int) bool {
	c := 0
	for _, e := range a {
		if e == n {
			c++
		}
	}
	return c > (len(a) / 3)
}