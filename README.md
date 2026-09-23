# go-base64

A Base64 encoder/decoder written in Go for learning purposes.

Go already provides an [encoding/base64](https://pkg.go.dev/encoding/base64) package, but implementing it from scratch is a good exercise for working with bitwise operations and the Go standard library.

## Usage

    base64enc [-d] [-h] [filename|--]

    -d         decode
    -h         print help
    --         read from stdin
    filename   file to encode or decode, falls back to treating the argument as a literal string

## Build

    go build -o enc .

## Examples

    ./enc go.mod
    ./enc -d -- <<< "aGVsbG8="
