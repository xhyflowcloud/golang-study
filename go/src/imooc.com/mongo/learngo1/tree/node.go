package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored!")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func main() {
	var root Node
	fmt.Println(root)

	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Right.Right = CreateNode(2)

	root.Right.Right.SetValue(3)
	root.Right.Right.Print()

	/*nodes := []Node {
		{value: 3},
		{},
		{6, nil, nil},
	}*/
	root.Traverse()

}
