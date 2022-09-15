package stack

import (
	"errors"
)

// TODO: Extract to common.
type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	top *Node
}

type SetofStacks struct {
	stacks    []Stack
	threshold int
}

// Stack base methods.
func (stack *Stack) Initialize() {
	stack.top = nil
}

func (stack *Stack) Push(aData int) {
	var lNode *Node = &Node{Next: nil, Data: aData}
	if stack.top == nil {
		stack.top = lNode
	} else {
		var lTemp *Node = stack.top
		stack.top = lNode
		stack.top.Next = lTemp
	}
}

func (stack *Stack) Pop() (int, error) {
	var lTop = stack.top
	if lTop == nil {
		return -1, errors.New("stack is empty")
	}
	stack.top = stack.top.Next
	return lTop.Data, nil
}

func (stack *Stack) Peek() (int, error) {
	if stack.top == nil {
		return -1, errors.New("stack is empty")
	} else {
		return stack.top.Data, nil
	}
}

func (stack *Stack) GetDataArray() ([]int, error) {
	var lDataArray []int = make([]int, 0)
	if stack.top == nil {
		return lDataArray, errors.New("stack is empty")
	} else {
		var lStackTop = stack.top
		for lStackTop != nil {
			lDataArray = append(lDataArray, lStackTop.Data)
			lStackTop = lStackTop.Next
		}
	}
	return lDataArray, nil
}

// Set of Stacks.
func (stackSet *SetofStacks) Initialize(aThreshold int) {
	stackSet.stacks = []Stack{}
	stackSet.threshold = aThreshold
}

func (stackSet *SetofStacks) addStack(aStack *Stack) {
	var lStacksSlice = make([]Stack, 0)
	copy(lStacksSlice, stackSet.stacks)
	lStacksSlice = append(lStacksSlice, *aStack)
	stackSet.stacks = lStacksSlice
}

// TODO: test this.
func (stackSet *SetofStacks) Push(aData int) {

	if len(stackSet.stacks) == 0 {
		var lStack = Stack{}
		lStack.Initialize()
		stackSet.addStack(&lStack)
	}

	var lStack = stackSet.stacks[len(stackSet.stacks)-1]

	var lStackLength = 0

	var lStackTop = lStack.top

	for lStackTop != nil {
		lStackLength++
		lStackTop = lStackTop.Next
	}

	if lStackLength < stackSet.threshold {
		lStack.Push(aData)
	} else {
		var lNewStack = Stack{}
		lNewStack.Initialize()
		lNewStack.Push(aData)
		stackSet.addStack(&lNewStack)
	}
}

func (stackSet *SetofStacks) GetStackArray() ([][]int, error) {
	if len(stackSet.stacks) == 0 {
		return [][]int{}, errors.New("stack set is empty")
	}

	var lTwoDArray = make([][]int, 0)
	for _, lStack := range stackSet.stacks {
		lStackArray, err := lStack.GetDataArray()
		if err == nil {
			lTwoDArray = append(lTwoDArray, lStackArray)
		} else {
			lTwoDArray = append(lTwoDArray, []int{})
		}
	}

	return lTwoDArray, nil
}
