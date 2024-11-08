MAIN=cmd/main.go
BIN=cmd/bin/
OUT=nri-gcp-l4-proxy

all: macos linux-arm linux-intel windows

macos:
	GOOS=darwin GOARCH=arm64 go build -o $(BIN)/macos-arch64/$(OUT) $(MAIN)
	
linux-arm:
	GOOS=linux GOARCH=arm64 go build -o $(BIN)/linux-arch64/$(OUT) $(MAIN)

linux-intel:
	GOOS=linux GOARCH=amd64 go build -o $(BIN)/linux-amd64/$(OUT) $(MAIN)

windows:
	GOOS=windows GOARCH=amd64 go build -o $(BIN)/windows-amd64/$(OUT) $(MAIN)

clean:
	rm -rf $(BIN)
