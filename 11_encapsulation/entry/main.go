package main

import (
	"fmt"

	tree ".."
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{
		myNode.node.Left,
	}
	left.postOrder()
	right := myTreeNode{
		myNode.node.Right,
	}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)
	root.Right.Left.SetValue(4)

	// nodes := []tree.Node{
	// 	{Value: 3},
	// 	{},
	// }
	// fmt.Println(nodes)

	// 树结构
	// 	   3
	//   /   \
	//  0     5
	//   \   /
	//    2 4

	// 中序 02345
	// 前序
	// 后序

	// pRoot := &root
	// pRoot.print()
	// pRoot.setValue(200)
	// pRoot.print()
	fmt.Println("中序遍历")
	root.Traverse()
	fmt.Println("前序遍历")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

}
