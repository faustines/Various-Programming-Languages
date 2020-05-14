/*
ECS 140a Hw2
Faustine Yiu 913062973
*/
package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	tests := []struct {
		a, b, c  int
		expected triangleType
	}{
		{1001, 5, 6, UnknownTriangle},
		// TODO add more tests for 100% test coverage
		{4, 2001, 6, UnknownTriangle},
		{4, 5, 3001, UnknownTriangle},
		{0, 5, 6, UnknownTriangle},
		{-1, 5, 6, UnknownTriangle},
		{4, 0, 6, UnknownTriangle},
		{4, -1, 6, UnknownTriangle},
		{4, 5, 0, UnknownTriangle},
		{4, 5, -1, UnknownTriangle},
		{2, 10, 5, InvalidTriangle},
		{6, 8, 10, RightTriangle},
		{2, 3, 4, AcuteTriangle},
		{3, 5, 4, ObtuseTriangle},
	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
