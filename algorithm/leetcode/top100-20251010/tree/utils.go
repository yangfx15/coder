package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Node1 = [1,2,3,null,5,null,4]
var Node1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Right: &TreeNode{Val: 5},
	},
	Right: &TreeNode{
		Val:   3,
		Right: &TreeNode{Val: 4},
	},
}

// Node2 = [1,2,3,4,null,null,null,5]
var Node2 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}},
	},
	Right: &TreeNode{
		Val: 3,
	},
}

// Node3 = [1,2,3]
var Node3 = &TreeNode{
	Val: 2,
	Left: &TreeNode{
		Val: 1,
	},
	Right: &TreeNode{
		Val: 3,
	},
}

// Node4 = [1,5,3,4,6]
var Node4 = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 1,
	},
	Right: &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 6},
	},
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTree(root.Left)
	fmt.Printf("%v,", root.Val)
	printTree(root.Right)
}
