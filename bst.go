package main

import "fmt"

// TreeNode represents a node in the Binary Search Tree (BST).
// SRP: The TreeNode struct has a single responsibility - representing a node in the tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTreeOperations defines an interface for basic tree operations.
// ISP: Separate interfaces for only essential operations.
type BinaryTreeOperations interface {
	Insert(value int)
	Search(value int) bool
	Delete(value int)
}

// TraversalOperations defines an interface for tree traversal.
// ISP: This interface is specifically for traversal-related methods.
type TraversalOperations interface {
	InOrderTraversal() []int
}

// BinarySearchTree implements BinaryTreeOperations and TraversalOperations.
// SRP: BinarySearchTree is responsible for managing the BST structure and operations.
type BinarySearchTree struct {
	Root *TreeNode
}

// Insert adds a new node with the given value to the BST.
// SRP: Responsible only for insertion logic.
// OCP: New insertion strategies can be added by implementing BinaryTreeOperations in other types.
func (bst *BinarySearchTree) Insert(value int) {
	bst.Root = insertNode(bst.Root, value)
}

// insertNode is a helper function that recursively finds the correct location for the new node.
func insertNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = insertNode(node.Right, value)
	}
	return node
}

// Search checks if a value exists in the BST.
func (bst *BinarySearchTree) Search(value int) bool {
	return searchNode(bst.Root, value)
}

// searchNode is a helper function that recursively searches for the value in the tree.
func searchNode(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return searchNode(node.Left, value)
	} else {
		return searchNode(node.Right, value)
	}
}

// Delete removes a node with the given value from the BST.
func (bst *BinarySearchTree) Delete(value int) {
	bst.Root = deleteNode(bst.Root, value)
}

// deleteNode is a helper function that deletes a node and maintains the BST properties.
func deleteNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}
	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minNode := findMin(node.Right)
		node.Value = minNode.Value
		node.Right = deleteNode(node.Right, minNode.Value)
	}
	return node
}

// findMin finds the minimum value node in a subtree.
func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// InOrderTraversal performs an in-order traversal of the BST.
// SRP: This method handles only in-order traversal logic.
// DIP: This method can be used on any type implementing TraversalOperations.
func (bst *BinarySearchTree) InOrderTraversal() []int {
	var result []int
	inOrder(bst.Root, &result)
	return result
}

// inOrder is a helper function for in-order traversal.
func inOrder(node *TreeNode, result *[]int) {
	if node != nil {
		inOrder(node.Left, result)
		*result = append(*result, node.Value)
		inOrder(node.Right, result)
	}
}

func main() {
	// DIP: Main depends on abstractions, not concrete implementations.
	var tree BinaryTreeOperations = &BinarySearchTree{}
	var traversal TraversalOperations = &BinarySearchTree{}

	// Insert nodes
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	// Search nodes
	fmt.Println("Search 30:", tree.Search(30)) // Output: true
	fmt.Println("Search 90:", tree.Search(90)) // Output: false

	// Delete node
	tree.Delete(20)
	fmt.Println("In-order traversal after deleting 20:", traversal.InOrderTraversal())

	// In-order traversal
	fmt.Println("In-order traversal:", traversal.InOrderTraversal())
}
