package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/Masterminds/sprig/v3"
)

var allowedFunctions = map[string]bool{
	"hello": true,

	"upper": true,
}

func main() {
	var filename string
	flag.StringVar(&filename, "file", "generated.gen.go", "generated file name")
	flag.Parse()
	p := ""
	for name, signature := range sprig.GenericFuncMap() {
		if !allowedFunctions[name] {
			continue
		}
		t := reflect.TypeOf(signature)
		if t.Kind() != reflect.Func {
			panic("<not a function>")
		}
		var params []string
		for i := 0; i < t.NumIn(); i++ {
			params = append(params, t.In(i).String())
		}
		var outputs []string
		for i := 0; i < t.NumOut(); i++ {
			outputs = append(outputs, t.Out(i).String())
		}
		p += jsonnetFunc(name, params, outputs)
	}

	content := fmt.Sprintf(`package jsonnet
import (
	"fmt"

	"github.com/Masterminds/sprig/v3"
	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func SprigFuncs() []*jsonnet.NativeFunction{
	return []*jsonnet.NativeFunction{
		%s
	}
}`, p)
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func jsonnetFunc(name string, params, outputs []string) string {
	returnNil := ", nil"
	if len(outputs) == 2 {
		returnNil = ""
	}
	var paramNames, paramNamesQuoted []string
	var typeConversions = []string{fmt.Sprintf(`if len(dataString) != %d {
			return nil, fmt.Errorf("bad arguments to %%q: needs %d argument", %q)
		}`, len(params), len(params), name)}
	for i, dataType := range params {
		paramName := fmt.Sprintf("p%d", i)
		paramNamesQuoted = append(paramNamesQuoted, fmt.Sprintf("%q", paramName))
		paramNames = append(paramNames, paramName)
		typeConversion := fmt.Sprintf(`
		%s, ok := dataString[%d].(%s)
		if !ok {
			return nil, fmt.Errorf("%%q failed to read input param %%q", %q, %q)
		}`, paramName, i, dataType, name, paramName)
		typeConversions = append(typeConversions, typeConversion)
	}

	return fmt.Sprintf(`
	{
		Name:   %q,
		Params: ast.Identifiers{%s},
		Func:   func(dataString []any) (res any, err error) {
			%s
		f, ok := sprig.GenericFuncMap()[%q].(func (%s) (%s))
		if !ok {
			return nil, fmt.Errorf("mismatch function defintion for %%q", %q)
		}
		return f(%s)%s
		},
	},`, name, strings.Join(paramNamesQuoted, `,`), strings.Join(typeConversions, "\n"), name, strings.Join(params, `,`), strings.Join(outputs, `,`), name, strings.Join(paramNames, `,`), returnNil)
}
