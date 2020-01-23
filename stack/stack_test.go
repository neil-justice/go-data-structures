package stack

import (
	"reflect"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	type fields struct {
		contents []string
	}
	tests := []struct {
		name  string
		stack *Stack
		want  string
		ok    bool
		size  int
	}{
		{"one element", CreateStack("a"), "a", true, 0},
		{"many elements", CreateStack("a", "b", "c"), "a", true, 2},
		{"no elements", CreateStack(), "", false, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := tt.stack
			got, ok := stack.Pop()
			if got != tt.want {
				t.Errorf("Stack.Pop() got = %v, want %v", got, tt.want)
			}
			if ok != tt.ok {
				t.Errorf("Stack.Pop() ok = %v, want %v", ok, tt.ok)
			}
			if stack.Size() != tt.size {
				t.Errorf("Stack.Pop() resulted in size = %d, want %d", stack.Size(), tt.size)
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type fields struct {
		contents []string
	}
	tests := []struct {
		name  string
		stack *Stack
		want  string
		ok    bool
		size  int
	}{
		{"one element", CreateStack("a"), "a", true, 1},
		{"many elements", CreateStack("a", "b", "c"), "a", true, 3},
		{"no elements", CreateStack(), "", false, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := tt.stack
			got, ok := stack.Peek()
			if got != tt.want {
				t.Errorf("Stack.Pop() got = %v, want %v", got, tt.want)
			}
			if ok != tt.ok {
				t.Errorf("Stack.Pop() ok = %v, want %v", ok, tt.ok)
			}
			if stack.Size() != tt.size {
				t.Errorf("Stack.Pop() resulted in size = %d, want %d", stack.Size(), tt.size)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		contents []string
	}
	tests := []struct {
		name  string
		stack *Stack
		push  string
		want  []string
	}{
		{"one element", CreateStack("a"), "a", []string{"a", "a"}},
		{"many elements", CreateStack("a", "b", "c"), "a", []string{"a", "b", "c", "a"}},
		{"no elements", CreateStack(), "a", []string{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := tt.stack
			stack.Push(tt.push)
			got := stack.contents
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}
