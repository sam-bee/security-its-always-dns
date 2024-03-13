## dns-exfil-nameserver binary

### Overview

This is the main command for running the nameserver/listener that will receive DNS lookups. It is here that the
exfiltrated data will be received.

After building as per the `README.md` file in the project root, you can test the nameserver by doing the following:

```
sudo ./bin/dnsreceiver
```

and, in another tab,

```
dig @127.0.0.1 example.com
```

You should see the correct IP address returned to you. Your binary must be run with sudo as it needs to listen on port 53.
