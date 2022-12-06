package lists

import (
	"lists/lists/array"
	"lists/lists/linkedlist"
)

type List interface {
	Add(element any)
	Remove(element any) (any, error)
	Find(element any) (any, error)
	Length() int
}

type ListType string

const (
	Array      ListType = "array"
	LinkedList ListType = "linkedlist"
	Queue      ListType = "queue"
	Stack      ListType = "stack"
)

func New(list ListType) List {
	switch list {
	case Array:
		return array.New()
	case LinkedList:
		return linkedlist.New()
	default:
		return nil
	}
}
