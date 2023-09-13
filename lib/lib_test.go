package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleParse(t *testing.T) {
	assert := assert.New(t)
	program := Parse("const a: number = 1")
	assert.NotNil(program)
	program.Accept(BaseTypeScriptParserVisitor{})
}
