/*
ECS 140a Hw2
Faustine Yiu 913062973
*/
package nfa

import (
	"testing"
)

func dagTransitions(st state, sym rune) []state {
	/*
	 * 0 -a-> 1
	 * 0 -a-> 2
	 * 1 -b-> 3
	 * 2 -c-> 3
	 */
	return map[state]map[rune][]state{
		0: map[rune][]state{
			'a': []state{1, 2},
		},
		1: map[rune][]state{
			'b': []state{3},
		},
		2: map[rune][]state{
			'c': []state{3},
		},
	}[st][sym]
}

func expTransitions(st state, sym rune) []state {
	/*
	 * 0 -a-> 1
	 * 0 -a-> 2
	 * 0 -b-> 2
	 * 1 -b->0
	 */
	return map[state]map[rune][]state{
		0: map[rune][]state{
			'a': []state{1, 2},
			'b': []state{2},
		},
		1: map[rune][]state{
			'b': []state{0},
		},
		2: map[rune][]state{},
	}[st][sym]
}

func langTransitions(st state, sym rune) []state {
	/*
	 * 0 -a-> 0
	 * 0 -b-> 1
	 * 1 -a-> 1
	 * 1 -b-> 0
	 */
	return map[state]map[rune][]state{
		0: map[rune][]state{
			'a': []state{0},
			'b': []state{1},
		},
		1: map[rune][]state{
			'a': []state{1},
			'b': []state{0},
		},
	}[st][sym]
}

