package main

import "fmt"

//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
//叶子节点 是指没有子节点的节点。
//
//示例 1：
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
//示例 2：
//
//输入：root = [1,2,3], targetSum = 5
//输出：[]
//示例 3：
//
//输入：root = [1,2], targetSum = 0
//输出：[]
func pathSum(root *TreeNode, targetSum int) [][]int {
	var dfs func(cur *TreeNode, target int, list []int)
	var res [][]int
	dfs = func(cur *TreeNode, target int, list []int) {
		if cur == nil {
			return
		}

		list = append(list, cur.Val)
		//fmt.Printf("cur:%v, target:%v, list:%v \n", cur.Val, target, list)
		if cur.Left == nil && cur.Right == nil {
			if cur.Val == target {
				temp := make([]int, len(list))
				copy(temp, list)
				res = append(res, temp)
			}
			return
		}

		dfs(cur.Left, target - cur.Val, list)
		dfs(cur.Right, target - cur.Val, list)
	}
	dfs(root, targetSum, []int{})
	return res
}

func main()  {
	fmt.Println(pathSum(Node6, 22))
}