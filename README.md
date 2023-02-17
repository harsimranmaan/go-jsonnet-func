# go-jsonnet-func

A collection of [jsonnet](https://jsonnet.org/) native functions that can be easily used in tools written in Golang.

Supports the following functions for jsonnet:

- `sha256`
- `parseUrl`

[Sprig](https://github.com/Masterminds/sprig) functions are supported in experimental stage and can be added to the [code generator](./generate.main.go) to enable those.

The following Sprig functions are supported at the moment. This list would cover most of the Sprig functions in the future.

- `upper`
- `snakecase`
- `camelcase`
- `kebabcase`
- `decryptAES`

## Usage

The library can you used to extend jsonnet functionality offered by your application.
To install run:

```bash
go get -u github.com/harsimranmaan/go-jsonnet-func
```

It can used in a sample program as shown in the [example](examples/main.go).
