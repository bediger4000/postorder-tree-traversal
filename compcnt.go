package main

/*
 * Make a Binary Search Tree with random number values,
 * get the post-order traverse, then recreate the tree
 * based on post-order traverse. Count comparisons needed
 * to make the new tree.
 */

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

 ./postorder 5 3 7 2 4 8

 should give us the tree above

*/

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
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

func find(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if node.data == value {
		return true
	}
	if node.left != nil && find(node.left, value) {
		return true
	}
	if node.right != nil && find(node.right, value) {
		return true
	}
	return false
}

// postorderInsert - given a possibly non-nil sub-tree, and
// the next value from a post-order traverse of the original
// binary search tree, give back a BST with the next value
// inserted. The key here is that the value inserted
// is always the root of the tree returned.
func postorderInsert(root *TreeNode, value int) (*TreeNode, int) {

	comparisonCount := 0

	newnode := &TreeNode{data: value}

	comparisonCount++
	if root == nil {
		return newnode, comparisonCount
	}

	if root.left == nil && root.right == nil {
		comparisonCount++
		if newnode.data > root.data {
			newnode.left = root
		}
		comparisonCount++
		if newnode.data < root.data {
			newnode.right = root
		}
		return newnode, comparisonCount
	}

	// find where to break off the "top" of the tree
	// to insert the new node
	var prev *TreeNode
	node := root

	comparisonCount++
	if newnode.data > root.data {
		newnode.left = root
		comparisonCount++
		for node != nil && newnode.data > node.data {
			comparisonCount++
			prev = node
			node = node.right
		}
		if node != nil {
			prev.right = nil
			newnode.right = node
		}

	} else if newnode.data < root.data {
		comparisonCount++
		newnode.right = root
		for node != nil && newnode.data < node.data {
			comparisonCount++
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
	return newnode, comparisonCount
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

	var treenodecnt int

	if val, e := strconv.Atoi(os.Args[1]); e != nil {
		log.Fatalf("Problem getting the number of tree nodes: %v\n", e)
	} else {
		treenodecnt = val
	}

	rand.Seed(time.Now().UnixNano())

	var inorderroot *TreeNode
	var inputvalues []int

	for i := 0; i < treenodecnt; i++ {
		value := rand.Intn(60000)
		if !find(inorderroot, value) {
			inputvalues = append(inputvalues, value)
			inorderroot = insert(inorderroot, value)
		}
	}

	if inorderroot == nil {
		return
	}

	var x []int

	post := postorderArray(inorderroot, x)

	var postorderroot *TreeNode

	var totalComparisons, count int
	for _, value := range post {
		postorderroot, count = postorderInsert(postorderroot, value)
		totalComparisons += count
	}

	fmt.Printf("%d\t%d\n", treenodecnt, totalComparisons)
}
