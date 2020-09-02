package jsonnet_test

import (
	"fmt"
	"testing"

	f "github.com/harsimranmaan/go-jsonnet-func"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseURL(t *testing.T) {
	var tests = []struct {
		input       []interface{}
		output      map[string]interface{}
		expectedErr string
	}{
		{nil, nil, "bad arguments to parseUrl: needs 1 argument"},
		{[]interface{}{"https://me:pwd@example.com:8080/abc?param=1&param=2#frag"}, map[string]interface{}{"fragment": "frag", "host": "example.com:8080", "hostname": "example.com", "opaque": "", "path": "/abc", "port": "8080", "query": "param=1&param=2", "scheme": "https", "userinfo": "me:pwd"}, ""},
		{[]interface{}{"data:text/plain;base64,SGVsbG8sIFdvcmxkIQ=="}, map[string]interface{}{"fragment": "", "host": "", "hostname": "", "opaque": "text/plain;base64,SGVsbG8sIFdvcmxkIQ==", "path": "", "port": "", "query": "", "scheme": "data", "userinfo": ""}, ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("parseUrl-%d", i), func(t *testing.T) {
			ret, err := f.ParseURL().Func(test.input)
			if test.expectedErr == "" {
				require.Nil(t, err)
				assert.Equal(t, test.output, ret)
			} else {
				require.NotNil(t, err)
				assert.Equal(t, test.expectedErr, err.Error())
			}
		})
	}
}
