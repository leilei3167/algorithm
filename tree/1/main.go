package main

import "fmt"

/* 链表实现二叉树 */
type TreeNode struct {
	Data  string    //存放数据
	Left  *TreeNode //左儿子
	Right *TreeNode //右儿子
}

//前序遍历，先访问根节点，再访问左儿子，右儿子
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	//先打印根节点
	fmt.Print(tree.Data, "")
	//再打印左儿子
	PreOrder(tree.Left)
	//再打印右儿子
	PreOrder(tree.Right)

}

//中序遍历，先访问左儿子，再根节点，最后右节点
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	MidOrder(tree.Left)

	//再打印根
	fmt.Print(tree.Data, "")

	//再打印右
	MidOrder(tree.Right)

}

//后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {

		return
	}
	PostOrder(tree.Left)
	PostOrder(tree.Right)
	fmt.Print(tree.Data, "")

}

//test
func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "b"}
	t.Right = &TreeNode{Data: "c"}
	t.Left.Left = &TreeNode{Data: "d"}
	t.Left.Right = &TreeNode{Data: "e"}
	t.Right.Left = &TreeNode{Data: "f"}

	//前序
	fmt.Println("前序遍历")
	PreOrder(t)

	fmt.Println("\n中序遍历")
	MidOrder(t)

	fmt.Println("\n后序遍历")
	PostOrder(t)

}
