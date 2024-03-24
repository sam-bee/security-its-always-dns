# It's Always DNS...

### _A DNS exfil tool, for smuggling data from a compromised machine using outgoing DNS lookups, with data hidden in the subdomains._

**This project is still under construction, and is not yet feature complete**

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
                                               |'----'|
                                               |      |
                                               |  DB  |
                                               (      )
                                                '----'
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

To build the DNS Receiver:
- `cp ./cmd/dnsreceiver/config/.env.example ./cmd/dnsreceiver/config/.env`
- Manually edit `./cmd/dnsreceiver/config/.env` to your needs
- `make build-dnsreceiver`

To build the DNS exfil tool:
- `cp ./cmd/dnsexfiltool/config/config.ini.example ./cmd/dnsexfiltool/config/config.ini`
- Manually edit `./cmd/dnsexfiltool/config/config.ini` to your needs
- `make build-dnsexfiltool`

Your binaries will end up in `./bin/`.

## Configuration

The DNS Receiver is dependent on a `.env` file like the one in `./cmd/dnsreceiver/.env.example`. You may override
these settings with environment variables on the server, which take precedence.

For the DNS Exfil Tool, however, you will need to set the `./cmd/dnsexfiltool/config.ini` file, which will be
included in the binary at compile time. You do not need the ability to copy your config file onto the target
machine. You do not need the ability to set environment variables on the target machine. You only need to get the
binary itself onto the target machine, and cause it to execute. For this convenience, the trade-off is that
recompiling the binary is the only way to change its configuration.

## dns-receiver binary

This is the main command for running the nameserver/listener that will receive DNS lookups. It is here that the
exfiltrated data will be received.

After building as per the `README.md` file in the project root, you can test the nameserver by doing the following:

```
sudo ./bin/dnsreceiver --config "./cmd/dnsreceiver/.env"
```

Your binary must be run with sudo as it needs to listen on port 53. In another tab,

```
dig @127.0.0.1 example.com
```

You should see the correct IP address returned to you. 

## Development

You can run the tests with this command:

```
make test
```
