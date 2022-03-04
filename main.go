package main

import (
	b "go-practice/basics"
)

func main() {
	// fmt.Println("Hello world")
	var ll b.LinkedList
	var newNode b.LinkedListNode
	newNode.Data = 1
	ll.Insert(&newNode)
	ll.PrintList()
	ll.InsertData(2)
	ll.PrintList()
	ll.InsertData(3)
	ll.PrintList()
	ll.InsertData(4)
	ll.PrintList()

	ll.Delete(3)
	ll.PrintList()
}
