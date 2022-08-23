package linkedlist_test

import (
	"testing"

	. "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/Common"
	linkedlist "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/DataStructures/SinglyLinkedList"
)

type LinkedList = linkedlist.LinkedList

func TestCase1_SLL_Insert(aTest *testing.T) {

	// Arrange
	var list LinkedList = LinkedList{}
	list.Initialize()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	var lDataArray = []int{1, 2, 3, 4, 5, 6}

	// Act
	got := list.GetDataArray()

	// Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expectd: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase2_SLL_Insert(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	list.Insert(1)
	var lDataArray = []int{1}

	//Act
	got := list.GetDataArray()

	//Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expectd: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase3_SLL_Insert(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	var lDataArray = []int{}

	//Act
	got := list.GetDataArray()

	//Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expectd: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase4_SLL_Prepend(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	list.Prepend(2)
	list.Insert(5)
	var lDataArray = []int{2, 5}

	//Act
	got := list.GetDataArray()

	//Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expectd: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase5_SLL_Get(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	list.Prepend(2)
	list.Insert(5)
	var lDataArray = []int{2, 5}

	//Act
	got := list.GetDataArray()

	//Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expectd: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase6_SLL_Get(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	list.Insert(1)
	list.Insert(5)
	list.Insert(10)
	list.Insert(11)
	list.Insert(15)

	var lElement int = 10

	//Act
	got, err := list.Get(2)

	//Assert
	if err != nil {
		aTest.Error("Expectd: ", lElement, "\n Recieved: ", got)
	}
}

func TestCase7_SLL_Get(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	list.Insert(1)
	list.Insert(5)
	list.Insert(10)
	list.Insert(11)
	list.Insert(15)

	var lElement int = -1

	//Act
	got, err := list.Get(-10)

	//Assert
	if err == nil {
		aTest.Error("Expectd: ", lElement, "\n Recieved: ", got)
	}
}

func TestCase8_SLL_Get(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	list.Insert(1)
	list.Insert(5)
	list.Insert(10)
	list.Insert(11)
	list.Insert(15)

	var lElement int = -1

	//Act
	got, err := list.Get(6)

	//Assert
	if err == nil {
		aTest.Error("Expectd: ", lElement, "\n Recieved: ", got)
	}
}

func TestCase9_SLL_Get(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	list.Insert(1)
	list.Insert(5)
	list.Insert(10)
	list.Insert(11)
	list.Insert(15)

	var lElement int = 15

	//Act
	got, err := list.Get(4)

	//Assert
	if err != nil {
		aTest.Error("Expectd: ", lElement, "\n Recieved: ", got)
	}
}

func TestCase10_SLL_Delete(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	list.Insert(1)
	list.Insert(5)
	list.Insert(10)
	list.Insert(11)
	list.Insert(15)

	var lDataArray = []int{1, 10, 11, 15}

	//Act
	list.Delete(5)
	got := list.GetDataArray()

	//Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expectd: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase11_SLL_Delete(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()
	list.Insert(1)

	var lDataArray = []int{}

	//Act
	list.Delete(1)
	got := list.GetDataArray()

	//Assert
	if !IsEqual(got, lDataArray) {
		aTest.Error("Expectd: ", lDataArray, "\n Recieved: ", got)
	}
}

func TestCase12_SLL_Delete(aTest *testing.T) {
	var list LinkedList = LinkedList{}

	//Arrange
	list.Initialize()

	//Act
	err := list.Delete(10)

	//Assert
	if err == nil {
		aTest.Error("Expected nothing")
	}
}
