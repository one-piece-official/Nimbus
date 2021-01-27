TAG=$(shell git rev-parse --short HEAD)


testing:
	go test -v -count 1 ./...
	golangci-lint run -v
