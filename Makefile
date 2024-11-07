all:
	GOOS=darwin GOARCH=arm64 go build -o cmd/bin/macos-arch64/nri-gcp-l4-proxy cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o cmd/bin/linux-arch64/nri-gcp-l4-proxy cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o cmd/bin/linux-amd64/nri-gcp-l4-proxy cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o cmd/bin/windows-amd64/nri-gcp-l4-proxy cmd/main.go
clean:
	rm -rf cmd/bin/
