package stack

type Stack struct {
	contents []string
}

func CreateStack(s ...string) *Stack {
	return &Stack{s}
}

func (stack *Stack) Pop() (string, bool) {
	if len(stack.contents) > 0 {
		s := stack.contents[0]
		stack.contents = stack.contents[1:]
		return s, true
	} else {
		return "", false
	}
}

func (stack *Stack) Push(s string) {
	stack.contents = append(stack.contents, s)
}

func (stack *Stack) Peek() (string, bool) {
	if len(stack.contents) > 0 {
		return stack.contents[0], true
	} else {
		return "", false
	}

}

func (stack *Stack) IsEmpty() bool {
	return len(stack.contents) == 0
}

func (stack *Stack) Size() int {
	return len(stack.contents)
}
