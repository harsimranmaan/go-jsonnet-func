package jsonnet

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

// SHA256 returns the sha256 string representation of the input value
func SHA256() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "sha256",
		Params: ast.Identifiers{"value"},
		Func: func(dataString []interface{}) (res interface{}, err error) {
			if len(dataString) != 1 {
				return nil, fmt.Errorf("bad arguments to sha256: needs %d argument", 1)
			}
			data, ok := dataString[0].(string)
			if !ok {
				return nil, errors.New("sha256 failed to read input")
			}
			h := sha256.New()
			_, err = h.Write([]byte(data))
			return fmt.Sprintf("%x", h.Sum(nil)), err
		},
	}
}
