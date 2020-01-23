package quickSort

func QuickSort(a []int) []int {
	if len(a) > 1 {
		pIdx := partition(a)
		first, second := QuickSort(a[:pIdx]), QuickSort(a[pIdx:])
		return append(first, second...)
	} else {
		return a
	}
}

func partition(a []int) int {
	pIdx := len(a) - 1
	pivot := a[pIdx]
	for i := 0; i < pIdx; i++ {
		if a[i] > pivot {
			if a[pIdx-1] > a[pIdx] {
				a[pIdx], a[pIdx-1] = a[pIdx-1], a[pIdx]
				pIdx--
			}
			// now a[pIdx-1] guaranteed <= pivot
			a[i], a[pIdx-1] = a[pIdx-1], a[i]
			a[pIdx], a[pIdx-1] = a[pIdx-1], a[pIdx]
			pIdx--
		}
	}
	return pIdx
}
