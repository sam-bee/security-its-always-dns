# It's Always DNS...

### _A DNS exfil tool, for smuggling data from a compromised machine using outgoing DNS lookups, with data hidden in the subdomains._

## Overview

DNS exfiltration is where data is sent from a compromised machine to a nameserver controlled by the attacker. The data is
hidden in subdomains, and sent out via DNS lookups. This may evade detection where DNS lookups are not monitored.

```text
COMPROMISED                                   ATTACKER'S 
  MACHINE                                     NAMESERVER
                   DNS lookup: where is
 ___________   "[hidden-data].example.com"?   __________
| DNS Exfil | -----------------------------> | DNS      |
| Tool      |                                | Receiver |
|___________| <----------------------------- |__________|
                        IP address                 |
                                       exfiltrated |
                                       hidden data |
                                                   v
                                                .----.
                                               (      )
                                               |^----^|
                                               |      |
                                               |  DB  |
                                               (      )
                                                ^----^
```

This project consists of two parts:
1. **DNS Exfil Tool** - a small binary for turning data into fully qualified domain names (FQDN's), and sending DNS
lookups
2. **DNS Receiver** - a tool that pretends to be a nameserver, receives exfiltrated data, and decodes/reassembles it.

## System Setup

You will need a cloud server to install the DNS Receiver. Start running the nameserver. Register a domain name for your
nameserver (such as `nameserver.com`), then register a domain name for exfiltration (such as `example.com`). You will
need to tell the registrar that `nameserver.com` is the nameserver for `*.example.com` domains. The nameserver will
then receive DNS lookups from any machine asking for `[hidden-data].example.com`. Your `dns_receiver` binary will
receive these requests, and log the hidden data. The `det` binary sends them, and should be installed on the target
machine.

## Building

To build the DNS Receiver, type: `make build-dnsreceiver`

To build the DNS exfil tool:
- `cp ./cmd/dnsexfiltool/config.ini.example ./cmd/dnsexfiltool/config.ini`
- Manually edit `./cmd/dnsexfiltool/config.ini` to your needs
- `make build-dnsexfiltool`

Your binaries will end up in `./bin/`.

## Development

You can run the tests with this command:

```
make test
```
