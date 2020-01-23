package quickSort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Empty case", []int{}, []int{}},
		{"One element", []int{1}, []int{1}},
		{"Test Partition", []int{1, 7, 5, 3, 4}, []int{1, 3, 4, 5, 7}},
		{"Many elements", []int{11, 3, 31, 4}, []int{3, 4, 11, 31}},
		{"Odd num of elements", []int{5, 1, 3}, []int{1, 3, 5}},
		{"All the same", []int{5, 5, 5}, []int{5, 5, 5}},
		{"Already sorted", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
