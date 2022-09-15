package binarytree_test

import (
	"fmt"
	"math"
	"testing"

	. "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/Common"
	binarytree "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/DataStructures/BinaryTree"
)

type BinaryNode = binarytree.BinaryNode
type BinaryTree = binarytree.BinaryTree

func TestCase1_Init(aTesting *testing.T) {
	// Arrange
	var lBinaryTree = &BinaryTree{}

	// Act
	lBinaryTree.Initialize()

	// Assert
	got := lBinaryTree
	if got == nil {
		aTesting.Error("Expected: Binary tree to not equal nil \n Got: ", got)
	}
}

func TestCase2_Insert(aTesting *testing.T) {
	// Arrange
	var lBinaryTree = &BinaryTree{}
	lBinaryTree.Initialize()

	// Act
	lBinaryTree.Insert(5)

	// Assert
	lDataArray := []int{5}
	binarytree.TraverseInOrder(lBinaryTree, lBinaryTree.Root)
	if !IsEqual(lBinaryTree.InOrderData, lDataArray) {
		aTesting.Error("Expected: ", lDataArray, "\n Got: ", lBinaryTree.InOrderData)
	}
}

func TestCase3_Insert(aTesting *testing.T) {
	// Arrange
	var lBinaryTree = &BinaryTree{}
	lBinaryTree.Initialize()

	lMax := int(math.Max(float64(4), float64(6)))
	fmt.Println(lMax)
	// Act
	lBinaryTree.Insert(5)
	lBinaryTree.Insert(2)
	lBinaryTree.Insert(1)

	// Assert
	lDataArray := []int{5, 2, 1}
	binarytree.TraverseInOrder(lBinaryTree, lBinaryTree.Root)
	if !IsEqual(lBinaryTree.InOrderData, lDataArray) {
		aTesting.Error("Expected: ", lDataArray, "\n Got: ", lBinaryTree.InOrderData)
	}
}
