#!/usr/bin/env zx

await $`protoc \
  --proto_path=./test_protos \
  --plugin=protoc-gen-ts=/usr/local/bin/protoc-gen-ts_proto \
  --ts_out=./test_protos \
  --plugin=protoc-gen-http=./main \
  --http_out=./test_protos \
  --http_opt=nameCase=pascal \
  ./test_protos/test.proto
  `