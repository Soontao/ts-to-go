package typescript_ast

import (
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
)

func TestSimpleParse(t *testing.T) {
	assert := assert.New(t)
	program, err := Parse("const a: number = 1")
	assert.Nil(err)
	assert.NotNil(program)
	antlr.ParseTreeWalkerDefault.Walk(&BaseTypeScriptParserListener{}, program)
}
