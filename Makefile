# Go Daily Code 测试 Makefile

PACKAGES := ./problems/... ./datastruct/... ./patterns/...
CODING_INTERVIEWS := ./problems/coding_interviews/...

.PHONY: help test test-verbose test-coverage test-bench test-all test-coding-interviews test-coding-interviews-bench test-coding-interviews-coverage test-coding-interviews-coverage-html test-coding-interviews-specific test-coding-interviews-bench-mem clean fmt vet check deps

help:
	@echo "=== Go Daily Code 测试命令 ==="
	@echo "  make test                         - 运行全部稳定目录测试"
	@echo "  make test-coding-interviews       - 仅运行 Interview 75 / LCR 测试"
	@echo "  make test-coverage                - 生成全部稳定目录覆盖率报告"
	@echo "  make test-bench                   - 运行全部 benchmark"
	@echo "  make fmt                          - 格式化全部稳定目录代码"
	@echo "  make vet                          - 检查全部稳定目录代码"
	@echo "  make check                        - 依次运行 fmt、vet 和 test"

test:
	go test $(PACKAGES)

test-verbose:
	go test -v $(PACKAGES)

test-coverage:
	go test -coverprofile=coverage.out $(PACKAGES)
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

test-bench:
	go test -bench=. $(PACKAGES)

test-all:
	go test -v -bench=. $(PACKAGES)

test-coding-interviews:
	go test $(CODING_INTERVIEWS)

test-coding-interviews-bench:
	go test -bench=. $(CODING_INTERVIEWS)

test-coding-interviews-coverage:
	go test -cover $(CODING_INTERVIEWS)

test-coding-interviews-coverage-html:
	go test -coverprofile=coverage.out $(CODING_INTERVIEWS)
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

test-coding-interviews-specific:
	@if [ -z "$(TEST)" ]; then \
		echo "请指定测试名称，例如: make test-coding-interviews-specific TEST=TestDecorateRecord"; \
		exit 1; \
	fi
	go test -v -run='$(TEST)' $(CODING_INTERVIEWS)

test-coding-interviews-bench-mem:
	go test -bench=. -benchmem $(CODING_INTERVIEWS)

clean:
	go clean -testcache
	rm -f coverage.out coverage.html

deps:
	go mod tidy
	go mod download

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

check: fmt vet test
	@echo "所有检查完成"
