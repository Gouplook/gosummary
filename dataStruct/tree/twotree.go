package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func Insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Val: val}
	}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
	}
	if val > root.Val {
		root.Right = Insert(root.Right, val)
	}
	return root
}

func Search(root *Node, val int) bool {
	if root == nil {
		return false
	}
	if root.Val == val {
		return true
	}
	if val < root.Val {
		return Search(root.Left, val)
	}
	if val > root.Val {
		return Search(root.Right, val)
	}
	return false
}

func Traverse(root *Node) {
	if root == nil {
		return
	}
	Traverse(root.Left)
	fmt.Printf("%v", root.Val)
	Traverse(root.Right)
}
func PreorderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v", root.Val)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
	return
}
func InorderTraversal(root *Node) {
	if root == nil {
		return
	}
	InorderTraversal(root.Left)
	fmt.Printf("%v", root.Val)
	InorderTraversal(root.Right)
	return
}
func PostorderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostorderTraversal(root.Left)
	PostorderTraversal(root.Right)
	fmt.Printf("%v", root.Val)
}

func Delete(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if val < root.Val {
		root.Left = Delete(root.Left, val)
	}
	if val > root.Val {
		root.Right = Delete(root.Right, val)
	}
	if val == root.Val {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
	}
	return root
}

func main() {
	var root *Node
	root = Insert(root, 5)
	root = Insert(root, 1)

	root = Insert(root, 3)
	root = Insert(root, 7)
	root = Insert(root, 2)
	root = Insert(root, 4)
	fmt.Println("Traversal of the tree:")
	Traverse(root)

}
