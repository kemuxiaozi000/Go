package tree

import "fmt"

// Node struct定义
type Node struct {
	Value       int
	Left, Right *Node
}

// Print 为结构体创建方法 （接收者）
func (node Node) Print() {
	fmt.Println(node.Value)
}

// SetValue 只有使用指针才可以改变结构内容
func (node *Node) SetValue(value int) {
	if node == nil {
		return
	}
	node.Value = value
}

// CreateTreeNode haha
func CreateTreeNode(value int) *Node {
	// 注意： 返回局部变量地址 （c++在栈上/堆上）
	return &Node{Value: value}
}

// Traverse 中序
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
