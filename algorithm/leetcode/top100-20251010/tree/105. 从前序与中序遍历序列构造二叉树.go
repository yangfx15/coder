package main

//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，
//请构造二叉树并返回其根节点。
//
//示例 1:
//
//
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//示例 2:
//
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder)
}

func build(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	tree := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	preorder = preorder[1:]

	if len(preorder) > 0 {
		//fmt.Printf("left:pre:%v, in:%v\n", preorder[:i], inorder[:i])
		//fmt.Printf("right:pre:%v, in:%v\n", preorder[i:], inorder[i+1:])
		tree.Left = build(preorder[:i], inorder[:i])
		tree.Right = build(preorder[i:], inorder[i+1:])
	}

	return tree
}

func main() {
	//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	//输出: [3,9,20,null,null,15,7]
	printTree(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	printTree(buildTree([]int{-1}, []int{-1}))
}
