package dfa

import "github.com/t14raptor/go-fast/ast"

type UseDef struct {
	Usage       *ast.Identifier
	Definitions []*ScopeDef
}
