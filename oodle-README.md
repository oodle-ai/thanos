## How to install protobuf

1. brew install protobuf
1. source .bingo/variables.env
1. go get github.com/golang/protobuf@v1.3.2
1. go install github.com/bwplotka/bingo@latest
1. bingo get protoc-gen-go@v1.31.0
1. bingo get protoc-gen-go-vtproto@v0.6.0

## How to update proto files

1. Create a seperate workspace for [oodle-thanos](https://github.com/oodle-ai/oodle-thanos) on Mac
2. Make changes to a proto file
3. Run `make proto` on Mac (not inside dev container)
