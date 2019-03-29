package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left,right *treeNode
}

// 值接受  指针接受
// 要改变数据内容要用指针，结构体数据大 也尽量用指针

func main()  {
	var tree_node treeNode
	tree_node.value = 2
	tree_node.left  = &treeNode{1,nil,nil}
	tree_node.right = &treeNode{4,nil,nil}

	fmt.Println(tree_node)

}
