#!/usr/bin/env zx

await $`protoc \
  --proto_path=./ \
  --plugin=protoc-gen-http=./main \
  --plugin=protoc-gen-ts=/usr/local/bin/protoc-gen-ts_proto \
  --http_out=./ \
  --ts_out=./ \
  ./test_protos/*.proto
  `