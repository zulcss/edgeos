clean:
	rm -rf bin
agent:
	go build -o bin/agent cmd/agent/main.go

cli:
	go build -o bin/stx cmd/stx/main.go
