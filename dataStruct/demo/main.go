package main

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 插入节点
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
		return root
	}
	if val > root.Val {
		root.Right = insert(root.Right, val)
		return root
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return getMin(root.Left)
}

// 在二叉树中删除节点
func delete(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		root.Left = delete(root.Left, val)
		return root
	}
	if val > root.Val {
		root.Right = delete(root.Right, val)
		return root
	}
	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	minNode := getMin(root.Right)
	root.Val = minNode.Val
	root.Right = delete(root.Right, minNode.Val)
	return root
}

// 先序遍历
func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	println(root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
	return
}

// 中序遍历
func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	println(root.Val)
	inOrder(root.Right)
	return
}

// 后序遍历
func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	println(root.Val)
	return
}

// 搜索
func search(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if val < root.Val {
		return search(root.Left, val)
	}
	return search(root.Right, val)
	return nil
}

//func main() {
//	root := &TreeNode{Val: 5, Left: nil, Right: nil}
//	r2 := insert(root, 3)
//	r3 := insert(r2, 4)
//	r4 := insert(r3, 6)
//	r5 := insert(r4, 7)
//	preOrder(r5)
//
//}
