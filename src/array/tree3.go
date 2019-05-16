package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var last int

func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if isValidBST(root.Left) {
		if root.Val > last {
			last = root.Val
			isValidBST(root.Right)
		}
	}

	return false

}


func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums , 0 ,len(nums) -1 )
}

//有序数组转平衡二叉树
func buildTree(nums[] int , l int , r int) *TreeNode{
	if l > r {
		return nil
	}
	m := l + (r - l+1) / 2;
	return &TreeNode{nums[m],buildTree(nums, l, m - 1),buildTree(nums, m + 1, r)}

}

//给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.
func averageOfLevels(root *TreeNode) []float64 {
	var nums[]float64
	var node[]TreeNode
	node = append(node,*root)
	for len(node) > 0 {
		size := len(node)
		cur_num := len(node)
		var sum float64
		for size >0 {
			size--
			cur := node[0]
			node = node[1:]
			sum += float64(cur.Val)
			if cur.Left != nil {
				node = append(node,*cur.Left)
			}
			if cur.Right != nil {
				node = append(node,*cur.Right)
			}
		}
		nums = append(nums,sum / float64(cur_num))
	}
	return nums

}



func levelOrderBottom(root *TreeNode) [][]int {



}

func main()  {

	//nums := []int{-10,-5,0,3,9}
	//fmt.Printf("%p",sortedArrayToBST(nums))
	//var node1 *TreeNode
	//node1.Val = 1
	//node1.Left = &TreeNode{1,nil,nil}
	//node1.Right = &TreeNode{2,&TreeNode{3,nil,nil},&TreeNode{4,nil,nil}}
	//
	//node1 := &TreeNode{1,&TreeNode{1,nil,nil},&TreeNode{2,&TreeNode{3,nil,nil},&TreeNode{4,nil,nil}}}

	//fmt.Print(averageOfLevels(node1))


	fmt.Print(isValidBST(&TreeNode{0,nil,nil}))


}
