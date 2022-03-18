package main

import (
	"fmt"
	b "go-practice/basics"
)

func main() {
	// fmt.Println("Hello world")
	// var ll b.LinkedList
	// var newNode b.LinkedListNode
	// newNode.Data = 1
	// ll.Insert(&newNode)
	// ll.PrintList()
	// ll.InsertData(2)
	// ll.PrintList()
	// ll.InsertData(3)
	// ll.PrintList()
	// ll.InsertData(4)
	// ll.PrintList()

	// ll.Delete(3)
	// ll.PrintList()

	var bt b.BinaryTree
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(1)
	bt.Insert(4)
	bt.Insert(2)
	bt.Insert(9)
	bt.Insert(8)

	bt.InOrderTraversal()
	bt.PostOrderTraversal()
	fmt.Println(bt.Maximum(bt.Root).Data)
	fmt.Println(bt.Minimum(bt.Root).Data)
	fmt.Println(bt.Successor(bt.Root).Data)
	fmt.Println(bt.Predecessor(bt.Root.Left.Left))
}
