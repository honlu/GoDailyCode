# Go Daily Code 测试指南

这个项目包含了完整的测试框架，可以本地运行各种测试。
暂时只支持coding-interviews包的测试。

## 🚀 快速开始

### 方法1: 使用 Makefile

```bash
# 查看所有可用命令
make help

# 运行 coding-interviews 包测试
make test-coding-interviews

# 运行 coding-interviews 基准测试
make test-coding-interviews-bench

# 运行 coding-interviews 覆盖率测试
make test-coding-interviews-coverage

# 运行所有测试（仅 coding-interviews 包）
make test-all
```

### 方法2: 使用交互式测试程序(推荐)

```bash
# 运行交互式测试程序
make run-tests
# 或者
cd test-runner && go run main.go
```

### 方法3: 直接使用 go test 命令

```bash
# 运行 coding-interviews 包测试
go test -v ./coding-interviews

# 运行基准测试
go test -bench=. ./coding-interviews

# 查看测试覆盖率
go test -cover ./coding-interviews

# 生成覆盖率报告
go test -coverprofile=coverage.out ./coding-interviews
go tool cover -html=coverage.out -o coverage.html

# 运行所有测试和基准测试
go test -v -bench=. ./coding-interviews
```

## 📁 测试文件结构

```
coding-interviews/
├── 75_120.go              # 主要算法实现
├── common.go              # 通用数据结构定义
└── decorate_record_test.go # 单元测试文件

test-runner/
└── main.go                # 交互式测试运行器
```

## 📋 当前支持的测试范围

**注意**: 当前测试框架仅支持 `coding-interviews` 包中的算法测试。

### 支持的算法
- `decorateRecord` - 二叉树的层序遍历算法

### 测试文件
- `coding-interviews/decorate_record_test.go` - 包含完整的测试用例

## 🧪 测试类型

### 1. 单元测试 (Unit Tests)
- 测试 `decorateRecord` 函数的各种场景
- 包括空树、单节点、完全二叉树、斜树等测试用例
- 使用表驱动测试模式

### 2. 基准测试 (Benchmark Tests)
- 测试算法性能
- 使用100个节点的树进行压力测试

### 3. 示例测试 (Example Tests)
- 提供使用示例
- 验证输出格式

## 🔧 测试辅助函数

### buildTree(vals []int) *TreeNode
根据层序遍历数组构建二叉树，-1表示空节点。

```go
// 示例：构建树 [3,9,20,null,null,15,7]
root := buildTree([]int{3, 9, 20, -1, -1, 15, 7})
```

### printTree(root *TreeNode)
打印二叉树的层序遍历结构，用于调试。

## 📊 测试用例说明

| 测试名称 | 输入 | 预期输出 | 说明 |
|---------|------|---------|------|
| 空树 | [] | [] | 空树应返回空结果 |
| 单节点树 | [1] | [[1]] | 只有一个节点的树 |
| 完全二叉树 | [3,9,20,-1,-1,15,7] | [[3],[9,20],[15,7]] | 标准完全二叉树 |
| 左斜树 | [1,2,-1,3,-1,-1,-1] | [[1],[2],[3]] | 所有节点都在左侧 |
| 右斜树 | [1,-1,2,-1,-1,-1,3] | [[1],[2],[3]] | 所有节点都在右侧 |
| 复杂树 | [1,2,3,4,5,6,7] | [[1],[2,3],[4,5,6,7]] | 四层完全二叉树 |

## 🎯 运行特定测试

```bash
# 运行特定测试函数
go test -run=TestDecorateRecord ./coding-interviews

# 运行示例测试
go test -run=Example ./coding-interviews

# 运行基准测试
go test -bench=BenchmarkDecorateRecord ./coding-interviews

# 运行所有测试并显示覆盖率
go test -v -cover ./coding-interviews

# 运行特定子测试
go test -run=TestDecorateRecord/空树 ./coding-interviews
```

## 📈 性能测试

基准测试使用100个节点的树来测试算法性能：

