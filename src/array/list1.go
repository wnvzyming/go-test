package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next

}

func reverseList(node *ListNode)  *ListNode{

	if node == nil {
		return node;
	}

	//var node1 ListNode
	node = reverseList(node.Next)


	node.Next.Next = node
	node.Next = nil

	return node

}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val{
		res = l2
		res.Next = mergeTwoLists(l1,l2.Next)
	}else{
		res = l1
		res.Next = mergeTwoLists(l1.Next,l2)
	}
	return res

}



func main() {
	node := new(ListNode)
	node.Val = 1
	node.Next = new(ListNode)
	node.Next.Val = 2
	node.Next.Next = new(ListNode)
	node.Next.Next.Val = 3
	node.Next.Next.Next = new(ListNode)
	node.Next.Next.Next.Val = 4


	node2 := new(ListNode)
	node2.Val = 2
	node2.Next = new(ListNode)
	node2.Next.Val = 3
	node2.Next.Next = new(ListNode)
	node2.Next.Next.Val = 5
	node2.Next.Next.Next = new(ListNode)
	node2.Next.Next.Next.Val = 7


	res := mergeTwoLists(node,node2)


	fmt.Printf("%p",res)


	//cur := node
	//for cur != nil {
	//	fmt.Println(cur.Val)
	//	cur = cur.Next
	//}
	////deleteNode(node.Next.Next)
	//reverseList(node)
	//
	//fmt.Println("+++++++++++++++++++")
	//cur = node
	//for cur != nil{
	//	fmt.Println(cur.Val)
	//	cur = cur.Next
	//}
}
