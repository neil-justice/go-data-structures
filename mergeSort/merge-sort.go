package mergeSort

func MergeSort(a []int) []int {
	if len(a) > 1 {
		first, second := MergeSort(a[:len(a)/2]), MergeSort(a[len(a)/2:])
		res := make([]int, len(a))
		for i := range a {
			if len(second) == 0 || (len(first) > 0 && first[0] <= second[0]) {
				res[i] = first[0]
				first = first[1:]
			} else {
				res[i] = second[0]
				second = second[1:]
			}
		}
		return res
	} else {
		return a
	}
}

// different merging algorithm
func MergeSort2(a []int) []int {
	if len(a) > 1 {
		first, second := MergeSort(a[:len(a)/2]), MergeSort(a[len(a)/2:])
		res := make([]int, 0, len(a))
		for len(first) > 0 && len(second) > 0 {
			if first[0] <= second[0] {
				res = append(res, first[0])
				first = first[1:]
			} else {
				res = append(res, second[0])
				second = second[1:]
			}
		}
		for len(first) > 0 {
			res = append(res, first[0])
			first = first[1:]
		}
		for len(second) > 0 {
			res = append(res, second[0])
			second = second[1:]
		}
		return res
	} else {
		return a
	}
}
