BIN := tbc
PREFIX := /usr/local/bin

$(BIN): $(wildcard *.go) $(wildcard */*.go)
	CGO_ENABLED=0 go build -ldflags='-s -d -buildid=' .

install: $(BIN)
	install -m 755 -D -t $(PREFIX) $(BIN)

clean:
	rm -rf $(BIN)
	go clean
