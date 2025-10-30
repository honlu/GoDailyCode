package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("=== Go Daily Code 测试运行器 ===")
	fmt.Println("当前工作目录:", getCurrentDir())

	// 检查是否在正确的目录
	if !isGoModule() {
		fmt.Println("❌ 错误: 当前目录不是Go模块")
		fmt.Println("请确保在项目根目录运行此程序")
		return
	}

	fmt.Println("✅ 检测到Go模块")

	// 显示可用的测试命令
	showTestCommands()

	// 运行基本测试
	fmt.Println("\n=== 开始运行测试 ===")
	runBasicTests()
}

func getCurrentDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func isGoModule() bool {
	// 检查上级目录的go.mod文件
	parentDir := filepath.Dir(getCurrentDir())
	_, err := os.Stat(filepath.Join(parentDir, "go.mod"))
	return !os.IsNotExist(err)
}

func showTestCommands() {
	fmt.Println("\n=== 可用的测试命令 ===")
	commands := []struct {
		name        string
		command     string
		description string
	}{
		{"单元测试", "go test -v ./coding-interviews", "运行coding-interviews包的单元测试"},
		{"基准测试", "go test -bench=. ./coding-interviews", "运行性能基准测试"},
		{"覆盖率测试", "go test -cover ./coding-interviews", "查看测试覆盖率"},
		{"所有测试", "go test -v -bench=. ./coding-interviews", "运行所有测试"},
		{"示例测试", "go test -run=Example ./coding-interviews", "运行示例测试"},
		{"整个项目测试", "go test -v ./...", "运行整个项目的所有测试"},
		{"特定测试", "go test -run=TestDecorateRecord ./coding-interviews", "运行特定测试函数"},
	}

	for i, cmd := range commands {
		fmt.Printf("%d. %s\n", i+1, cmd.name)
		fmt.Printf("   命令: %s\n", cmd.command)
		fmt.Printf("   说明: %s\n\n", cmd.description)
	}
}

func runBasicTests() {
	// 切换到项目根目录
	projectRoot := filepath.Dir(getCurrentDir())
	os.Chdir(projectRoot)

	// 运行coding-interviews包的测试
	fmt.Println("1. 运行 coding-interviews 包测试...")
	runCommand("go", "test", "-v", "./coding-interviews")

	fmt.Println("\n2. 运行测试覆盖率...")
	runCommand("go", "test", "-cover", "./coding-interviews")

	fmt.Println("\n3. 运行基准测试...")
	runCommand("go", "test", "-bench=.", "./coding-interviews")
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("执行命令: %s %v\n", name, args)
	fmt.Println("----------------------------------------")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ 命令执行失败: %v\n", err)
	} else {
		fmt.Println("✅ 命令执行成功")
	}
	fmt.Println("----------------------------------------")
}
