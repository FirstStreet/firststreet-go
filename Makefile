test:
	go test ./...

test-v:
	go test ./... -v

coverage:
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic
