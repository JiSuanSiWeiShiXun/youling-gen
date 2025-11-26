# Youling Gen - 使用示例

## 基本使用

### 1. 创建新项目

```bash
youling-gen create my-project
```

这将启动交互式向导，提示你输入：
- 项目描述
- 项目作者
- 是否初始化 Git 仓库

### 2. 指定选项创建项目

```bash
# 使用指定的模板分支
youling-gen create my-project -b main

# 跳过 Git 初始化
youling-gen create my-project --skip-git

# 组合使用
youling-gen create my-project -b develop --skip-git
```

### 3. 查看可用模板

```bash
youling-gen list
```

## 完整示例

### 创建一个微服务项目

```bash
$ youling-gen create payment-service

🚀 欢迎使用 Youling 项目生成器！

? 请输入项目描述: 支付微服务
? 请输入项目作者: John Doe
? 是否初始化 Git 仓库? Yes
 正在下载模板...
✓ 模板下载成功！
 正在配置项目...
✓ 项目配置完成！
 正在初始化 Git 仓库...
✓ Git 仓库初始化成功！

✨ 项目创建成功！

下一步操作:
  cd payment-service
  go mod tidy
  make build

更多信息请查看 README.md
```

### 进入项目并启动

```bash
# 进入项目目录
cd payment-service

# 安装依赖
go mod tidy

# 编译 proto 文件
make build

# 运行 HTTP 服务（端口 8080）
go run cmd/api-gateway/main.go

# 运行 gRPC 服务（端口 50051）
go run cmd/adhoc-server/main.go
```

## 项目结构说明

创建的项目包含以下主要目录：

```
payment-service/
├── api/                    # Proto 定义文件
├── cmd/                    # 服务入口
│   ├── api-gateway/       # HTTP 网关
│   └── adhoc-server/      # gRPC 服务
├── internal/              # 内部实现
│   ├── shared/           # 共享模块
│   ├── api/              # API 服务实现
│   └── adhoc/            # Adhoc 服务实现
├── pkg/                   # 公共库
├── gen/                   # 生成的代码
├── config.yml            # 配置文件
├── Makefile              # 构建脚本
└── README.md             # 项目文档
```

## 常见问题

### Q: 如何更换模板版本？

A: 使用 `-b` 参数指定分支：
```bash
youling-gen create my-project -b v1.0.0
```

### Q: 创建失败怎么办？

A: 确保：
1. Git 已安装并可访问
2. 有网络连接可以访问 GitHub
3. 目标目录不存在

### Q: 如何自定义项目配置？

A: 项目创建后，可以修改：
1. `config.yml` - 应用配置
2. `go.mod` - 依赖管理
3. `Makefile` - 构建脚本

### Q: 生成的项目包含什么？

A: 包含：
- Hertz HTTP 框架
- gRPC 服务框架
- Proto 代码生成工具
- 分层架构设计
- 配置管理（Viper）
- 日志系统（Zap）
- 完整的示例代码

## 高级用法

### 自动化脚本

创建多个项目：

```bash
#!/bin/bash

projects=(
  "user-service"
  "order-service"
  "payment-service"
)

for project in "${projects[@]}"; do
  echo "创建 $project..."
  echo -e "$project 微服务\nTeam Name\ny" | youling-gen create "$project"
done
```

### CI/CD 集成

在 CI/CD 流程中使用：

```yaml
# .github/workflows/create-service.yml
name: Create Service

on:
  workflow_dispatch:
    inputs:
      service_name:
        description: '服务名称'
        required: true

jobs:
  create:
    runs-on: ubuntu-latest
    steps:
      - name: Install youling-gen
        run: go install github.com/JiSuanSiWeiShiXun/youling-gen@latest
      
      - name: Create service
        run: |
          youling-gen create ${{ github.event.inputs.service_name }} --skip-git
```

## 更新日志

### v1.0.0 (2025-11-26)
- ✨ 初始版本发布
- ✅ 支持从 microservice-template 创建项目
- ✅ 交互式命令行界面
- ✅ 自动配置项目信息
- ✅ Git 仓库初始化
