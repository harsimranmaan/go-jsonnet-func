package main

import (
	"flag"
	"fmt"
	"go/format"
	"os"
	"reflect"
	"strings"

	"github.com/Masterminds/sprig/v3"
)

// only add hermitic functions here
var allowedFunctions = []string{
	"hello",
	"upper",
	"snakecase",
	"camelcase",
	"kebabcase",

	// Crypto:
	//"encryptAES",
	"decryptAES",
}

func main() {
	var outputType = make(map[string]bool)
	var filename string
	flag.StringVar(&filename, "file", "generated.gen.go", "generated file name")
	flag.Parse()
	p := ""
	for _, funcName := range allowedFunctions {
		signature, ok := sprig.GenericFuncMap()[funcName]
		if !ok {
			panic(fmt.Sprintf("function not found: %s", funcName))
		}
		t := reflect.TypeOf(signature)
		if t.Kind() != reflect.Func {
			panic(fmt.Sprintf("not a function: %s", funcName))
		}
		var params []string
		for i := 0; i < t.NumIn(); i++ {
			params = append(params, t.In(i).String())
		}
		var outputs []string
		for i := 0; i < t.NumOut(); i++ {
			output := t.Out(i).String()
			outputs = append(outputs, output)
			outputType[output] = true
		}
		p += jsonnetFunc(funcName, params, outputs)
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
	formatted, err := format.Source([]byte(content))
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filename, formatted, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(outputType)
}

func jsonnetFunc(name string, params, outputs []string) string {
	returnErr := ""
	for _, o := range outputs {
		if o == "error" {
			returnErr = ", err"
		}
	}
	if len(outputs) > 2 {
		panic(fmt.Sprintf("upto 1 output variable is supported: %s", name))
	}
	var paramNames, paramNamesQuoted []string
	var typeConversions = []string{fmt.Sprintf(`if len(dataString) != %d {
			return nil, fmt.Errorf("bad arguments to %%q: needs %d argument", %q)
		}`, len(params), len(params), name)}
	for i, dataType := range params {
		paramName := fmt.Sprintf("p%d", i)
		paramNamesQuoted = append(paramNamesQuoted, fmt.Sprintf("%q", paramName))
		paramNames = append(paramNames, paramName)
		// todo: extract method
		dataTypeCast := dataType
		// if dataTypeCast == "int" {
		// 	dataTypeCast = "float64"
		// }
		typeConversion := fmt.Sprintf(`
		%s, ok := dataString[%d].(%s)`, paramName, i, dataTypeCast)
		typeConversion += fmt.Sprintf(`
		if !ok {
			return nil, fmt.Errorf("%%q failed to read input param %%q", %q, %q)
		}`, name, paramName)
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
			return nil, fmt.Errorf("mismatch function definition for %%q", %q)
		}
		o1%s := f(%s)
		return o1, err
		},
	},`, fmt.Sprintf("sprig.%s", name), strings.Join(paramNamesQuoted, `,`), strings.Join(typeConversions, "\n"), name, strings.Join(params, `,`), strings.Join(outputs, `,`), name, returnErr, strings.Join(paramNames, `,`))
}

// TODO: Handle []string to []interface, handle jsonnet number to int conversions
// would unlock a majority of functions
