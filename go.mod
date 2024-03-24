module github.com/sam-bee/security-itsalwaysdns/dns_receiver

go 1.22

require (
	github.com/sam-bee/security-itsalwaysdns/dns_exfil_tool v1.0.0
	github.com/sam-bee/security-itsalwaysdns/integration_tests v1.0.0
	github.com/sam-bee/security-itsalwaysdns/shared v1.0.0
)

replace (
	github.com/sam-bee/security-itsalwaysdns/dns_exfil_tool v1.0.0 => ../dns_exfil_tool
	github.com/sam-bee/security-itsalwaysdns/dns_receiver v1.0.0 => ../dns_receiver
	github.com/sam-bee/security-itsalwaysdns/integration_tests v1.0.0 => ../integration_tests
	github.com/sam-bee/security-itsalwaysdns/shared v1.0.0 => ../shared
)
