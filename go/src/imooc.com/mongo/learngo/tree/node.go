package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node TreeNode) print() {
	fmt.Println(node.Value)
}

func (node *TreeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored!")
		return
	}
	node.Value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.print()
	node.Left.Traverse()
	node.Right.Traverse()
}

func createNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func main() {
	var root TreeNode
	fmt.Println(root)

	root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)
	root.Right.Right = createNode(2)

	root.Right.Right.setValue(3)
	root.Right.Right.print()

	/*nodes := []TreeNode {
		{value: 3},
		{},
		{6, nil, nil},
	}*/
	root.Traverse()
	var pRoot *TreeNode
	pRoot.setValue(200)
}
