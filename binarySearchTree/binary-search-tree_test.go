package binarySearchTree

import (
	"reflect"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	tests := []struct {
		name string
		bst  *BST
		arg  int
		want []int
	}{
		{"smaller", CreateBST(10), 5, []int{5, 10}},
		{"larger", CreateBST(10), 15, []int{10, 15}},
		{"deep", CreateBST(10, 3, 21, 4, 15), 16, []int{3, 4, 10, 15, 16, 21}},
		{"same", CreateBST(10), 10, []int{10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bst.Insert(tt.arg)
			if got := tt.bst.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() gave BST %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Contains(t *testing.T) {
	tests := []struct {
		name string
		bst  *BST
		arg  int
		want bool
	}{
		{"root true", CreateBST(10), 10, true},
		{"root false", CreateBST(10), 11, false},
		{"deep true", CreateBST(10, 3, 15, 11, 6, 4, 22, 8), 11, true},
		{"deep false", CreateBST(10, 3, 15, 11, 6, 4, 22, 8), 211, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bst.Contains(tt.arg); got != tt.want {
				t.Errorf("BST.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_Delete(t *testing.T) {
	tests := []struct {
		name string
		bst  *BST
		arg  int
		want []int
	}{
		{"no children", CreateBST(10, 5, 15), 5, []int{10, 15}},
		{"left child", CreateBST(10, 5, 15, 3), 5, []int{3, 10, 15}},
		{"right child", CreateBST(10, 5, 15, 7), 5, []int{7, 10, 15}},
		{"2 children", CreateBST(10, 5, 15, 3, 7), 5, []int{3, 7, 10, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bst.Delete(tt.arg)
			if got := tt.bst.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BST.Delete() gave BST %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBST_findMin(t *testing.T) {
	tests := []struct {
		name string
		bst  *BST
		want int
	}{
		{"1 node", CreateBST(10), 10},
		{"many nodes", CreateBST(10, 5, 15, 3), 3},
		{"degenerate", CreateBST(5, 4, 3, 2, 1), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bst.findMin().Key; got != tt.want {
				t.Errorf("BST.findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
