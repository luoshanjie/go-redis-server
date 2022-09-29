.DEFAULT_GOAL=run
PROJECT_NAME=go-redis-server
BINARY=rediserver
OUTFILE=c.out

# 依赖
dep:
	@echo "[${PROJECT_NAME}] get go framework dependencies"
	@go mod tidy

# 清理
clean:
	@echo "[${PROJECT_NAME}] clean"
	@if [ -f bin/${BINARY} ]; then rm bin/${BINARY}; fi
	@if [ -f ${OUTFILE} ]; then rm ${OUTFILE}; fi

# 编译
build: clean
	@echo "[${PROJECT_NAME}] build"
	@go build -o bin/${BINARY}

# 运行
run: build
	@echo "[${PROJECT_NAME}] run"
	@bin/${BINARY}

# 测试
test:
	@echo "[${PROJECT_NAME}] test"
	@go test ./...

.PHONY: dep build test cover clean check lint help
