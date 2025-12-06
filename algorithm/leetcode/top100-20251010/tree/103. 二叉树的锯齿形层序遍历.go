package main

import "fmt"

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//示例 1：
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
//示例 2：
//
//输入：root = [1]
//输出：[[1]]
//示例 3：
//
//输入：root = []
//输出：[]
//
//提示：
//
//树中节点数目在范围 [0, 2000] 内
//-100 <= Node.val <= 100

func zigzagLevelOrder(root *TreeNode) [][]int {
	// 每次初始化
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		var temp []int

		// 同层级遍历
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]

			temp = append(temp, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, temp)
	}

	for i := 0; i < len(res); i++ {
		if i%2 == 1 {
			temp := res[i]
			for l, r := 0, len(temp)-1; l < r; l, r = l+1, r-1 {
				temp[l], temp[r] = temp[r], temp[l]
			}
		}
	}

	return res
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	//root = &TreeNode{
	//	Val:  1,
	//	Left: &TreeNode{Val: 2},
	//}
	//
	fmt.Println(zigzagLevelOrder(root))
}
