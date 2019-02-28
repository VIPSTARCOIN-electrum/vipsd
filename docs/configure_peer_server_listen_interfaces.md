vipsd allows you to bind to specific interfaces which enables you to setup
configurations with varying levels of complexity.  The listen parameter can be
specified on the command line as shown below with the -- prefix or in the
configuration file without the -- prefix (as can all long command line options).
The configuration file takes one entry per line.

**NOTE:** The listen flag can be specified multiple times to listen on multiple
interfaces as a couple of the examples below illustrate.

Command Line Examples:

|Flags|Comment|
|----------|------------|
|--listen=|all interfaces on default port which is changed by `--testnet` and `--regtest` (**default**)|
|--listen=0.0.0.0|all IPv4 interfaces on default port which is changed by `--testnet` and `--regtest`|
|--listen=::|all IPv6 interfaces on default port which is changed by `--testnet` and `--regtest`|
|--listen=:31915|all interfaces on port 31915|
|--listen=0.0.0.0:31915|all IPv4 interfaces on port 31915|
|--listen=[::]:31915|all IPv6 interfaces on port 31915|
|--listen=127.0.0.1:31915|only IPv4 localhost on port 31915|
|--listen=[::1]:31915|only IPv6 localhost on port 31915|
|--listen=:8336|all interfaces on non-standard port 8336|
|--listen=0.0.0.0:8336|all IPv4 interfaces on non-standard port 8336|
|--listen=[::]:8336|all IPv6 interfaces on non-standard port 8336|
|--listen=127.0.0.1:8337 --listen=[::1]:31915|IPv4 localhost on port 8337 and IPv6 localhost on port 31915|
|--listen=:31915 --listen=:8337|all interfaces on ports 31915 and 8337|

The following config file would configure vipsd to only listen on localhost for both IPv4 and IPv6:

```text
[Application Options]

listen=127.0.0.1:31915
listen=[::1]:31915
```
