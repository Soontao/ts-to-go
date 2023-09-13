package lib

import "github.com/antlr4-go/antlr/v4"

func Parse(code string) IProgramContext {
	p := NewTypeScriptParser(antlr.NewCommonTokenStream(NewTypeScriptLexer(antlr.NewInputStream("const a: number = 1")), antlr.TokenDefaultChannel))
	return p.Program()
}
