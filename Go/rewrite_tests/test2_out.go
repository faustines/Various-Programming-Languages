package main

import (
	"fmt"
	"hw2/expr"
)

func main() {
	var result float64

	result = expr.ParseAndEval("1 + 2")
	fmt.Printf("%d\n", result)
}
