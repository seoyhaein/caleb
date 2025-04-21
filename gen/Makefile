# Path to the protoc compiler
PROTOC_PATH=/usr/local/bin/protoc

# Path to the Protocol Buffers folder
PROTO_PATH=../protos

# Directory to save the generated files
OUT_DIR=$(PROTO_PATH)

# Path to the protoc-gen-go plugin
PROTOC_GEN_GO_PATH=/usr/local/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC_PATH=/usr/local/bin/protoc-gen-go-grpc

# Find all .proto files in the PROTO_PATH
PROTO_FILES=$(wildcard $(PROTO_PATH)/*.proto)

# Default target
all: generate

# Generate Protocol Buffers code
generate: $(PROTO_FILES)
	@for file in $^; do \
		$(PROTOC_PATH) --proto_path=$(PROTO_PATH) --plugin=$(PROTOC_GEN_GO_PATH) --go_out=$(OUT_DIR) --go_opt=paths=source_relative $$file; \
		$(PROTOC_PATH) --proto_path=$(PROTO_PATH) --plugin=$(PROTOC_GEN_GO_GRPC_PATH) --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative $$file; \
	done

# Clean generated files
clean:
	rm -f $(OUT_DIR)/*.pb.go

.PHONY: all generate clean
