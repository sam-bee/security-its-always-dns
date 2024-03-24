## Codec package

### Overview

For encoding and decoding data during DNS exfiltration.
- Encoding happens on the target machine, where data is broken down into 120 byte chunks and hidden in a subdomains.
- Data should then be exfiltrated from the network, hidden in DNS lookups.
- Decoding happens on the attacker's nameserver, which receives the lookup.

The scope of this package is just the encoding and decoding: `data > domain names > data`.

### Example encoding

```go
exfilPacketDomains := GetDomainsToLookUp(stringToExfil, "example.com")
// exfilPacketDomains[0] will look something like [exfil-id]-0.[payload-data].[payload-data].[payload-data].example.com
```

Each fully qualified domain name in the set will contain up to 120 bytes of exfil data. It is converted to base 36 (`[a-z0-9]`), but the data passed into this package is not otherwise obfuscated or compressed.

### Example decoding

```go
// fqdns is a []string of fully qualified domain names.
// "example.com" here is the base domain to strip off the end before decoding them.
decodedData, err := DecodeDataFromFqdns(fqdns, "example.com")
```
