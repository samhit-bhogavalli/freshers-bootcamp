package main

import "fmt"

//struct for tree node
type treeNode struct {
	value string
	left *treeNode
	right *treeNode
}

//function for preorder traversal
func (root treeNode)preorder() {
	fmt.Println(root.value)
	if root.left != nil {
		root.left.preorder()
	}
	if root.right != nil {
		root.right.preorder()
	}
}

//function for postorder traversal
func (root treeNode)postorder(){
	if root.left != nil {
		root.left.postorder()
	}
	if root.right != nil {
		root.right.postorder()
	}
	fmt.Println(root.value)
}

func main(){
	//testing
	tree := treeNode{"+", &treeNode{"a",nil,nil},&treeNode{"-",&treeNode{"b",nil,nil},&treeNode{"c",nil,nil}} }
	fmt.Println("preorder")
	tree.preorder()
	fmt.Println("postorder")
	tree.postorder()
}