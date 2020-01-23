package heap

type PriorityQueue struct {
	heap *Heap
}

func (q *PriorityQueue) Add(i int) {
	q.heap.Insert(i)
}

func (q *PriorityQueue) Remove() (int, bool) {
	return q.heap.Extract()
}

func (q *PriorityQueue) Peek() (int, bool) {
	return q.heap.Peek()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.heap.Size() == 0
}
