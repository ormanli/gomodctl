.PHONY: build
build:
	go build -o gomodctl -v main.go

.PHONY: test
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -tags=integration ./...
