package main

import (
	tree ".."
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)

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
	root.Traverse()

}
