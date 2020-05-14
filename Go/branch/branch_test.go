/*
ECS 140a Hw2
Faustine Yiu 913062973
*/
package branch

import (
	"io/ioutil"
	"testing"
)

func TestComputeBranchFactors(t *testing.T) {
	var testCode = `
	package main

	import (
		"fmt"
		"eval"
	)

	func sequential_if() {
		if true {
			fmt.Println("is true")
		} else {
			fmt.Println("isn't true")
		}

		if true {
			fmt.Println("is true")
		} else {
			fmt.Println("isn't true")
		}

		if true {
			fmt.Println("is true")
		} else {
			fmt.Println("isn't true")
		}

		if true {
			fmt.Println("is true")
		} else {
			fmt.Println("isn't true")
		}
	}

	func single_for() {
		for i := 0; i < 10; i += 1 {
			fmt.Println("is loop")
		}

		for i := 0; i < 10; i += 1 {
			fmt.Println("is loop")
		}

		for i := 0; i < 10; i += 1 {
			fmt.Println("is loop")
		}
	}

	func single_range() {
		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}

		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}

		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}

		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}

		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}
	}

	func no_branches() {
		return 42
	}

	func single_switch() {
		switch 5 {
		case 0:
			// pass
		case 5:
			fmt.Println("It's five!")
		default:
			fmt.Println("It isn't five...")
		}

		switch 5 {
		case 0:
			// pass
		case 5:
			fmt.Println("It's five!")
		default:
			fmt.Println("It isn't five...")
		}
	}

	func single_typeswitch() {
		var x interface{} = "test"
		switch x.(type) {
		case uint:
			// pass
		case string:
			fmt.Println("It's a string!")
		default:
			fmt.Println("It's not a string...")
		}
		switch x.(type) {
		case uint:
			// pass
		case string:
			fmt.Println("It's a string!")
		default:
			fmt.Println("It's not a string...")
		}
		switch x.(type) {
		case uint:
			// pass
		case string:
			fmt.Println("It's a string!")
		default:
			fmt.Println("It's not a string...")
		}
	}

	func nested_if(x int) float64 {
		var result float64
		if x < 0 {
			if x > -5 {
				result := -0.5
			} else {
				result := -1
			}
		} else if x > 0 {
			result := 1
		} else {
			result := 0
		}
		return result
	}

	func nested_for_if() {
		for i := 0; i < 10; i += 1 {
			if i > 5 {
				fmt.Println("is filter")
			}
		}
	}

	func nested_switch_if(x int) {
		switch x > 5 {
		case true:
			if x > 10 {
				fmt.Println("is really big")
			}
		default:
			fmt.Println("is default")
		}
	}

	func mixed_switch_no_default_for_if() {
		switch 5 {
		case 0:
			// pass
		case 5:
			fmt.Println("It's five!")
		}

		for i := 0; i < 10; i += 1 {
			if i > 5 {
				fmt.Println("is filter")
			}
		}
	}

	func single_typeswitch_no_default() {
		var x interface{} = "test"
		switch x.(type) {
		case uint:
			// pass
		case string:
			fmt.Println("It's a string!")
		}
	}

	func nested_if_no_else(x int) float64 {
		var result float64 = 0
		if x < 0 {
			if x > -5 {
				result := -0.5
			} else {
				result := -1
			}
		} else if x > 0 {
			result := 1
		}
		return result
	}
	`

	tests := []struct {
		funcName string
		branches uint
	}{
		{"sequential_if", 4},
		{"single_for", 3},
		{"single_range", 5},
		{"no_branches", 0},
		{"single_switch", 2},
		{"single_typeswitch", 3},
		{"nested_if", 3},
		{"nested_for_if", 2},
		{"nested_switch_if", 2},
		{"mixed_switch_no_default_for_if", 3},
		{"single_typeswitch_no_default", 1},
		{"nested_if_no_else", 3},
	}

	branchFactors := ComputeBranchFactors(testCode)

	for _, test := range tests {
		if branchFactors[test.funcName] != test.branches {
			t.Errorf("ComputeBranchFactors(%v) = %d, expected %d\n",
				test.funcName, branchFactors[test.funcName], test.branches)
		}
	}
}
func TestSimplifyParseAndEval_Failure(t *testing.T) {
	tests := []struct {
		in_filename string
	}{
		{"test4_in.go"},
	}

	for _, test := range tests {
		in_file_raw, err := ioutil.ReadFile("../rewrite_tests/" + test.in_filename)
		if err != nil {
			panic("Unable to read test input: " + test.in_filename)
		}

		in_file := string(in_file_raw)

		func() {
			defer func() {
				if recover() == nil {
					t.Errorf("SimplifyParseAndEval()[%v] did not panic, but should\n",
						test.in_filename)
				}
			}()

			ComputeBranchFactors(in_file)
		}()
	}
}
