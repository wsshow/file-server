package utils

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	if e := stack.list.Back(); e != nil {
		stack.list.Remove(e)
		return e.Value
	}

	return nil
}

func (stack *Stack) Contain(value interface{}) bool {
	for e := stack.list.Front(); e != nil; e = e.Next() {
		if e == value {
			return true
		}
	}
	return false
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.Len() == 0
}
