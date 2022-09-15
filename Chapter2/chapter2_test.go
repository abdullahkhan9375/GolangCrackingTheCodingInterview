package chapter2_test

import (
	"testing"

	chapter2 "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/Chapter2"
	. "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/Common"
	linkedlist "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/DataStructures/SinglyLinkedList"
)

type LinkedList = linkedlist.LinkedList

func TestCase1_RemoveDuplicates(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(4)
	list.Insert(6)
	var lDataArray = []int{1, 2, 3, 4, 6}

	// Act
	chapter2.RemoveDuplicates(&list)
	got := list.GetDataArray()
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase2_RemoveDuplicates(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)
	list.Insert(2)
	list.Insert(2)
	list.Insert(2)
	list.Insert(2)
	list.Insert(6)
	var lDataArray = []int{1, 2, 6}

	// Act
	chapter2.RemoveDuplicates(&list)
	got := list.GetDataArray()
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase3_RemoveDuplicates(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(3)
	list.Insert(4)
	list.Insert(4)
	var lDataArray = []int{1, 2, 3, 4}

	// Act
	chapter2.RemoveDuplicates(&list)
	got := list.GetDataArray()
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase4_RemoveDuplicates(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(10)
	list.Insert(4)
	list.Insert(4)
	list.Insert(10)
	var lDataArray = []int{1, 2, 3, 10, 4}

	// Act
	chapter2.RemoveDuplicates(&list)
	got := list.GetDataArray()
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase5_KthTolast(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(10)
	list.Insert(4)
	list.Insert(4)
	list.Insert(10)

	// K is 3. So third last element.
	var KthElement = 4

	// Act
	got := chapter2.KthtoLast(&list, 3)
	if KthElement != got {
		aTest.Error("Expected: ", KthElement, "\n Recieved: ", got)
	}
}

func TestCase6_DeleteMiddleNode(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)
	list.Insert(2)

	// Act
	got := chapter2.DeleteMiddleNode(&list)

	// Assert
	if got == nil {
		aTest.Error("Expected: error", "\n Recieved: ", got)
	}
}

func TestCase7_DeleteMiddleNode(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()

	// Act
	got := chapter2.DeleteMiddleNode(&list)

	// Assert
	if got == nil {
		aTest.Error("Expected: error", "\n Recieved: ", got)
	}
}

func TestCase8_DeleteMiddleNode(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)
	list.Insert(2)
	list.Insert(5)
	list.Insert(6)

	lDataArray := []int{1, 2, 6}
	// Act
	chapter2.DeleteMiddleNode(&list)
	got := list.GetDataArray()

	// Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase9_DeleteMiddleNode(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)
	list.Insert(2)
	list.Insert(5)
	list.Insert(6)
	list.Insert(5)
	list.Insert(9)
	list.Insert(10)

	lDataArray := []int{1, 2, 5, 5, 9, 10}
	// Act
	chapter2.DeleteMiddleNode(&list)
	got := list.GetDataArray()

	// Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase10_Partition(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)

	// Act
	got := chapter2.DeleteMiddleNode(&list)

	// Assert
	if got == nil {
		aTest.Error("Expected: 1 node error")
	}
}

func TestCase11_Partition(aTest *testing.T) {
	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(11)
	list.Insert(10)
	list.Insert(5)
	list.Insert(5)
	list.Insert(2)
	list.Insert(1)

	lDataArray := []int{2, 1, 5, 5, 11, 10}
	// Act
	chapter2.Partition(&list, 5)
	got := list.GetDataArray()

	// Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expected: ", lDataArray, "\n Recieved: ", got)
	}
}
