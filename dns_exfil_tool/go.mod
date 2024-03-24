module github.com/sam-bee/security-itsalwaysdns/dns_exfil_tool

go 1.22

require (
	github.com/pelletier/go-toml v1.9.5
	github.com/sam-bee/security-itsalwaysdns/shared v1.0.0
)

replace github.com/sam-bee/security-itsalwaysdns/shared v1.0.0 => ../shared
