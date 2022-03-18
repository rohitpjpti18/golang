package basics

import "fmt"

type BinaryTreeNode struct {
	Data   int
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
	Parent *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func (b *BinaryTree) InOrder(node *BinaryTreeNode) {
	if node != nil {
		b.InOrder(node.Left)
		fmt.Printf("%d ", node.Data)
		b.InOrder(node.Right)
	}
}

func (b *BinaryTree) InOrderTraversal() {
	b.InOrder(b.Root)
	fmt.Println()
}

func (b *BinaryTree) PostOrder(node *BinaryTreeNode) {
	if node != nil {
		b.PostOrder(node.Left)
		b.PostOrder(node.Right)
		fmt.Printf("%d ", node.Data)
	}
}

func (b *BinaryTree) PostOrderTraversal() {
	if b.Root == nil {
		fmt.Printf("null")
	}

	b.PostOrder(b.Root)
	fmt.Println()
}

func (b *BinaryTree) Minimum(node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return nil
	}
	ptr := node

	for ptr.Left != nil {
		ptr = ptr.Left
	}
	return ptr
}

func (b *BinaryTree) Maximum(node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return nil
	}
	ptr := node

	for ptr.Right != nil {
		ptr = ptr.Right
	}

	return ptr
}

func (b *BinaryTree) Successor(node *BinaryTreeNode) *BinaryTreeNode {
	if node.Right != nil {
		return b.Minimum(node.Right)
	}

	parent := node.Parent
	ptr := node

	for parent != nil {
		if parent.Left == ptr {
			break
		}
		ptr = parent
		parent = ptr.Parent
	}

	return parent
}

func (b *BinaryTree) Predecessor(node *BinaryTreeNode) *BinaryTreeNode {
	if node.Left != nil {
		return b.Maximum(node.Left)
	}

	parent := node.Parent
	ptr := node

	for parent != nil {
		if parent.Right == ptr {
			break
		}
		ptr = parent
		parent = ptr.Parent
	}

	return parent
}

func (b *BinaryTree) Insert(val int) {
	var newNode BinaryTreeNode
	newNode.Data = val

	if b.Root == nil {
		b.Root = &newNode
		return
	}
	ptr := b.Root

	for ptr != nil {
		if ptr.Data == val {
			return
		} else if ptr.Data > val {
			if ptr.Left == nil {
				newNode.Parent = ptr
				ptr.Left = &newNode
				return
			} else {
				ptr = ptr.Left
			}
		} else {
			if ptr.Right == nil {
				newNode.Parent = ptr
				ptr.Right = &newNode
				return
			} else {
				ptr = ptr.Right
			}
		}
	}
}
