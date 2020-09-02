package jsonnet_test

import (
	"fmt"
	"testing"

	f "github.com/harsimranmaan/go-jsonnet-func"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSHA256(t *testing.T) {
	var tests = []struct {
		input       []interface{}
		output      interface{}
		expectedErr string
	}{
		{nil, "", "bad arguments to sha256: needs 1 argument"},
		{[]interface{}{""}, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", ""},
		{[]interface{}{`{ "a": 1, "b":"def"}`}, "70e1e7e546a1e3b61733f597295d7227e94d68d59b42f8b1a57155ccb593bb1d", ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("sha256-%d", i), func(t *testing.T) {
			ret, err := f.SHA256().Func(test.input)
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
