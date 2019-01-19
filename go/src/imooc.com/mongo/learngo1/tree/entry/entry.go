package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"imooc.com/mongo/learngo1/tree"
)

type MyTreeNode struct {
	node *tree.Node
}

func (myTreeNode *MyTreeNode) postOrder() {
	if myTreeNode == nil || myTreeNode.node == nil {
		return
	}
	left := MyTreeNode{myTreeNode.node.Left}
	right := MyTreeNode{myTreeNode.node.Right}
	left.postOrder()
	right.postOrder()
	myTreeNode.node.Print()
}

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(10000))
}
func main() {
	var root tree.Node
	//fmt.Println(root)

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateNode(2)

	root.Right.Right.SetValue(3)
	//root.Right.Right.Print()

	/*nodes := []Node {
		{value: 3},
		{},
		{6, nil, nil},
	}*/
	root.Traverse()
	/*var myTreeNode MyTreeNode
	myTreeNode.node = &root*/
	//myTreeNode.postOrder()

	//testSparse()
	/*nodeCount := 0
	 root.TraverseFunc(func(node *tree.Node) {
		 nodeCount++
	 })

	 fmt.Println("Node count:", nodeCount)*/

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		fmt.Println(node)
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value:", maxNode)

}
