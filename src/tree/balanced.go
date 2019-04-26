package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}



func isBalanced(node1 *TreeNode ) bool {


	return false

}

func main()  {

	var node1 TreeNode
	node1.Value = 1
	node1.Left = &TreeNode{1,nil,nil}
	node1.Right = &TreeNode{2,nil,nil}



	fmt.Print(isBalanced(&node1))


}
