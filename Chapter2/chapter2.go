package chapter2

import (
	"errors"
	"fmt"
	"math"

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

// Implement an algorithm to delete a node in the middle (i.e., any node but
// the first and last node, not necessarily the exact middle) of a singly linked list, given only access to
// that node.

// 1 Node. -> return.
// 2 Nodes. -> return.
// 3 Nodes. -> A -> B -> C ===> A -> C
// K Nodes. -> A -> ... -> K -> ... -> Z ===> A -> ... -> Z

func DeleteMiddleNode(aLinkedList *LinkedList) error {
	var lLength int = aLinkedList.Length
	if lLength >= 0 && lLength <= 2 {
		return errors.New("linked list must have > 2 nodes")
	}
	fmt.Println("Length is: ", lLength)
	var lMiddleNode = aLinkedList.Head
	var lMiddlePoint int = int(math.Floor(float64(lLength) / 2))
	var lIndex int = 0
	for lIndex < lMiddlePoint {

		lMiddleNode = lMiddleNode.Next
		lIndex++
	}
	fmt.Println(lMiddleNode)
	lRepNode := lMiddleNode.Next
	lMiddleNode.Data = lRepNode.Data
	lMiddleNode.Next = lRepNode.Next
	return nil
}

// Partition: Write code to partition a linked list around a value x, such that all nodes less than x come
// before all nodes greater than or equal to x. If x is contained within the list, the values of x only need
// to be after the elements less than x (see below). The partition element x can appear anywhere in the
// "right partition"; it does not need to appear between the left and right partitions.

// 1 Node -> return;
// 2 Nodes -> switch nodes.
// 3 Nodes -> Pivot according to the selected node.
// k Nodes
func Partition(aLinkedList *LinkedList, aData int) error {
	var lLength int = aLinkedList.Length
	if lLength == 0 || lLength == 1 {
		return errors.New("linked list should have more than 1 node")
	}

	lLLCopy := *aLinkedList

	lHead := lLLCopy.Head
	for lHead.Next != nil {
		if lHead.Data > aData {
			lBuffer := lHead
			lHead = nil
			aLinkedList.Insert(lBuffer.Data)
		} else if lHead.Data < aData {
			lBuffer := lHead
			lHead = nil
			aLinkedList.Prepend(lBuffer.Data)
		}
		lHead = lHead.Next
		fmt.Println(lHead)
	}
	return nil
}
