package chapter2

import (
	linkedlist "github.com/abdullahkhan9375/GoLangCrackingTheCodingInterview/DataStructures/SinglyLinkedList"
)

type Node = linkedlist.Node
type LinkedList = linkedlist.LinkedList

func RemoveDuplicates(aLinkedList *LinkedList) {
	// A map of all nodes within a linked list.
	var lLookup map[int]int = make(map[int]int, 0)
	var lPrevNode *Node
	var lHead = aLinkedList.Head
	for lHead != nil {
		lData := lHead.Data
		if _, exists := lLookup[lData]; exists {
			lPrevNode.Next = lHead.Next
		} else {
			lLookup[lData] = 0
			lPrevNode = lHead
		}
		lHead = lHead.Next
	}
}

func KthtoLast(aLinkedList *LinkedList, aKth int) int {
	var lLength = aLinkedList.Length
	var lIndex = lLength - aKth

	var lHead = aLinkedList.Head
	for lIndex > 0 {
		lHead = lHead.Next
		lIndex--
	}
	return lHead.Data
}
