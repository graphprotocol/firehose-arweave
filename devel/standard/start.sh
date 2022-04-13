#!/usr/bin/env bash

ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

clean=
sfarweave="$ROOT/../firearweave"

main() {
  pushd "$ROOT" &> /dev/null

  while getopts "hc" opt; do
    case $opt in
      h) usage && exit 0;;
      c) clean=true;;
      \?) usage_error "Invalid option: -$OPTARG";;
    esac
  done
  shift $((OPTIND-1))
  [[ $1 = "--" ]] && shift

  set -e

  if [[ $clean == "true" ]]; then
    rm -rf fire-data &> /dev/null || true
  fi

  # check if thegarii exists
  check_thegarii

  exec $sfarweave -c $(basename $ROOT).yaml start "$@"
}

usage_error() {
  message="$1"
  exit_code="$2"

  echo "ERROR: $message"
  echo ""
  usage
  exit ${exit_code:-1}
}

usage() {
  echo "usage: start.sh [-c]"
  echo ""
  echo "Start $(basename $ROOT) environment."
  echo ""
  echo "Options"
  echo "    -c             Clean actual data directory first"
}

install_rust() {
    if ! command -v cargo &> /dev/null
    then
        echo "rust toolchain not exists, installing rust..."
        curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
        source $HOME/.cargo/env
    fi
}


check_thegarii() {
    if ! command -v thegarii &> /dev/null
    then
        install_rust
        echo "thegarii has not been installed, installing thegarii..."
        cargo install thegarii
    fi
}

main "$@"
