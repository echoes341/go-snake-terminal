# go-snake-terminal
Porting of [Plutov's snake telnet game](https://github.com/plutov/go-snake-telnet) to standard terminal

# Install it and run it
To build and run this package you'll need [dep](https://github.com/tools/godep)

	go get -d github.com/echoes341/go-snake-terminal
	cd `go env GOPATH`/src/github.com/echoes341/go-snake-terminal
	dep ensure --vendor-only
	go build -o snake
	./snake