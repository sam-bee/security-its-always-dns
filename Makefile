PHONY: build-dnsexfiltool build-dnsreceiver

build-dnsexfiltool:
	echo "Building DNS exfil tool..."
	if [ ! -f ./cmd/dnsexfiltool/config.ini ]; then echo "ERROR: no ini file found at ./cmd/dnsexfiltool/config.ini"; exit 1; fi
	go build -o bin/det cmd/dnsexfiltool/main.go

build-dnsreceiver:
	echo "Building DNS receiver..."
	go build -o bin/dnsreceiver cmd/dnsreceiver/main.go

test:
	go fmt ./...
	go build ./...
	go test -v ./...
