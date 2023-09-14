package typescript_ast

import "github.com/antlr4-go/antlr/v4"

func Parse(code string) (IProgramContext, antlr.RecognitionException) {
	p := NewTypeScriptParser(antlr.NewCommonTokenStream(NewTypeScriptLexer(antlr.NewInputStream("const a: number = 1")), antlr.TokenDefaultChannel))
	return p.Program(), p.GetError()
}
