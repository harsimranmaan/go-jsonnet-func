.PHONY: test
test:
	go test -coverprofile=coverage.txt -covermode=atomic -race ./...
.PHONY: get
get:
	go get ./...

.PHONY: gen
gen:
	go generate ./...
