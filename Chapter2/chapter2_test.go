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
