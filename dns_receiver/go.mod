module github.com/sam-bee/security-itsalwaysdns/dns_receiver

go 1.22

require (
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/miekg/dns v1.1.58
	github.com/sam-bee/security-itsalwaysdns/shared v1.0.0
)

require (
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
)

replace github.com/sam-bee/security-itsalwaysdns/shared v1.0.0 => ../shared
