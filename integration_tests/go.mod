module github.com/sam-bee/security-itsalwaysdns/integration_tests

go 1.22

require (
    github.com/sam-bee/security-itsalwaysdns/dns_exfil_tool v1.0.0
)

replace github.com/sam-bee/security-itsalwaysdns/dns_exfil_tool => ../dns_exfil_tool
