package main

import (
	"fmt"
	"math"
)

type NodeTree struct {
	Value int
	Left,Right *NodeTree
}

//二叉树深度
func maxDepth(node *NodeTree) int {

	if node == nil {
		return 0
	}
	left := maxDepth(node.Left) + 1
	right := maxDepth(node.Right) + 1

	return int(math.Max(float64(left),float64(right)))

}

//对称二叉树
func isSymmetric(node1 *NodeTree , node2 *NodeTree)  bool{

	if node1 == nil && node2 == nil {
		return true
	}
	if node1.Value == node2.Value {
		return isSymmetric(node1.Left,node2.Right) && isSymmetric(node1.Right,node2.Left)
	}
	return false
	
}

//高度平衡的二叉树
func isBalanced()  bool {

	return false

}

func main()  {

	var node1,node2 NodeTree
	node1.Value = 1
	node1.Left = &NodeTree{1,nil,nil}
	node1.Right = &NodeTree{2,&NodeTree{3,nil,nil},&NodeTree{4,nil,nil}}

	fmt.Print(maxDepth(&node1))

	node2.Value = 1
	node2.Left = &NodeTree{1,nil,nil}
	node2.Right = &NodeTree{2,&NodeTree{3,nil,nil},&NodeTree{4,nil,nil}}

	fmt.Print(isSymmetric(&node1,&node2))


}

