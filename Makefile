# 定义项目名称
PROJECT_NAME := golang-example
# 定义 Go 命令
GO := go

# 构建项目
build:
	$(GO) build -o $(PROJECT_NAME) .

# 运行项目
run: build
	./$(PROJECT_NAME)

# 测试项目
test:
	$(GO) test ./...

# 清理生成的文件
clean:
	rm -f $(PROJECT_NAME)

# 默认目标，当只输入 `make` 时执行
.DEFAULT_GOAL := build