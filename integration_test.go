package jsonnet_test

import (
	"fmt"

	"github.com/google/go-jsonnet"
	f "github.com/harsimranmaan/go-jsonnet-func"
)

func ExampleParseURL() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(f.ParseURL())
	output, _ := vm.EvaluateSnippet("main.jsonnet", `local a = std.native("parseUrl")("https://example.com/test?param=1#link"); {out:a}`)
	fmt.Println(output)
	// Output:{
	//    "out": {
	//       "fragment": "link",
	//       "host": "example.com",
	//       "hostname": "example.com",
	//       "opaque": "",
	//       "path": "/test",
	//       "port": "",
	//       "query": "param=1",
	//       "scheme": "https",
	//       "userinfo": ""
	//    }
	// }
}

func ExampleSHA256() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(f.SHA256())
	output, _ := vm.EvaluateSnippet("main.jsonnet", `local a = std.native("sha256")("test"); {out:a}`)
	fmt.Println(output)
	// Output: {
	//    "out": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	// }
}
