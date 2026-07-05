# ========================================
# TEMPLATE BEGIN: TEST AND COVERAGE CONFIG
# ========================================

COVERAGE_DIR ?= .coverage.out

test:
	@if [ -d $(COVERAGE_DIR) ]; then rm -r $(COVERAGE_DIR); fi
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...

# ========================================
# TEMPLATE END: TEST AND COVERAGE CONFIG
# ========================================

# ========================================
# PROTOBUF CODE GENERATION
# ========================================

# Install buf (protobuf code generation)
# 安装 buf 工具
.PHONY: init
init:
	go install github.com/bufbuild/buf/cmd/buf@latest
	@echo "buf 安装完成!"

# Generate Go code from proto files
# 从 proto 文件生成 Go 代码
.PHONY: generate
generate:
	buf generate
	@echo "protobuf 代码生成完成!"

# Remove generated .pb.go files
# 清理生成的 .pb.go 文件
.PHONY: clean
clean:
	rm -f protos/protoenumstatus/*.pb.go
	rm -f protos/protoenumresult/*.pb.go
	@echo "清理生成文件完成!"

# Show available targets
# 显示可用的命令目标
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  test     - Run tests with coverage"
	@echo "  init     - Install buf (protobuf code generation)"
	@echo "  generate - Generate Go code from proto files"
	@echo "  clean    - Remove generated .pb.go files"
	@echo "  help     - Show this help message"
