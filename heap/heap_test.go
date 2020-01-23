package heap

import (
	"fmt"
	"testing"
)

func TestMinHeap_Insert(t *testing.T) {
	tests := []struct {
		name string
		heap *Heap
		arg  int
	}{
		{"one insert", CreateMinHeap(), 5},
		{"two inserts", CreateMinHeap(3), 5},
		{"same inserts", CreateMinHeap(3, 3, 3), 3},
		{"many inserts", CreateMinHeap(11, 3, 4, 8, 5), 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.Insert(tt.arg)
			fmt.Println(tt.heap.a)
			if !tt.heap.IsHeap() {
				t.Errorf("Heap.IsHeap() was false after inserting %v", tt.arg)
			}
		})
	}
}
func TestMaxHeap_Insert(t *testing.T) {
	tests := []struct {
		name string
		heap *Heap
		arg  int
	}{
		{"one insert", CreateMaxHeap(), 5},
		{"two inserts", CreateMaxHeap(3), 5},
		{"same inserts", CreateMaxHeap(3, 3, 3), 3},
		{"many inserts", CreateMaxHeap(3, 4, 8, 5, 11), 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.Insert(tt.arg)
			if !tt.heap.IsHeap() {
				t.Errorf("Heap.IsHeap() was false after inserting %v", tt.arg)
			}
		})
	}
}

func TestMinHeap_Extract(t *testing.T) {
	tests := []struct {
		name string
		heap *Heap
		want int
	}{
		{"extract from 1", CreateMinHeap(5), 5},
		{"extract from 2", CreateMinHeap(10, 5), 5},
		{"extract from many", CreateMinHeap(11, 3, 4, 8, 5), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.heap.Extract(); got != tt.want {
				t.Errorf("Heap.Extract() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.heap.a)
			if !tt.heap.IsHeap() {
				t.Errorf("Heap.IsHeap() was false after extracting")
			}
		})
	}
}
func TestMaxHeap_Extract(t *testing.T) {
	tests := []struct {
		name string
		heap *Heap
		want int
	}{
		{"extract from 1", CreateMaxHeap(5), 5},
		{"extract from 2", CreateMaxHeap(5, 10), 10},
		{"extract from many", CreateMaxHeap(3, 4, 8, 5, 11), 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.heap.Extract(); got != tt.want {
				t.Errorf("Heap.Extract() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.heap.a)
			if !tt.heap.IsHeap() {
				t.Errorf("Heap.IsHeap() was false after extracting")
			}
		})
	}
}

func TestHeap_Size(t *testing.T) {
	tests := []struct {
		name string
		heap *Heap
		want int
	}{
		{"one insert", CreateMaxHeap(), 0},
		{"two inserts", CreateMaxHeap(3), 1},
		{"same inserts", CreateMaxHeap(3, 3, 3), 3},
		{"many inserts", CreateMaxHeap(3, 4, 8, 5, 11), 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.heap.Size(); got != tt.want {
				t.Errorf("Heap.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
