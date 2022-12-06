package linkedlist

import "fmt"

type LinkedList struct {
	head   *Link
	length int
}

func New() *LinkedList {
	return &LinkedList{
		head:   nil,
		length: 0,
	}
}

func (l *LinkedList) Add(element any) {
	if l.length == 0 {
		l.head = NewLink(element, nil)

	} else {
		newHead := NewLink(element, l.head)
		l.head = newHead
	}
	l.length++
}

func (l *LinkedList) Remove(element any) (any, error) {
	if l.length < 1 {
		return nil, fmt.Errorf("unable to remove element. unable to find element '%s'", element)
	}

	if l.head.data == element {
		removed := l.head.data
		l.length--
		l.head = l.head.next
		return removed, nil
	}

	current := l.head
	for current.next != nil {
		if current.next.data == element {
			removed := current.next.data
			l.length--
			current.next = current.next.next
			return removed, nil
		}
		current = current.next
	}
	return nil, fmt.Errorf("unable to remove element. unable to find element '%s'", element)
}

func (l *LinkedList) Find(element any) (any, error) {
	if l.length < 1 {
		return nil, fmt.Errorf("unable to find element '%s'", element)
	}

	current := l.head
	for current != nil {
		if current.data == element {
			return current.data, nil
		}
		current = current.next
	}
	return nil, fmt.Errorf("unable to find element '%s'", element)
}

func (l *LinkedList) Length() int {
	return l.length
}

type Link struct {
	next *Link
	data any
}

func NewLink(element any, next *Link) *Link {
	return &Link{
		next: next,
		data: element,
	}
}
