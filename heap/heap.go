package heap

type Heap struct {
	a        []int
	lessThan func(int, int) bool
}

func childrenIndices(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func CreateMaxHeap(values ...int) *Heap {
	a := make([]int, 0, len(values))
	h := &Heap{a, func(a int, b int) bool {
		return a > b
	}}
	for _, v := range values {
		h.Insert(v)
	}
	return h
}

func CreateMinHeap(values ...int) *Heap {
	a := make([]int, 0, len(values))
	h := &Heap{a, func(a int, b int) bool {
		return a < b
	}}
	for _, v := range values {
		h.Insert(v)
	}
	return h
}

func (h *Heap) Insert(value int) {
	h.a = append(h.a, value)
	index := len(h.a) - 1
	pIndex := parentIndex(index)
	for index > 0 && h.lessThan(h.a[index], h.a[pIndex]) {
		h.a[index], h.a[pIndex] = h.a[pIndex], h.a[index]
		index = pIndex
		pIndex = parentIndex(index)
	}
}

func (h *Heap) Extract() (int, bool) {
	if h.Size() == 0 {
		return 0, false
	}
	res := h.a[0]
	h.a[0] = h.a[h.Size()-1]
	h.a = h.a[:h.Size()-1]
	index := 0
	c1Index, c2Index := childrenIndices(index)
	for (c1Index < h.Size() && h.lessThan(h.a[c1Index], h.a[index])) ||
		(c2Index < h.Size() && h.lessThan(h.a[c2Index], h.a[index])) {
		if h.lessThan(h.a[c1Index], h.a[c2Index]) {
			h.a[index], h.a[c1Index] = h.a[c1Index], h.a[index]
			index = c1Index
		} else {
			h.a[index], h.a[c2Index] = h.a[c2Index], h.a[index]
			index = c2Index
		}
		c1Index, c2Index = childrenIndices(index)
	}
	return res, true
}

func (h *Heap) Size() int {
	return len(h.a)
}

func (h *Heap) Peek() (int, bool) {
	if h.Size() == 0 {
		return 0, false
	} else {
		return h.a[0], true
	}
}

func (h *Heap) IsHeap() bool {
	for index := range h.a {
		c1Index, c2Index := childrenIndices(index)
		if c1Index < h.Size() && h.lessThan(h.a[c1Index], h.a[index]) {
			return false
		}
		if c2Index < h.Size() && h.lessThan(h.a[c2Index], h.a[index]) {
			return false
		}
	}
	return true
}
