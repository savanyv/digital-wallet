# Variables
PROTO_DIR := proto
GO_OUT_DIR := .
PROTOC := protoc
PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc
PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto")

.PHONY: all
all: check-deps generate

.PHONY: check-deps
check-deps:
	@which $(PROTOC) > /dev/null || (echo "Error: protoc not installed. Please install protobuf compiler"; exit 1)
	@which $(PROTOC_GEN_GO) > /dev/null || (echo "Error: protoc-gen-go not installed. Run: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"; exit 1)
	@which $(PROTOC_GEN_GO_GRPC) > /dev/null || (echo "Error: protoc-gen-go-grpc not installed. Run: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"; exit 1)

.PHONY: generate
generate: check-deps
	@echo "Generating gRPC code..."
	@for proto in $(PROTO_FILES); do \
		echo "Generating for $$proto..."; \
		$(PROTOC) \
			--go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
			--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
			-I$(PROTO_DIR) $$proto; \
	done
	@echo "Generation completed successfully"

.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	@find $(GO_OUT_DIR) -name "*.pb.go" -type f -delete
	@echo "Clean completed"
