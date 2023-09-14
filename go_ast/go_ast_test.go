package go_ast

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateGoCode(t *testing.T) {
	assert := assert.New(t)
	file := &ast.File{
		Name: ast.NewIdent("main"),
		Decls: []ast.Decl{
			&ast.FuncDecl{
				Name: ast.NewIdent("myFunction"),
				Type: &ast.FuncType{},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: ast.NewIdent("fmt.Println"),
								Args: []ast.Expr{
									ast.NewIdent(`"Hello, World!"`),
								},
							},
						},
					},
				},
			},
		},
	}

	// Convert the buffer content to a string.
	code, err := GenerateGoCode(file)
	assert.Nil(err)
	assert.NotNil(code)
}