```bash
# 运行基准测试
go test -bench=. ./coding-interviews

# 运行基准测试并显示内存分配
go test -bench=. -benchmem ./coding-interviews
```

## 🐛 调试测试

如果测试失败，可以使用以下方法调试：

1. 使用 `-v` 参数查看详细输出
2. 使用 `printTree` 函数打印树结构
3. 检查测试用例的输入和预期输出

```bash
# 运行详细测试
go test -v ./coding-interviews

# 运行特定测试并显示详细输出
go test -v -run=TestDecorateRecord ./coding-interviews
```

## 📝 添加新测试

要添加新的测试用例，在 `decorate_record_test.go` 文件中的 `tests` 切片中添加新的测试用例：

```go
{
    name:     "新测试用例",
    input:    []int{1, 2, 3},
    expected: [][]int{{1}, {2, 3}},
},
```

## 🔍 故障排除

### 常见问题

1. **模块未找到错误**
   ```bash
   # 确保在项目根目录运行
   cd /path/to/GoDailyCode
   go mod tidy
   ```

2. **测试超时**
   ```bash
   # 增加超时时间
   go test -timeout=60s ./coding-interviews
   ```

3. **内存不足**
   ```bash
   # 减少基准测试的节点数量
   # 修改 decorate_record_test.go 中的 vals 数组大小
   ```

4. **测试文件未找到**
   ```bash
   # 确保在正确的目录运行测试
   # 测试文件位于 coding-interviews/decorate_record_test.go
   go test -v ./coding-interviews
   ```

5. **包名冲突**
   ```bash
   # 如果遇到包名冲突，确保所有文件都使用 package main
   # 检查 75_120.go, common.go, decorate_record_test.go 的包声明
   ```

## 🎯 实际测试结果

运行测试后，你应该看到类似以下的输出：

```bash
$ go test -v ./coding-interviews
=== RUN   TestDecorateRecord
=== RUN   TestDecorateRecord/空树
    decorate_record_test.go:96: ✓ Test passed: 空树
=== RUN   TestDecorateRecord/单节点树
    decorate_record_test.go:96: ✓ Test passed: 单节点树
=== RUN   TestDecorateRecord/完全二叉树
    decorate_record_test.go:96: ✓ Test passed: 完全二叉树
=== RUN   TestDecorateRecord/左斜树
    decorate_record_test.go:96: ✓ Test passed: 左斜树
=== RUN   TestDecorateRecord/右斜树
    decorate_record_test.go:96: ✓ Test passed: 右斜树
=== RUN   TestDecorateRecord/复杂树
    decorate_record_test.go:96: ✓ Test passed: 复杂树
--- PASS: TestDecorateRecord (0.00s)
=== RUN   ExampleDecorateRecord
--- PASS: ExampleDecorateRecord (0.00s)
PASS
ok  	DailyCopyCode/coding-interviews	0.189s
```

基准测试结果：
```bash
$ go test -bench=. ./coding-interviews
goos: darwin
goarch: arm64
pkg: DailyCopyCode/coding-interviews
cpu: Apple M2
BenchmarkDecorateRecord-8   	  863229	      1440 ns/op
PASS
ok  	DailyCopyCode/coding-interviews	2.390s
```

## 🔮 扩展测试框架

如果你想为其他算法添加测试，可以按照以下步骤：

1. **在 `coding-interviews` 包中添加新的算法文件**
   ```go
   // 例如：new_algorithm.go
   package main
   
   func newAlgorithm(input []int) []int {
       // 算法实现
   }
   ```

2. **创建对应的测试文件**
   ```go
   // 例如：new_algorithm_test.go
   package main
   
   import "testing"
   
   func TestNewAlgorithm(t *testing.T) {
       // 测试用例
   }
   ```

3. **更新 Makefile 和测试运行器**
   - 添加新的测试命令
   - 更新交互式测试程序

## 📚 更多信息

- [Go 测试文档](https://golang.org/pkg/testing/)
- [Go 基准测试](https://golang.org/pkg/testing/#hdr-Benchmarks)
- [Go 覆盖率工具](https://golang.org/cmd/cover/)