func TestReachable(t *testing.T) {
	nfas := map[string]TransitionFunction{
		"dagTransitions":  dagTransitions,
		"expTransitions":  expTransitions,
		"langTransitions": langTransitions,
	}

	tests := []struct {
		nfa          string
		start, final state
		input        []rune
		expected     bool
	}{
		{"dagTransitions", 0, 3, []rune{'a', 'b'}, true},
		{"dagTransitions", 0, 3, []rune{'a', 'c'}, true},
		{"dagTransitions", 0, 1, []rune{'a'}, true},
		{"dagTransitions", 0, 0, nil, true},
		{"dagTransitions", 0, 3, []rune{'a', 'a'}, false},
		{"dagTransitions", 0, 3, []rune{'a'}, false},
		{"dagTransitions", 0, 1, []rune{'b'}, false},
		{"dagTransitions", 0, 0, []rune{'b'}, false},

		{"expTransitions", 0, 0, []rune{'a', 'b'}, true},
		{"expTransitions", 0, 2, []rune{'a', 'b', 'a'}, true},
		{"expTransitions", 0, 2, []rune{'a', 'b', 'a', 'b', 'a'}, true},
		{"expTransitions", 0, 0, []rune{'a', 'a'}, false},
		{"expTransitions", 0, 2, []rune{'a', 'b', 'a', 'b'}, false},

		{"langTransitions", 0, 0, []rune{'a', 'b', 'b'}, true},
		{"langTransitions", 0, 1, []rune{'a', 'a', 'b'}, true},
		{"langTransitions", 0, 0, []rune{'a', 'a', 'a', 'a', 'a'}, true},
		{"langTransitions", 0, 0, nil, true},
		{"langTransitions", 0, 1, []rune{'a', 'a'}, false},
		{"langTransitions", 0, 0, []rune{'a', 'b', 'a', 'a'}, false},

		// TODO add more tests for 100% test coverage
		{"dagTransitions", 0, 1, nil, false},
		{"dagTransitions", 0, 0, nil, true},
		{"dagTransitions", 0, 3, nil, false},
		{"dagTransitions", 0, 2, nil, false},
		{"dagTransitions", 0, 2, []rune{'a'}, true},
		{"dagTransitions", 0, 2, []rune{'b'}, false},
		{"dagTransitions", 0, 3, []rune{'a','c','c'}, false},
		{"dagTransitions", 1, 1, nil, true},
		{"dagTransitions", 1, 0, nil, false},
		{"dagTransitions", 1, 3, nil, false},
		{"dagTransitions", 1, 2, nil, false},
		{"dagTransitions", 1, 1, []rune{'b'}, false},
		{"dagTransitions", 1, 3, []rune{'a'}, false},
		{"dagTransitions", 1, 3, []rune{'b'}, true},
		{"dagTransitions", 1, 3, []rune{'c'}, false},
		{"dagTransitions", 1, 1, []rune{'a'}, false},
		{"dagTransitions", 1, 0, nil, false},
		{"dagTransitions", 1, 1, nil, true},
		{"dagTransitions", 1, 0, []rune{'a'}, false},
		{"dagTransitions", 1, 0, []rune{'b'}, false},
		{"dagTransitions", 1, 0, []rune{'c'}, false},
		{"dagTransitions", 1, 3, []rune{'b', 'b'}, false},
		{"dagTransitions", 2, 3, []rune{'a'}, false},
		{"dagTransitions", 2, 3, []rune{'b'}, false},
		{"dagTransitions", 2, 3, []rune{'c'}, true},
		{"dagTransitions", 2, 0, nil, false},
		{"dagTransitions", 2, 2, nil, true},
		{"dagTransitions", 2, 3, nil, false},
		{"dagTransitions", 2, 1, nil, false},
		{"dagTransitions", 2, 0, []rune{'a'}, false},
		{"dagTransitions", 2, 0, []rune{'b'}, false},
		{"dagTransitions", 2, 0, []rune{'c'}, false},
		{"dagTransitions", 2, 1, []rune{'a'}, false},
		{"dagTransitions", 2, 1, []rune{'b'}, false},
		{"dagTransitions", 2, 1, []rune{'c'}, false},

		{"expTransitions", 0, 0, nil, true},
		{"expTransitions", 0, 2, nil, false},
		{"expTransitions", 0, 1, nil, false},
		{"expTransitions", 0, 2, []rune{'b', 'b'}, false},
		{"expTransitions", 1, 0, []rune{'b'}, true},
		{"expTransitions", 1, 0, []rune{'a'}, false},
		{"expTransitions", 1, 1, nil, true},
		{"expTransitions", 1, 2, nil, false},
		{"expTransitions", 1, 0, nil, false},
		{"expTransitions", 1, 2, []rune{'b', 'a', 'b', 'a', 'b','a'}, true},
		{"expTransitions", 1, 2, []rune{'a', 'b', 'a', 'b','a'}, false},
		{"expTransitions", 1, 2, []rune{'b', 'b', 'b', 'b','b'}, false},
		{"expTransitions", 2, 0, []rune{'a'}, false},
		{"expTransitions", 2, 0, []rune{'b'}, false},
		{"expTransitions", 2, 1, []rune{'a'}, false},
		{"expTransitions", 2, 1, []rune{'b'}, false},
		{"expTransitions", 2, 1, nil, false},
		{"expTransitions", 2, 2, nil, true},
		{"expTransitions", 2, 0, nil, false},

		{"langTransitions", 0, 1, []rune{'a', 'b', 'b'}, false},
		{"langTransitions", 0, 1, []rune{'a', 'b', 'a','a'}, true},
		{"langTransitions", 0, 1, []rune{'b', 'a','a'}, true},
		{"langTransitions", 0, 1, []rune{'b'}, true},
		{"langTransitions", 1, 1, nil, true},
		{"langTransitions", 1, 0, nil, false},
		{"langTransitions", 1, 1, []rune{'a', 'a', 'a', 'a', 'a'}, true},
		{"langTransitions", 1, 1, []rune{'a', 'b', 'a', 'b', 'a'}, true},
		{"langTransitions", 1, 1, []rune{'a', 'b'}, false},
		{"langTransitions", 1, 1, []rune{'a', 'b', 'a', 'a'}, false},
	}

	for _, test := range tests {
		func() {
			defer func() {
				if recover() != nil {
					t.Errorf("Reachable panicked on (%s, %d, %d, %v); expected %t",
						test.nfa, test.start, test.final, string(test.input),
						test.expected)
				}
			}()
			actual := Reachable(nfas[test.nfa], test.start, test.final, test.input)
			if test.expected != actual {
				t.Errorf("Reachable failed on (%s, %d, %d, %v); expected %t, actual %t",
					test.nfa, test.start, test.final, string(test.input),
					test.expected, actual)
			}
		}()
	}
}
