package main

import "fmt"

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
//有效 二叉搜索树定义如下：
//
//节点的左子树只包含 严格小于 当前节点的数。
//节点的右子树只包含 严格大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。

var rootList []int

func isValidBST(root *TreeNode) bool {
	rootList = []int{}
	leftOrder(root)
	for i := 0; i < len(rootList)-1; i++ {
		if rootList[i] >= rootList[i+1] {
			return false
		}
	}

	return true
}

func leftOrder(root *TreeNode) {
	if root == nil {
		return
	}
	leftOrder(root.Left)
	rootList = append(rootList, root.Val)
	leftOrder(root.Right)
}

func main() {
	fmt.Println(isValidBST(Node3))
	fmt.Println(isValidBST(Node4))
}
