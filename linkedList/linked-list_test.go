package linkedList

import (
	"testing"
)

func TestAppend(t *testing.T) {
	cases := []struct {
		list   *LinkedList
		append string
		want   string
	}{
		{CreateLinkedList("x"), "a", "[x a]"},
		{CreateLinkedList(), "b", "[b]"},
	}
	for _, c := range cases {
		c.list.Append(c.append)
		got := c.list.String()
		if got != c.want {
			t.Errorf("Append() == %q, want %q", got, c.want)
		}
		if c.append != c.list.tail.v {
			t.Errorf("list.tail.v == %q, want %q", c.append, c.list.tail.v)
		}
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		list  *LinkedList
		want  string
		ok    bool
		index int
	}{
		{CreateLinkedList("a", "b", "c"), "a", true, 0},
		{CreateLinkedList("a", "b", "c"), "c", true, 2},
		{CreateLinkedList("a", "b", "c"), "", false, 3},
	}
	for _, c := range cases {
		got, ok := c.list.head.Get(c.index)
		if got != c.want || ok != c.ok {
			t.Errorf("Get(%d) == (%s, %t) want (%s, %t)", c.index, got, ok, c.want, c.ok)
		}
	}
}

func TestLinkedList_Delete(t *testing.T) {
	tests := []struct {
		name string
		list *LinkedList
		want string
		arg  string
		ok   bool
	}{
		{"Delete head", CreateLinkedList("a", "b", "c"), "[b c]", "a", true},
		{"Delete tail", CreateLinkedList("a", "b", "c"), "[a b]", "c", true},
		{"Delete middle", CreateLinkedList("a", "b", "c"), "[a c]", "b", true},
		{"Doesn't exist", CreateLinkedList("a", "b", "c"), "[a b c]", "d", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.Delete(tt.arg)
			if got != tt.ok {
				t.Errorf("LinkedList.Delete() = %v, want %v", got, tt.ok)
			}
			if tt.list.String() != tt.want {
				t.Errorf("LinkedList.String() = %v, want %v", tt.list, tt.want)
			}
			if last, _ := tt.list.head.GetKLast(1); last != tt.list.tail.v {
				t.Errorf("LinkedList.tail.v = %v, want %v", tt.list.tail.v, last)
			}
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name string
		list *LinkedList
		want string
		left string
		ok   bool
	}{
		{"Pop", CreateLinkedList("a", "b", "c"), "a", "[b c]", true},
		{"Pop only element", CreateLinkedList("a"), "a", "[]", true},
		{"Pop empty", CreateLinkedList(), "", "[]", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.list.Pop()
			if got != tt.want {
				t.Errorf("LinkedList.Delete() = %v, want %v", got, tt.want)
			}
			if ok != tt.ok {
				t.Errorf("LinkedList.Delete() = %v, want %v", ok, tt.ok)
			}
			if tt.list.String() != tt.left {
				t.Errorf("LinkedList.String() = %v, want %v", tt.list, tt.left)
			}
			if last, ok := tt.list.head.GetKLast(1); ok == true && last != tt.list.tail.v {
				t.Errorf("LinkedList.tail.v = %v, want %v", tt.list.tail.v, last)
			}
		})
	}
}
