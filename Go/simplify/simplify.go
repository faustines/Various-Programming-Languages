/*
ECS 140a Hw2
Faustine Yiu 913062973
Source: TA: Tara Hariri Office Hour on 1/29
*/
package simplify

import "hw2/expr"
import "fmt"

// Depth should return the simplified expresion
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	//TODO implement the simplify
	//TODO implement the simplify
	switch v := e.(type) {
		case expr.Literal:
			// return v

		case expr.Var:
			index := e.(expr.Var) //cast
			//check if e is in v
			if val, ok := env[index]; ok {
	    	return expr.Literal(val)
			}
			return e

		case expr.Unary:
			v.X = Simplify(v.X,env)
			//check for v.x expr.Literal
			if _, ok := v.X.(expr.Literal); ok {
	    	return expr.Literal(e.Eval(env))
			}
			return e

		case expr.Binary:
			v.X = Simplify(v.X,env)
			v.Y = Simplify(v.Y,env)
			//check for both v.x and v.y expr.Literal
			if _, ok1 := v.X.(expr.Literal); ok1 {
				if _, ok2 := v.Y.(expr.Literal); ok2 {
					return expr.Literal(e.Eval(env))
				}
			}

			switch v.Op {
				case '+':
					if v.X == expr.Literal(0){
						return v.Y
					}else if v.Y == expr.Literal(0){
						return v.X
					}else if v.X == v.X{
						return v
					}

				case '-':
					if v.X == expr.Literal(0){
						return v.Y
					}else if v.Y == expr.Literal(0){
						return v.X
					}else if v.X == v.X{
						return v
					}

				case '*':
					if v.X == expr.Literal(0) || v.Y == expr.Literal(0) {
						return expr.Literal(0)
					}else if v.X == expr.Literal(1) {
						return v.Y
					}else if v.Y == expr.Literal(1) {
						return v.X
					}else if v.X == v.X {
						return v
					}

				case '/':
					if v.X == expr.Literal(0) || v.Y == expr.Literal(0) {
						return expr.Literal(0)
					}
					if v.X == expr.Literal(1) {
						return v.Y
					}else if v.Y == expr.Literal(1) {
						return v.X
					}else if v.X == v.X {
						return v
					}
			}
		default:
			panic(fmt.Sprintf("unknown Expr: %T", e))
	}
	return e
}
