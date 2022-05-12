package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		qLen := len(queue)
		level := []int{}
		for idx := 0; idx < qLen; idx++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				level = append(level, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		if len(level) != 0 {
			result = append(result, level)
		}
	}
	return result
}
