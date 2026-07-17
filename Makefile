LOCAL_BIN := $(CURDIR)/bin

install:
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@v1.29.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v3@v3.7.1

protogen:
	$(LOCAL_BIN)/buf generate proto --path proto/items

mocks:
	$(LOCAL_BIN)/mockery

tests:
	go test ./...