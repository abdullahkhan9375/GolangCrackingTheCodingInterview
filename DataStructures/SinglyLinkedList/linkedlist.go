package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Next *Node
	Data int
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (linkedList *LinkedList) Initialize() {
	linkedList.Head = nil
	linkedList.Length = 0
}

func (linkedList *LinkedList) validateIndex(aIndex int) error {
	if aIndex >= linkedList.Length {
		return errors.New("index can't be greater than length")
	} else if aIndex < 0 {
		return errors.New("index can't be lower than 0")
	}
	return nil
}

func (linkedList *LinkedList) Prepend(aData int) {
	var lNode *Node
	if linkedList.Head == nil {
		lNode = &Node{Next: nil, Data: aData}
	} else {
		lNode = &Node{Next: linkedList.Head.Next, Data: aData}
	}

	linkedList.Head = lNode
	linkedList.Length++
}

func (linkedList *LinkedList) Insert(aData int) {
	var lNode = &Node{Next: nil, Data: aData}
	if linkedList.Head == nil {
		linkedList.Head = lNode
	} else {
		var lHead = linkedList.Head
		for lHead.Next != nil {
			lHead = lHead.Next
		}
		lHead.Next = lNode
	}
	linkedList.Length++
}

func (linkedList *LinkedList) Get(aIndex int) (int, error) {
	var lHead *Node = linkedList.Head
	err := linkedList.validateIndex(aIndex)
	if err != nil {
		return -1, err
	}

	if aIndex == 0 {
		return lHead.Data, nil
	}

	var lIndex int = 0
	for lIndex != aIndex {
		lHead = lHead.Next
		lIndex++
	}
	return lHead.Data, nil
}

func (linkedList *LinkedList) PopAt(aIndex int) (*Node, error) {
	var lHead *Node = linkedList.Head
	var lPrevHead *Node
	var lPopNode *Node
	err := linkedList.validateIndex(aIndex)
	if err != nil {
		return &Node{}, err
	}
	if aIndex == 0 {
		lPopNode = linkedList.Head
		linkedList.Head = linkedList.Head.Next
	} else {
		lIndex := 0
		lHead = linkedList.Head
		for lIndex != aIndex {
			lPrevHead = lHead
			lHead = lHead.Next
			lIndex++
		}
		lPopNode = lHead
		lPrevHead.Next = lHead.Next
	}
	linkedList.Length--
	return lPopNode, nil
}

func (linkedList *LinkedList) Find(aData int) (int, error) {
	var lIndex int = 0
	if linkedList.Head == nil {
		return -1, errors.New("linked list is empty")
	}
	if linkedList.Head.Data == aData {
		return lIndex, nil
	}
	var lHead *Node = linkedList.Head
	for lHead.Data != aData && lHead.Next != nil {
		lHead = lHead.Next
		lIndex++
	}
	if lHead.Data == aData {
		return lIndex, nil
	}
	return -1, errors.New("could not find the node")
}

func (linkedList *LinkedList) Delete(aData int) error {
	var lElementIndex, err = linkedList.Find(aData)
	if err != nil {
		return err
	}
	linkedList.PopAt(lElementIndex)
	linkedList.Length--
	return nil
}

func (linkedList *LinkedList) GetDataArray() []int {
	var lHead *Node = linkedList.Head
	var lDataArray []int = make([]int, 0)
	for lHead != nil {
		lDataArray = append(lDataArray, lHead.Data)
		lHead = lHead.Next
	}
	return lDataArray
}

func (linkedList *LinkedList) PrintListData() {
	toPrint := linkedList.Head

	for linkedList.Length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		linkedList.Length--
	}
	fmt.Printf("\n")
}
