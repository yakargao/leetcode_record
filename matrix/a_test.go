package matrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		matrix [][]int
		order  []int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			order: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			order: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix: [][]int{},
			order:  []int{},
		},
		{
			matrix: [][]int{
				{1},
			},
			order: []int{1},
		},
	}

	for _, test := range tests {
		result := spiralOrder(test.matrix)
		if !reflect.DeepEqual(result, test.order) {
			t.Errorf("spiralOrder(%v) = %v, want %v", test.matrix, result, test.order)
		}
	}
}
