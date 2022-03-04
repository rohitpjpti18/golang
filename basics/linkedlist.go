package basics

import "fmt"

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
}

func (l *LinkedList) Insert(newNode *LinkedListNode) bool {
	if l.Head == nil {
		l.Head = newNode
		return true
	}

	var ptr *LinkedListNode = l.Head
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = newNode

	return true
}

func (l *LinkedList) InsertData(Data int) bool {
	var newNode LinkedListNode
	newNode.Data = Data
	newNode.Next = nil

	if l.Head == nil {
		l.Head = &newNode
		return true
	}

	var ptr *LinkedListNode = l.Head
	for ptr.Next != nil {
		ptr = ptr.Next
	}
	ptr.Next = &newNode

	return true
}

func (l *LinkedList) Delete(Data int) bool {
	var tempNode LinkedListNode
	tempNode.Data = -1
	tempNode.Next = l.Head
	l.Head = &tempNode

	ptr := l.Head
	for ptr.Next != nil {
		if ptr.Next.Data == Data {
			ptr.Next = ptr.Next.Next
			l.Head = l.Head.Next
			return true
		} else {
			ptr = ptr.Next
		}
	}
	l.Head = l.Head.Next

	return false
}

func (l *LinkedList) PrintList() {
	var ptr *LinkedListNode = l.Head

	for ptr != nil {
		fmt.Printf("%d->", ptr.Data)
		ptr = ptr.Next
	}
	fmt.Println("NULL")
}
