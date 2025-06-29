
bench:
	go test -bench .

test:
	go test .

lint:
	go mod tidy
	gofmt -w -s *.go
	golangci-lint run .

period:
	go build cmd/period/main.go

