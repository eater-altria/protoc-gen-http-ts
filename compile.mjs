#!/usr/bin/env zx

await $`CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build main.go`
await $`mv main release/protoc-gen-http-1.0.0-macos-arm64`

await $`CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go`
await $`mv main release/protoc-gen-http-1.0.0-macos-amd64`

await $`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go`
await $`mv main release/protoc-gen-http-1.0.0-linux-amd64`

await $`CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go`
await $`mv main.exe release/protoc-gen-http-1.0.0-windows-amd64.exe`