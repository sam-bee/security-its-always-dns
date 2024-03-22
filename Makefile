.PHONY: build-dnsexfiltool build-dnsreceiver build test

build-dnsexfiltool:
	echo "Building DNS exfil tool..."
	if [ ! -f "./cmd/dnsexfiltool/config.toml" ]; then exit 1; fi
	go build -o bin/det cmd/dnsexfiltool/main.go
	echo "Built to ./bin/det"

build-dnsreceiver:
	echo "Building DNS receiver..."
	go build -o bin/dnsreceiver cmd/dnsreceiver/main.go
	echo "Built to ./bin/dnsreceiver"

build: build-dnsexfiltool build-dnsreceiver

test:
	go fmt ./...
	go build ./...
	go test -v ./...
