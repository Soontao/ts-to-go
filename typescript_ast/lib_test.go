package typescript_ast

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
)

func TestSimpleParse(t *testing.T) {
	assert := assert.New(t)
	program := Parse("const a: number = 1")
	assert.NotNil(program)
	assert.Nil(program.GetParser().GetError())
	antlr.ParseTreeWalkerDefault.Walk(&BaseTypeScriptParserListener{}, program)
}
