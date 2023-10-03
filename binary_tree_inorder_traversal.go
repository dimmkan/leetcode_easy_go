/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.



Example 1:


Input: root = [1,null,2,3]
Output: [1,3,2]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	for root != nil {
		if root.Left == nil {
			ans = append(ans, root.Val)
			root = root.Right
		} else {
			prev := root.Left
			for prev.Right != nil && prev.Right != root {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = root
				root = root.Left
			} else {
				ans = append(ans, root.Val)
				prev.Right = nil
				root = root.Right
			}
		}
	}
	return ans
}
