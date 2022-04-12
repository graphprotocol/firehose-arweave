# firehose-arweave

Firehose integration for Arweave.

`firehose-arweave` use [thegarii](https://github.com/ChainSafe/thegarii) as the source
of Arweave blocks.


## Quick Start

> NOTE
>
> this script will install rust and thegarii automatically ( if not exist ).
```
# Arguments shared with both thegarii:
#
# --mindreader-node-start-block-num
# --mindreader-node-stop-block-num
./devel/standard/start.sh -- --mindreader-node-start-block-num 911988
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


## License

[Apache 2.0](LICENSE)
