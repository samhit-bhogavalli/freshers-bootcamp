package main

import "fmt"
//struct for tree node
type treeNode struct {
	value string
	left *treeNode
	right *treeNode
}
//function for preorder traversal
func preorder(root *treeNode){
	fmt.Println(root.value)
	if root.left != nil {
		preorder(root.left)
	}
	if root.right != nil {
		preorder(root.right)
	}
}
//function for postorder traversal
func postorder(root *treeNode){
	if root.left != nil {
		postorder(root.left)
	}
	if root.right != nil {
		postorder(root.right)
	}
	fmt.Println(root.value)
}
func main(){
	//testing
	tree := treeNode{"+", &treeNode{"a",nil,nil},&treeNode{"-",&treeNode{"b",nil,nil},&treeNode{"c",nil,nil}} }
	fmt.Println("preorder")
	preorder(&tree)
	fmt.Println("postorder")
	postorder(&tree)
}