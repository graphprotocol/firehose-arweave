# Firehose on Dummy Blockchain
[![reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/streamingfast/firehose-acme)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# Usage

> NOTE
>
> this script will install rust and thegarii automatically.

```
./devel/standard/start.sh
```

# Environments of thegarii
    
| KEY          | DEFAULT_VALUE          | DESCRIPTION                                 |
|--------------|------------------------|---------------------------------------------|
| ENDPOINTS    | "https://arweave.net"  | for multiple endpoints, split them with ',' |
| BATCH_BLOCKS | 50                     | how many blocks batch at one time           |
| CONFIRMS     | 20                     | irreversibility condition                   |
| PTR_PATH     | $APP_DATA/thegarii/ptr | the file stores the block ptr for polling   |
| retry        | 10                     | retry times when failed on http requests    |
| timeout      | 120_000                | timeout of http requests                    |


for configuring these variables with arguments, see `thegarii -h`

```
thegaril 0.0.2
info@chainsafe.io
env arguments for CLI

USAGE:
    thegarii [FLAGS] [OPTIONS]

FLAGS:
    -d, --debug      Activate debug mode
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -B, --batch-blocks <batch-blocks>    how many blocks polling at one time
    -b, --block-time <block-time>        time cost for producing a new block in arweave
    -c, --confirms <confirms>            safe blocks against to reorg in polling
    -D, --db-path <db-path>              storage db path ( only works with full features )
    -e, --endpoints <endpoints>...       client endpoints
    -g, --grpc-addr <grpc-addr>          grpc address
    -p, --ptr-path <ptr-path>            block ptr file path
    -r, --retry <retry>                  retry times when failed on http requests
    -t, --timeout <timeout>              timeout of http requests

```


## Release

Use the `./bin/release.sh` Bash script to perform a new release. It will ask you questions
as well as driving all the required commands, performing the necessary operation automatically.
The Bash script runs in dry-mode by default, so you can check first that everything is all right.

Releases are performed using [goreleaser](https://goreleaser.com/).

## Contributing

**Issues and PR in this repo related strictly to the Firehose on Dummy Blockchain.**

Report any protocol-specific issues in their
[respective repositories](https://github.com/streamingfast/streamingfast#protocols)

**Please first refer to the general
[StreamingFast contribution guide](https://github.com/streamingfast/streamingfast/blob/master/CONTRIBUTING.md)**,
if you wish to contribute to this code base.

This codebase uses unit tests extensively, please write and run tests.

## License

[Apache 2.0](LICENSE)
