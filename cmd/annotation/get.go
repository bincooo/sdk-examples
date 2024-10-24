package annotation

import (
	"github.com/bincooo/sdk/gen/annotation"
	"go/ast"
)

type GetMapping struct {
	Path string `annotation:"name=path,default=/"`
}

var _ annotation.M = (*GetMapping)(nil)

func (g GetMapping) Name() string {
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

func (g GetMapping) Match(node ast.Node) error {
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

func (g GetMapping) As() annotation.M {
	return annotation.Router{
		Method: "GET",
		Path:   g.Path,
	}
}
