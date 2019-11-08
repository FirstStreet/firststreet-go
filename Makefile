lint:
	LOG_LEVEL=error golangci-lint run --tests=false --skip-dirs /*mock -v

test:
	go test ./...

test-v:
	go test ./... -v

coverage:
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic
