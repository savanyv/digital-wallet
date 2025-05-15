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
	@which $(PROTOC) > /dev/null || (echo "protoc not found"; exit 1)
	@which $(PROTOC_GEN_GO) > /dev/null || (echo "protoc-gen-go not found"; exit 1)
	@which $(PROTOC_GEN_GO_GRPC) > /dev/null || (echo "protoc-gen-go-grpc not found"; exit 1)

.PHONY: generate
generate: check-deps
	@echo "Generating code..."
	@for proto in $(PROTO_FILES); do \
		echo "Generating for $$proto"; \
		$(PROTOC) \
			-I=$(PROTO_DIR) \
			--go_out=$(GO_OUT_DIR) --go_opt=paths=source_relative \
			--go-grpc_out=$(GO_OUT_DIR) --go-grpc_opt=paths=source_relative \
			$$proto; \
	done
	@echo "Done."

.PHONY: clean
clean:
	@find $(GO_OUT_DIR) -name "*.pb.go" -delete
