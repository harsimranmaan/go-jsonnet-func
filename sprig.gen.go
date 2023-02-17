package jsonnet

import (
	"fmt"

	"github.com/Masterminds/sprig/v3"
	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func SprigFuncs() []*jsonnet.NativeFunction {
	return []*jsonnet.NativeFunction{
		{
			Name:   "upper",
			Params: ast.Identifiers{"p0"},
			Func: func(dataString []any) (res any, err error) {
				if len(dataString) != 1 {
					return nil, fmt.Errorf("bad arguments to %q: needs 1 argument", "upper")
				}

				p0, ok := dataString[0].(string)
				if !ok {
					return nil, fmt.Errorf("%q failed to read input param %q", "upper", "p0")
				}
				f, ok := sprig.GenericFuncMap()["upper"].(func(string) string)
				if !ok {
					return nil, fmt.Errorf("mismatch function defintion for %q", "upper")
				}
				return f(p0), nil
			},
		},
		{
			Name:   "hello",
			Params: ast.Identifiers{},
			Func: func(dataString []any) (res any, err error) {
				if len(dataString) != 0 {
					return nil, fmt.Errorf("bad arguments to %q: needs 0 argument", "hello")
				}
				f, ok := sprig.GenericFuncMap()["hello"].(func() string)
				if !ok {
					return nil, fmt.Errorf("mismatch function defintion for %q", "hello")
				}
				return f(), nil
			},
		},
	}
}
