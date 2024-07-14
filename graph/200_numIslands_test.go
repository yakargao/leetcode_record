package graph

import "testing"

func TestNumIslands(t *testing.T) {
	testCases := []struct {
		grid     [][]byte
		expected int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		// Add more test cases if necessary
	}

	for _, tc := range testCases {
		actual := numIslands(tc.grid)
		if actual != tc.expected {
			t.Errorf("numIslands(%v) = %d; expected %d", tc.grid, actual, tc.expected)
		}
	}
}
