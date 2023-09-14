package transpile

import "fornever.org/ts-to-go/typescript_ast"

func Transpile(tsCode string) (goCode string) {
	typescript_ast.Parse(tsCode)
	return
}
