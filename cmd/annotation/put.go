package annotation

import (
	"github.com/bincooo/sdk/gen/annotation"
	"go/ast"
)

type PutMapping struct {
	Path string `annotation:"name=path,default=/"`
}

var _ annotation.M = (*PutMapping)(nil)

func (g PutMapping) Name() string {
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

func (g PutMapping) Match(node ast.Node) error {
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

func (g PutMapping) As() annotation.M {
	return annotation.Router{
		Method: "PUT",
		Path:   g.Path,
	}
}
