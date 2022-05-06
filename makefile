build:
	go build -o im
start:
	go build -o im && ./im
test:
	go test -v ./... -race -covermode=atomic