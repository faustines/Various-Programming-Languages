/*
ECS 140a Hw2
Faustine Yiu 913062973
Source: TA: Tara Hariri Office Hour on 1/29
*/
package depth

import "hw2/expr"
import "fmt"

// Depth returns the maximum number of AST nodes between the
// root of the tree and any leaf (literal or variable) in the tree.
func Depth(e expr.Expr) uint {
	// TODO: implement this function
	switch v := e.(type) {

			case expr.Literal:
				return 1

			case expr.Var:
				return 1

			case expr.Unary:
				return 1 + Depth(v.X)

			case expr.Binary:
				if Depth(v.X) > Depth(v.Y) {
					return 1 + Depth(v.X)
				}else{
					return 1 + Depth(v.Y)
				}

			default:
				panic(fmt.Sprintf("unknown Expr: %T", e))
	}
}
