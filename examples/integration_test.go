package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/go-jsonnet"
	gjf "github.com/harsimranmaan/go-jsonnet-func"
)

var vm *jsonnet.VM

func TestMain(m *testing.M) {
	vm = jsonnet.MakeVM()
	vm.NativeFunction(gjf.SHA256())
	vm.NativeFunction(gjf.ParseURL())
	os.Exit(m.Run())
}

func ExampleParseURL() {
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
	output, _ := vm.EvaluateSnippet("main.jsonnet", `local a = std.native("sha256")("test"); {out:a}`)
	fmt.Println(output)
	// Output: {
	//    "out": "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	// }
}
