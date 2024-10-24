package annotation

import (
	"github.com/bincooo/sdk/gen/annotation"
	"go/ast"
)

type DelMapping struct {
	Path string `annotation:"name=path,default=/"`
}

var _ annotation.M = (*DelMapping)(nil)

func (g DelMapping) Name() string {
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

func (g DelMapping) Match(node ast.Node) error {
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

func (g DelMapping) As() annotation.M {
	return annotation.Router{
		Method: "DELETE",
		Path:   g.Path,
	}
}
