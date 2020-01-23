package mergeSort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"Empty case", []int{}, []int{}},
		{"One element", []int{1}, []int{1}},
		{"Many elements", []int{11, 3, 31, 4}, []int{3, 4, 11, 31}},
		{"Odd num of elements", []int{5, 1, 3}, []int{1, 3, 5}},
		{"Duplicate elements", []int{5, 5, 3}, []int{3, 5, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
			if got := MergeSort2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
