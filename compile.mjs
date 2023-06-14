#!/usr/bin/env zx

const version = "1.0.4"

await $`CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build main.go`
await $`mv main release/protoc-gen-http-ts-${version}-macos-arm64`

await $`CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go`
await $`mv main release/protoc-gen-http-ts-${version}-macos-amd64`

await $`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go`
await $`mv main release/protoc-gen-http-ts-${version}-linux-amd64`

await $`CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go`
await $`mv main.exe release/protoc-gen-http-ts-${version}-windows-amd64.exe`