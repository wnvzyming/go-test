package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func isSameTree(node1 *Node , node2 *Node) bool {

	if node1 == nil || node2 == nil {
		if node1 == nil && node2 == nil {
			return true
		}
		return false
	}

	if node1.Value == node2.Value {
		return isSameTree(node1.Left,node2.Left) && isSameTree(node2.Right,node2.Right)
	}

	return false

}

func main()  {

	var node1,node2 Node
	node1.Value = 1
	node1.Left = &Node{1,nil,nil}
	node1.Right = &Node{2,nil,nil}

	node2.Value = 1
	node2.Left = &Node{1,nil,nil}
	node2.Right = &Node{2,nil,nil}

	fmt.Print(isSameTree(&node1,&node2))


}
