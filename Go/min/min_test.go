/*
ECS 140a Hw2
Faustine Yiu 913062973
*/
package min

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{0, -1, 1, 2, -4}, -4},
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{}, 0},
		{nil, 0},
		// TODO add more tests for 100% test coverage
		{[]int{-1}, -1},
		{[]int{0, 1, 2}, 0},
		{[]int{1, 2, 3}, 1},
	}

	for i, test := range tests {
		actual := Min(test.in)
		if actual != test.expected {
			t.Errorf("#%d: Min(%v)=%d; expected %d", i, test.in, actual, test.expected)
		}
	}
}
