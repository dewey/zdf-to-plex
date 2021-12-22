#!/bin/bash
set -e

source develop.env

function cleanup() {
    rm -f zdf-to-plex
}
trap cleanup EXIT

GO111MODULE=on GOGC=off go build -mod=vendor -v -o zdf-to-plex
./zdf-to-plex