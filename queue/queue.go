package queue

import (
	"dsa/linkedList"
)

type Queue struct {
	list *linkedList.LinkedList
}

func (q *Queue) Add(s string) {
	q.list.Append(s)
}

func (q *Queue) Remove() (string, bool) {
	return q.list.Pop()
}

func (q *Queue) Peek() (string, bool) {
	return q.list.Get(0)
}

func (q *Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue) ToSlice() []string {
	return q.list.ToSlice()
}
