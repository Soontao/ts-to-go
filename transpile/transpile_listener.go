package transpile

import (
	"go/ast"

	"fornever.org/ts-to-go/typescript_ast"
)

type TranspileListener struct {
	typescript_ast.BaseTypeScriptParserListener
	tsAst *typescript_ast.ProgramContext
	goAst *ast.File
}

func (l *TranspileListener) GetAst() *ast.File {
	return l.goAst
}
