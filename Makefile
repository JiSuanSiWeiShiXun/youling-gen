.PHONY: build install clean test run help

BINARY_NAME=youling-gen
BUILD_DIR=bin
MAIN_FILE=main.go

# 默认目标
all: build

# 构建
build:
	@echo "构建 $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "✓ 构建完成: $(BUILD_DIR)/$(BINARY_NAME)"

# 安装到 GOPATH/bin
install:
	@echo "安装 $(BINARY_NAME)..."
	@go install
	@echo "✓ 安装完成"

# 清理
clean:
	@echo "清理构建文件..."
	@rm -rf $(BUILD_DIR)
	@echo "✓ 清理完成"

# 运行测试
test:
	@echo "运行测试..."
	@go test -v ./...

# 运行程序
run:
	@go run $(MAIN_FILE)

# 格式化代码
fmt:
	@echo "格式化代码..."
	@go fmt ./...
	@echo "✓ 格式化完成"

# 代码检查
lint:
	@echo "代码检查..."
	@golangci-lint run ./...

# 下载依赖
deps:
	@echo "下载依赖..."
	@go mod download
	@go mod tidy
	@echo "✓ 依赖下载完成"

# 帮助信息
help:
	@echo "可用命令:"
	@echo "  make build    - 构建项目"
	@echo "  make install  - 安装到 GOPATH/bin"
	@echo "  make clean    - 清理构建文件"
	@echo "  make test     - 运行测试"
	@echo "  make run      - 运行程序"
	@echo "  make fmt      - 格式化代码"
	@echo "  make lint     - 代码检查"
	@echo "  make deps     - 下载依赖"
	@echo "  make help     - 显示此帮助信息"
