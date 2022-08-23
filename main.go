package main

import (
	"fmt"

	linkedlist "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/DataStructures/SinglyLinkedList"
)

func main() {
	list := linkedlist.LinkedList{}
	list.Initialize()

	list.Insert(2)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Insert(10)
	list.Insert(3)
	fmt.Println(list.GetDataArray())
}
