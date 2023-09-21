package evaluator

import (
	"github.com/xtommas/monkey-interpreter/ast"
	"github.com/xtommas/monkey-interpreter/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
