package binarytree

import "fmt"

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

func (binaryNode *BinaryNode) Insert(aNode *BinaryNode) {
	if binaryNode == nil {
		return
	}
	if binaryNode.data < aNode.data {
		if binaryNode.left == nil {
			binaryNode.left = aNode
		} else {
			binaryNode.left.Insert(aNode)
		}
	} else {
		if binaryNode.right == nil {
			binaryNode.right = aNode
		} else {
			binaryNode.right.Insert(aNode)
		}
	}
	fmt.Println("Added: ", aNode.data)
}

// Binary Tree
type BinaryTree struct {
	Root *BinaryNode

	// Diagnostic
	InOrderData []int
}

func (binaryTree *BinaryTree) Initialize() {
	binaryTree.Root = nil
	binaryTree.InOrderData = []int{}
}

func (binaryTree *BinaryTree) AppendToDataArray(aData int) {
	var lDataArray []int = make([]int, 0)
	lDataArray = append(lDataArray, aData)
	binaryTree.InOrderData = lDataArray
}

func (binaryTree *BinaryTree) Insert(aData int) {
	// leaf level.
	if binaryTree.Root == nil {
		binaryTree.Root = &BinaryNode{left: nil, right: nil, data: aData}
	} else {
		binaryTree.Root.Insert(&BinaryNode{left: nil, right: nil, data: aData})
	}
}

func TraverseInOrder(aBinaryTree *BinaryTree, aNode *BinaryNode) {
	if aNode == nil {
		return
	}
	TraverseInOrder(aBinaryTree, aNode.left)
	aBinaryTree.AppendToDataArray(aNode.data)
	TraverseInOrder(aBinaryTree, aNode.right)

}
