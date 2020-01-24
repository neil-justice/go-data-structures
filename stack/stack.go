package stack

// Stack an array-backed stack.
type Stack struct {
	contents []string
}

// CreateStack create a stack with the given elements, with the first
// element at the top of the stack.
func CreateStack(s ...string) *Stack {
	return &Stack{s}
}

// Pop remove the topmost element of the stack and return it.
// If the stack is empty returns (nil, false)
func (stack *Stack) Pop() (string, bool) {
	if len(stack.contents) > 0 {
		s := stack.contents[0]
		stack.contents = stack.contents[1:]
		return s, true
	}
	return "", false
}

// Push add an element to the top of the stack
func (stack *Stack) Push(s string) {
	stack.contents = append(stack.contents, s)
}

// Peek returns the top element of the stack if there is one,
// or (nil, false) if the stack is empty.
func (stack *Stack) Peek() (string, bool) {
	if len(stack.contents) > 0 {
		return stack.contents[0], true
	}
	return "", false
}

// IsEmpty returns true iff there is nothing on the stack
func (stack *Stack) IsEmpty() bool {
	return len(stack.contents) == 0
}

// Size number of elements in the stack
func (stack *Stack) Size() int {
	return len(stack.contents)
}
