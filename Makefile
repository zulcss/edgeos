clean:
	rm -rf bin
agent:
	go build -o bin/agent cmd/agent/main.go
