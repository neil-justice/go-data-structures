package linkedList

import (
	"fmt"
	"strings"
)

type Node struct {
	v    string
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func CreateLinkedList(entries ...string) *LinkedList {
	list := &LinkedList{}
	for _, s := range entries {
		list.Append(s)
	}
	return list
}

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

func (list *LinkedList) Pop() (string, bool) {
	if list.head != nil {
		if list.tail == list.head {
			list.tail = list.tail.next
		}
		s := list.head.v
		list.head = list.head.next
		return s, true
	} else {
		return "", false
	}
}

func (list *LinkedList) Get(index int) (string, bool) {
	if list.head != nil {
		return list.head.Get(index)
	} else {
		return "", false
	}
}

func (n *Node) Get(index int) (string, bool) {
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

func (n *Node) ForEach(fn func(string)) {
	for n != nil {
		fn(n.v)
		n = n.next
	}
}

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

func (list *LinkedList) ToSlice() []string {
	slice := make([]string, 0)
	list.head.ForEach(func(s string) { slice = append(slice, s) })
	return slice
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func main() {
	list := CreateLinkedList("a", "b", "c")
	list.Append("d")
	list.Delete("a")
	fmt.Println(list.head.GetKLast(0))
	fmt.Println(list.head)
	fmt.Println(list.ToSlice())
}
