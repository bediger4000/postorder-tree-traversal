package main

/*
Given the sequence of keys visited by a postorder traversal of a binary
search tree, reconstruct the tree.

For example, given the sequence 2, 4, 3, 8, 7, 5, you should construct
the following tree:

    5
   / \
  3   7
 / \   \
2   4   8

*/

import (
	"fmt"
	"os"
	"strconv"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func insert(node *TreeNode, value int) *TreeNode {

	if node == nil {
		return &TreeNode{data: value}
	}

	n := &(node.left)
	if value >= node.data {
		n = &(node.right)
	}
	*n = insert(*n, value)
	return node
}

// postorderInsert - given a possibly non-nil sub-tree, and
// the next value from a post-order traverse of the original
// binary search tree, give back a BST with the next value
// inserted. The key here is that the value inserted
// is always the root of the tree returned.
func postorderInsert(root *TreeNode, value int) *TreeNode {

	newnode := &TreeNode{data: value}

	if root == nil {
		return newnode
	}

	if root.left == nil && root.right == nil {
		if newnode.data > root.data {
			newnode.left = root
		}
		if newnode.data < root.data {
			newnode.right = root
		}
		return newnode
	}

	// find where to break off the "top" of the tree
	// to insert the new node
	var prev *TreeNode
	node := root

	if newnode.data > root.data {
		newnode.left = root
		for node != nil && newnode.data > node.data {
			prev = node
			node = node.right
		}
		if node != nil {
			prev.right = nil
			newnode.right = node
		}

	} else if newnode.data < root.data {
		newnode.right = root
		for node != nil && newnode.data < node.data {
			prev = node
			node = node.left
		}
		if node != nil {
			prev.left = nil
			newnode.left = node
		}
	}
	// value == root.data will just fall through,
	// and screw up construction of the tree.
	return newnode
}

func inorderArray(node *TreeNode, values []int) []int {
	if node.left != nil {
		values = inorderArray(node.left, values)
	}
	values = append(values, node.data)
	if node.right != nil {
		values = inorderArray(node.right, values)
	}
	return values
}

func postorderArray(node *TreeNode, values []int) []int {
	if node.left != nil {
		values = postorderArray(node.left, values)
	}
	if node.right != nil {
		values = postorderArray(node.right, values)
	}
	return append(values, node.data)
}

func inorderTraverse(node *TreeNode) {
	if node == nil {
		return
	}
	inorderTraverse(node.left)
	fmt.Printf("%d ", node.data)
	inorderTraverse(node.right)
}

func postorderTraverse(node *TreeNode) {
	if node == nil {
		return
	}
	postorderTraverse(node.left)
	postorderTraverse(node.right)
	fmt.Printf("%d ", node.data)
}

func drawTree(node *TreeNode, prefix string) {
	fmt.Printf("%s%p [label=\"%d\"];\n", prefix, node, node.data)
	if node.left != nil {
		drawTree(node.left, prefix)
		fmt.Printf("%s%p -> %s%p;\n", prefix, node, prefix, node.left)
	} else {
		fmt.Printf("%s%pL [shape=\"point\"];\n", prefix, node)
		fmt.Printf("%s%p -> %s%pL;\n", prefix, node, prefix, node)
	}
	if node.right != nil {
		drawTree(node.right, prefix)
		fmt.Printf("%s%p -> %s%p;\n", prefix, node, prefix, node.right)
	} else {
		fmt.Printf("%s%pR [shape=\"point\"];\n", prefix, node)
		fmt.Printf("%s%p -> %s%pR;\n", prefix, node, prefix, node)
	}
}

func main() {

	var inorderroot *TreeNode
	// var postorderroot *TreeNode
	var inputvalues []int

	for _, str := range os.Args[1:] {
		i, e := strconv.Atoi(str)
		if e == nil {
			inputvalues = append(inputvalues, i)
			inorderroot = insert(inorderroot, i)
			// postorderroot = postorderInsert(postorderroot, i)
		}
	}

	fmt.Printf("/* Input : %v */\n", inputvalues)

	if inorderroot == nil {
		return
	}

	var x []int
	in := inorderArray(inorderroot, x)
	fmt.Printf("/* In order: %v */\n", in)

	post := postorderArray(inorderroot, x)
	fmt.Printf("/* Post order: %v */\n", post)

	var postorderroot *TreeNode

	for _, value := range post {
		postorderroot = postorderInsert(postorderroot, value)
	}

	fmt.Printf("digraph g1 {\n")
	fmt.Printf("subgraph cluster_0 {\n\tlabel=\"input order insert\"\n")
	drawTree(inorderroot, "a")
	fmt.Printf("\n}\n")
	fmt.Printf("subgraph cluster_1 {\n\tlabel=\"postorder insert\"\n")
	drawTree(postorderroot, "b")
	fmt.Printf("\n}\n")
	fmt.Printf("\n}\n")
}
