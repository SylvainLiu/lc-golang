package robber

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res := robTree(root)
	return int(math.Max(res[0], res[1]))
}

func robTree(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0, 0}
	}

	left := robTree(root.Left)
	right := robTree(root.Right)

	robCurrent := float64(root.Val) + left[0] + right[0]
	notRobCurrent := math.Max(left[0], left[1]) + math.Max(right[0], right[1])

	return []float64{notRobCurrent, robCurrent}
}
