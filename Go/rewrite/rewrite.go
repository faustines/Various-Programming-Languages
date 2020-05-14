/*
ECS 140a Hw2
Faustine Yiu 913062973
Source: https://play.golang.org/p/mRKeN7SZD8g
Source: TA: Tara Hariri Office Hour on 1/31
*/
package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"hw2/expr"
	"hw2/simplify"
	"strconv"
)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	//TODO Write the rewriteCalls function
	var s string
	var emptyenv expr.Env
	var result string

    ast.Inspect(node, func (node ast.Node) bool {
        // We only want to operate on call expressions,
        // so let's filter for those kinds of AST node.
			switch v := node.(type) {
      case *ast.CallExpr:
      	switch fun := v.Fun.(type) {
        case *ast.SelectorExpr:
					if fun.Sel.Name == "ParseAndEval" {
						if len(v.Args) == 2 {
							switch c:= v.Args[0].(type){
							case *ast.BasicLit:
								// fmt.Println(c.Value)
								s,_ = strconv.Unquote(c.Value)
								 if e, err := expr.Parse(s); err == nil {
									 e = simplify.Simplify(e,emptyenv)
									 //Format the resulting expr.Expr into a string.
									 result = expr.Format(e)
									 c.Value = strconv.Quote(result)
								 }
          }
				}else{
					return false
				}
        }
      }
		}
        // If we return true, we keep recursing under this AST node.
        // If we return false, we won't visit anything under this AST node.
        return true
    })
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
