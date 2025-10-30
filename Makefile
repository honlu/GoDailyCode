# Go Daily Code 测试 Makefile
# 当前仅支持 coding-interviews 包测试

.PHONY: help test test-verbose test-coverage test-bench test-all test-coding-interviews test-coding-interviews-bench test-coding-interviews-coverage test-coding-interviews-coverage-html test-coding-interviews-specific test-coding-interviews-bench-mem run-tests clean fmt vet check deps

# 默认目标
help:
	@echo "=== Go Daily Code 测试命令 ==="
	@echo ""
	@echo "基本测试命令:"
	@echo "  make test-coding-interviews      - 运行 coding-interviews 包测试（推荐）"
	@echo "  make test-all                    - 运行所有测试和基准测试"
	@echo "  make run-tests                   - 运行交互式测试程序（推荐）"
	@echo ""
	@echo "详细测试命令:"
	@echo "  make test-coding-interviews-bench - 运行 coding-interviews 基准测试"
	@echo "  make test-coding-interviews-coverage - 运行 coding-interviews 覆盖率测试"
	@echo "  make test-coding-interviews-coverage-html - 生成HTML覆盖率报告"
	@echo "  make test-coding-interviews-bench-mem - 基准测试（显示内存分配）"
	@echo "  make test-coding-interviews-specific TEST=TestName - 运行特定测试"
	@echo ""
	@echo "代码质量命令:"
	@echo "  make fmt                         - 格式化代码"
	@echo "  make vet                         - 运行代码检查"
	@echo "  make check                       - 运行所有检查"
	@echo ""
	@echo "其他命令:"
	@echo "  make clean                       - 清理测试缓存"
	@echo "  make deps                        - 安装依赖"

# 运行 coding-interviews 包测试（默认）
test:
	@echo "运行 coding-interviews 包测试..."
	go test -v ./coding-interviews

# 运行详细测试
test-verbose:
	@echo "运行详细测试..."
	go test -v ./coding-interviews

# 运行测试覆盖率
test-coverage:
	@echo "运行测试覆盖率..."
	go test -cover ./coding-interviews
	@echo "生成详细覆盖率报告..."
	go test -coverprofile=coverage.out ./coding-interviews
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

# 运行基准测试
test-bench:
	@echo "运行基准测试..."
	go test -bench=. ./coding-interviews

# 运行所有测试和基准测试
test-all:
	@echo "运行所有测试和基准测试..."
	go test -v -bench=. ./coding-interviews

# 只运行 coding-interviews 包测试
test-coding-interviews:
	@echo "运行 coding-interviews 包测试..."
	go test -v ./coding-interviews

# 运行 coding-interviews 基准测试
test-coding-interviews-bench:
	@echo "运行 coding-interviews 基准测试..."
	go test -bench=. ./coding-interviews

# 运行 coding-interviews 覆盖率测试
test-coding-interviews-coverage:
	@echo "运行 coding-interviews 覆盖率测试..."
	go test -cover ./coding-interviews

# 运行 coding-interviews 覆盖率测试并生成报告
test-coding-interviews-coverage-html:
	@echo "运行 coding-interviews 覆盖率测试并生成HTML报告..."
	go test -coverprofile=coverage.out ./coding-interviews
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

# 运行 coding-interviews 特定测试
test-coding-interviews-specific:
	@echo "运行 coding-interviews 特定测试..."
	@echo "可用测试:"
	@echo "  make test-specific TEST=TestDecorateRecord"
	@echo "  make test-specific TEST=ExampleDecorateRecord"
	@echo "  make test-specific TEST=BenchmarkDecorateRecord"
	@if [ -z "$(TEST)" ]; then \
		echo "请指定测试名称，例如: make test-specific TEST=TestDecorateRecord"; \
		exit 1; \
	fi
	go test -v -run=$(TEST) ./coding-interviews

# 运行 coding-interviews 基准测试并显示内存分配
test-coding-interviews-bench-mem:
	@echo "运行 coding-interviews 基准测试（显示内存分配）..."
	go test -bench=. -benchmem ./coding-interviews

# 运行交互式测试程序
run-tests:
	@echo "运行交互式测试程序..."
	cd test-runner && go run main.go

# 清理测试缓存
clean:
	@echo "清理测试缓存..."
	go clean -testcache
	rm -f coverage.out coverage.html

# 安装依赖
deps:
	@echo "安装依赖..."
	go mod tidy
	go mod download

# 格式化代码
fmt:
	@echo "格式化代码..."
	go fmt ./coding-interviews

# 运行代码检查
vet:
	@echo "运行代码检查..."
	go vet ./coding-interviews || echo "注意: 发现一些警告，但代码可以正常运行"

# 运行所有检查（格式化、检查、测试）
check: fmt vet test-coding-interviews
	@echo "所有检查完成"
