package jsonnet

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

// ParseURL returns a jsonnet function parseUrl which returns a dictionary representing the input url
func ParseURL() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "parseUrl",
		Params: ast.Identifiers{"url"},
		Func: func(dataString []interface{}) (res interface{}, err error) {
			if len(dataString) != 1 {
				return nil, fmt.Errorf("bad arguments to parseUrl: needs %d argument", 1)
			}
			data, ok := dataString[0].(string)
			if !ok {
				return nil, errors.New("parseUrl failed to read input")
			}
			parsedSegments := map[string]interface{}{}
			parsedURL, err := url.Parse(data)
			if err != nil {
				return nil, fmt.Errorf("unable to parseUrl: %s", err)
			}
			parsedSegments["scheme"] = parsedURL.Scheme
			parsedSegments["host"] = parsedURL.Host
			parsedSegments["hostname"] = parsedURL.Hostname()
			parsedSegments["port"] = parsedURL.Port()
			parsedSegments["opaque"] = parsedURL.Opaque
			parsedSegments["path"] = parsedURL.Path
			parsedSegments["query"] = parsedURL.RawQuery
			parsedSegments["fragment"] = parsedURL.Fragment
			if parsedURL.User != nil {
				parsedSegments["userinfo"] = parsedURL.User.String()
			} else {
				parsedSegments["userinfo"] = ""
			}

			return parsedSegments, nil
		},
	}
}
