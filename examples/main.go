package main

import (
	"fmt"

	"github.com/google/go-jsonnet"
	gjf "github.com/harsimranmaan/go-jsonnet-func"
)

func main() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(gjf.SHA256())
	output, _ := vm.EvaluateSnippet("main.jsonnet", `local a = std.native("sha256")("test"); {out:a}`)
	fmt.Println(output)
	// Output: {
	//    "out": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	// }
}
