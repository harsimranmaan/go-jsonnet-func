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
			Name:   "sprig.hello",
			Params: ast.Identifiers{},
			Func: func(dataString []any) (res any, err error) {
				if len(dataString) != 0 {
					return nil, fmt.Errorf("bad arguments to %q: needs 0 argument", "hello")
				}
				f, ok := sprig.GenericFuncMap()["hello"].(func() string)
				if !ok {
					return nil, fmt.Errorf("mismatch function definition for %q", "hello")
				}
				o1 := f()
				return o1, err
			},
		},
		{
			Name:   "sprig.upper",
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
					return nil, fmt.Errorf("mismatch function definition for %q", "upper")
				}
				o1 := f(p0)
				return o1, err
			},
		},
		{
			Name:   "sprig.snakecase",
			Params: ast.Identifiers{"p0"},
			Func: func(dataString []any) (res any, err error) {
				if len(dataString) != 1 {
					return nil, fmt.Errorf("bad arguments to %q: needs 1 argument", "snakecase")
				}

				p0, ok := dataString[0].(string)
				if !ok {
					return nil, fmt.Errorf("%q failed to read input param %q", "snakecase", "p0")
				}
				f, ok := sprig.GenericFuncMap()["snakecase"].(func(string) string)
				if !ok {
					return nil, fmt.Errorf("mismatch function definition for %q", "snakecase")
				}
				o1 := f(p0)
				return o1, err
			},
		},
		{
			Name:   "sprig.camelcase",
			Params: ast.Identifiers{"p0"},
			Func: func(dataString []any) (res any, err error) {
				if len(dataString) != 1 {
					return nil, fmt.Errorf("bad arguments to %q: needs 1 argument", "camelcase")
				}

				p0, ok := dataString[0].(string)
				if !ok {
					return nil, fmt.Errorf("%q failed to read input param %q", "camelcase", "p0")
				}
				f, ok := sprig.GenericFuncMap()["camelcase"].(func(string) string)
				if !ok {
					return nil, fmt.Errorf("mismatch function definition for %q", "camelcase")
				}
				o1 := f(p0)
				return o1, err
			},
		},
		{
			Name:   "sprig.kebabcase",
			Params: ast.Identifiers{"p0"},
			Func: func(dataString []any) (res any, err error) {
				if len(dataString) != 1 {
					return nil, fmt.Errorf("bad arguments to %q: needs 1 argument", "kebabcase")
				}

				p0, ok := dataString[0].(string)
				if !ok {
					return nil, fmt.Errorf("%q failed to read input param %q", "kebabcase", "p0")
				}
				f, ok := sprig.GenericFuncMap()["kebabcase"].(func(string) string)
				if !ok {
					return nil, fmt.Errorf("mismatch function definition for %q", "kebabcase")
				}
				o1 := f(p0)
				return o1, err
			},
		},
		{
			Name:   "sprig.decryptAES",
			Params: ast.Identifiers{"p0", "p1"},
			Func: func(dataString []any) (res any, err error) {
				if len(dataString) != 2 {
					return nil, fmt.Errorf("bad arguments to %q: needs 2 argument", "decryptAES")
				}

				p0, ok := dataString[0].(string)
				if !ok {
					return nil, fmt.Errorf("%q failed to read input param %q", "decryptAES", "p0")
				}

				p1, ok := dataString[1].(string)
				if !ok {
					return nil, fmt.Errorf("%q failed to read input param %q", "decryptAES", "p1")
				}
				f, ok := sprig.GenericFuncMap()["decryptAES"].(func(string, string) (string, error))
				if !ok {
					return nil, fmt.Errorf("mismatch function definition for %q", "decryptAES")
				}
				o1, err := f(p0, p1)
				return o1, err
			},
		},
	}
}
