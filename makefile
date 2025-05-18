PROTO_DIR := proto
PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto")

.PHONY: all
all: generate

.PHONY: generate
generate:
	@echo "Generating Go code from .proto files..."
	@for file in $(PROTO_FILES); do \
		protoc \
			--proto_path=$(PROTO_DIR) \
			--go_out=$(PROTO_DIR) \
			--go_opt=paths=source_relative \
			--go-grpc_out=$(PROTO_DIR) \
			--go-grpc_opt=paths=source_relative \
			$$file; \
	done
	@echo "Done."

.PHONY: clean
clean:
	@find $(GO_OUT_DIR) -name "*.pb.go" -delete
