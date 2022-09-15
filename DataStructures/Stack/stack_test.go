package stack_test

import (
	"fmt"
	"testing"

	. "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/Common"
	stack "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/DataStructures/Stack"
)

type Stack = stack.Stack
type SetofStacks = stack.SetofStacks
type Node = stack.Node

func TestCase1_Push(aTest *testing.T) {
	// Arrange
	var lStack Stack = Stack{}
	lStack.Initialize()
	lStack.Push(4)
	var lDataArray = []int{4}

	// Act
	got, _ := lStack.GetDataArray()

	//Assert
	if !IsEqual(lDataArray, got) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase2_Push(aTest *testing.T) {
	// Arrange
	var lStack Stack = Stack{}
	lStack.Initialize()
	lStack.Push(2)
	lStack.Push(3)
	lStack.Push(9)
	lStack.Push(5)
	lStack.Push(10)

	var lDataArray = []int{10, 5, 9, 3, 2}

	// Act
	got, _ := lStack.GetDataArray()

	//Assert
	if !IsEqual(lDataArray, got) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase3_Push(aTest *testing.T) {
	// Arrange
	var lStack Stack = Stack{}
	lStack.Initialize()

	var lDataArray = []int{}

	// Act
	got, _ := lStack.GetDataArray()

	//Assert
	if !IsEqual(lDataArray, got) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase4_Pop(aTest *testing.T) {
	// Arrange
	var lStack Stack = Stack{}
	lStack.Initialize()
	lStack.Push(2)
	lStack.Push(3)
	lStack.Push(9)
	lStack.Push(5)
	lStack.Push(10)

	var lPopData = 10

	// Act
	got, _ := lStack.Pop()
	if lPopData != got {
		aTest.Error("Expected: ", lPopData, "\n Recieved: ", got)
	}
}

func TestCase5_Pop(aTest *testing.T) {
	// Arrange
	var lStack Stack = Stack{}
	lStack.Initialize()

	// Act
	_, err := lStack.Pop()
	if err == nil {
		aTest.Error("Expected nothing")
	}
}

func TestCase6_SetofStacks(aTest *testing.T) {
	// Arrange
	var lStackSet SetofStacks = SetofStacks{}
	lStackSet.Initialize(3)
	lStackSet.Push(5)
	lStackSet.Push(6)
	lStackSet.Push(10)
	lStackSet.Push(11)
	lStackSet.Push(12)

	lDataArray := [][]int{{5, 6, 10}, {11, 12}}
	// Act
	lStackArray, err := lStackSet.GetStackArray()

	if err == nil {
		fmt.Println("Data array: ", lDataArray)
		fmt.Println("Stack array: ", lStackArray)
	}
}
