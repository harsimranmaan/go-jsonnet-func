package main

import (
	"fmt"

	"github.com/google/go-jsonnet"
	gjf "github.com/harsimranmaan/go-jsonnet-func"
)

func main() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(gjf.SHA256())
	output, _ := vm.EvaluateSnippet("main.jsonnet", `local a = std.native("sha256")("some string"); {out:a}`)
	fmt.Println(output)
}
