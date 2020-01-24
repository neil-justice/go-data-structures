package linkedlist

import (
	"strings"
)

// Node singly linked list node
type Node struct {
	v    string
	next *Node
}

// LinkedList singly-linked list of strings which tracks its tail node
// (so that it can be used as a backing for a queue)
type LinkedList struct {
	head *Node
	tail *Node
}

// CreateLinkedList Create a linked list with the given entries
// in the same order
func CreateLinkedList(entries ...string) *LinkedList {
	list := &LinkedList{}
	for _, s := range entries {
		list.Append(s)
	}
	return list
}

// Append add an entry to the end of the list
func (list *LinkedList) Append(s string) {
	if list.head == nil {
		list.head = &Node{s, nil}
		list.tail = list.head
	} else {
		list.tail = list.tail.append(s)
	}
}

func (n *Node) append(s string) *Node {
	for n.next != nil {
		n = n.next
	}
	n.next = &Node{s, nil}
	return n.next
}

// Delete delete the first node with the given value from the list
// Returns true if there was a node to delete, otherwise false
func (list *LinkedList) Delete(s string) bool {
	var prev, n *Node = nil, list.head
	for n != nil && n.v != s {
		prev = n
		n = n.next
	}

	if n != nil {
		if prev != nil {
			prev.next = n.next
			if prev.next == nil {
				list.tail = prev
			}
		} else {
			list.head = n.next
		}
		return true
	}
	return false
}

// Pop delete and return the last node from the list
// 2nd return value is true if there was a node to delete, otherwise false
func (list *LinkedList) Pop() (string, bool) {
	if list.head != nil {
		if list.tail == list.head {
			list.tail = list.tail.next
		}
		s := list.head.v
		list.head = list.head.next
		return s, true
	}
	return "", false
}

// Get the nth entry in the list. If n > list.Len() - 1,
// the 2nd return value is false
func (list *LinkedList) Get(index int) (string, bool) {
	if list.head != nil {
		return list.head.get(index)
	}
	return "", false
}

func (n *Node) get(index int) (string, bool) {
	i := 0
	for n != nil {
		if i == index {
			return n.v, true
		}
		n = n.next
		i++
	}
	return "", false
}

// GetKLast the nth last entry in the list. If n > list.Size() - 1,
// the 2nd return value is false.
// if n == 1, returns the last element.
// if n == 2, returns the 2nd-to-last element
func (n *Node) GetKLast(index int) (string, bool) {
	if index <= 0 {
		return n.v, false
	}
	p1 := n
	p2 := n
	for i := 0; i < index; i++ {
		if p1 != nil {
			p1 = p1.next
		} else {
			return "", false
		}
	}
	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2.v, true
}

// ForEach run a function on each entry in the list
func (n *Node) ForEach(fn func(string)) {
	for n != nil {
		fn(n.v)
		n = n.next
	}
}

// Len the list length
func (n *Node) Len() int {
	i := 0
	n.ForEach(func(s string) { i++ })
	return i
}

func (list *LinkedList) String() string {
	return list.head.string()
}

func (n *Node) string() string {
	var sb strings.Builder
	sb.WriteString("[")
	for n != nil {
		sb.WriteString(n.v)
		if n.next != nil {
			sb.WriteString(" ")
		}
		n = n.next
	}
	sb.WriteString("]")
	return sb.String()
}

// ToSlice copy the list contents to a slice
func (list *LinkedList) ToSlice() []string {
	slice := make([]string, 0)
	list.head.ForEach(func(s string) { slice = append(slice, s) })
	return slice
}

// IsEmpty returns true iff there are no elements in the list
func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}
