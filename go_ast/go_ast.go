package go_ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
)

func GenerateGoCode(f *ast.File) (string, error) {
	fs := token.NewFileSet()
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fs, f); err != nil {
		fmt.Println("Error generating Go code:", err)
		return "", err
	}
	return buf.String(), nil
}
