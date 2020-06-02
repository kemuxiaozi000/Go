package main

import "fmt"

// Node ..
type Node struct {
	value       int
	left, right *Node
}

func (node *Node) setValue(value int) {
	if node == nil {
		return
	}
	node.value = value
}

func createNode(value int) *Node {
	// 注意： 返回局部变量地址 （c++在栈上/堆上）
	return &Node{value: value}
}

// func (node *node) f() {
// 	fmt.Println(node.value)
// 	return
// }

func (node *Node) traverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.left.traverseFunc(f)
	f(node)
	node.right.traverseFunc(f)
}

func (node *Node) traverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.traverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}

func main() {
	var root Node
	root = Node{value: 3}
	root.left = &Node{}
	root.right = &Node{5, nil, nil}
	root.right.left = new(Node)
	root.left.right = createNode(2)
	root.right.left.setValue(4)
	c := root.traverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.value > maxNode {
			maxNode = node.value
		}
	}
	fmt.Println("Max node value:", maxNode)
}

// 树结构
// 	   3
//   /   \
//  0     5
//   \   /
//    2 4
