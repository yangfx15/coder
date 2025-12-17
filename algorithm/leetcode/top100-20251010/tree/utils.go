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

// Node5 = [3,5,1,6,2,0,8,null,null,7,4]
var Node5 = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 6},
		Right: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 4},
		},
	},
	Right: &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 8},
	},
}

// Node6 = [5,4,8,11,null,13,4,7,2,null,null,5,1]
var Node6 = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{Val: 11,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 2},
		},
	},
	Right: &TreeNode{Val: 8,
		Left: &TreeNode{Val: 13},
		Right: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 1},
		},
	},
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTree(root.Left)
	fmt.Printf("%v, ", root.Val)
	printTree(root.Right)
}
