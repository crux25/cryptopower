# cryptopower

A cross-platform desktop wallet built with [gio](https://gioui.org/).

## Building

Note: You need to have [Go 1.19](https://golang.org/dl/) or above to build.

Then `go build`.

### Linux

To build **cryptopower** on Linux these [gio dependencies](https://gioui.org/doc/install/linux) are required.

Arch Linux:
`pacman -S vulkan-headers libxkbcommon-x11`

## FreeBSD

To build **cryptopower** on FreeBSD you will need to `pkg install vulkan-headers` as root. This is a gio dependency.

## Running cryptopower

### General usage

By default, **cryptopower** runs on Mainnet network type. However, cryptopower can run on testnet by issuing commands on the terminal in the format:

```bash
cryptopower [options]
```

- Run `./cryptopower --network=testnet` to run cryptopower on the testnet network.
- Run `cryptopower -h` or `cryptopower help` to get general information of commands and options that can be issued on the cli.
- Use `cryptopower <command> -h` or `cryptopower help <command>` to get detailed information about a command.

## Profiling

Cryptopower uses [pprof](https://github.com/google/pprof) for profiling. It creates a web server which you can use to save your profiles. To setup a profiling web server, run cryptopower with the --profile flag and pass a server port to it as an argument.

So, after running the build command above, run the command

`./cryptopower --profile=6060`

You should now have a local web server running on 127.0.0.1:6060.

To save a profile, you can simply use

`curl -O localhost:6060/debug/pprof/profile`

## Contributing

See [CONTRIBUTING.md](https://github.com/crypto-power/cryptopower/blob/master/.gitlab/ci/CONTRIBUTING.md)

## Other

Earlier experimental work with other user interface toolkits can be found at [godcr-old](https://github.com/raedahgroup/godcr-old).

## Bootstrappable Cryptopower Builds

The [reproduciblebuilds](https://github.com/crypto-power/cryptopower/tree/master/reproduciblebuilds) directory contains the files necessary to perform bootstrappable Cryptopower builds.

Bootstrappability furthers our binary security guarantees by allowing us to audit and reproduce our toolchain instead of blindly trusting binary downloads.

We achieve bootstrappability by using Docker.