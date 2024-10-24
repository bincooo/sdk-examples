package annotation

import (
	"github.com/bincooo/sdk/gen/annotation"
	"go/ast"
)

type PostMapping struct {
	Path string `annotation:"name=path,default=/"`
}

var _ annotation.M = (*PostMapping)(nil)

func (g PostMapping) Name() string {
	var m annotation.M = g
	for {
		if n := m.As(); n == nil {
			break
		} else {
			m = n
		}
	}
	return m.Name()
}

func (g PostMapping) Match(node ast.Node) error {
	var m annotation.M = g
	for {
		if n := m.As(); n == nil {
			break
		} else {
			m = n
		}
	}
	return m.Match(node)
}

func (g PostMapping) As() annotation.M {
	return annotation.Router{
		Method: "POST",
		Path:   g.Path,
	}
}
