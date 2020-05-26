package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 为结构体创建方法 （接收者）
func (node treeNode) print() {
	fmt.Println(node.value)
	return
}

// 只有使用指针才可以改变结构内容
func (node *treeNode) setValue(value int) {
	if node == nil {
		return
	}
	node.value = value
}

// 遍历
// 中序
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 工厂函数
func createTreeNode(value int) *treeNode {
	// 注意： 返回局部变量地址 （c++在栈上/堆上）
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createTreeNode(2)

	// nodes := []treeNode{
	// 	{value: 3},
	// 	{},
	// 	{6, nil, &root},
	// }
	// fmt.Println(nodes)

	// 结构体方法
	root.right.left.setValue(4)

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
	root.traverse()

}
