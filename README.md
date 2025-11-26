# Youling Gen

一个用 Go 实现的脚手架工具，用于快速创建基于 [microservice-template](https://github.com/JiSuanSiWeiShiXun/microservice-template) 的微服务项目。

## ✨ 特性

- 🚀 快速创建 Go 微服务项目
- 📦 基于 Hertz (HTTP) + gRPC 混合架构
- 🎨 交互式命令行界面
- ⚙️ 自动配置项目信息
- 🔧 支持自定义模板分支
- 📝 自动生成项目元数据

## 📦 安装

### 从源码安装

```bash
git clone https://github.com/JiSuanSiWeiShiXun/youling-gen.git
cd youling-gen
make install
```

### 使用 go install

```bash
go install github.com/JiSuanSiWeiShiXun/youling-gen@latest
```

## 🚀 快速开始

### 创建新项目

```bash
# 基本用法
youling-gen create my-project

# 指定模板分支
youling-gen create my-project -b main

# 跳过 Git 初始化
youling-gen create my-project --skip-git
```

### 列出可用模板

```bash
youling-gen list
```

### 查看帮助

```bash
youling-gen --help
youling-gen create --help
```

## 📖 使用示例

创建一个新的微服务项目：

```bash
$ youling-gen create my-microservice

🚀 欢迎使用 Youling 项目生成器！

? 请输入项目描述: 我的第一个微服务项目
? 请输入项目作者: Your Name
? 是否初始化 Git 仓库? Yes
✓ 模板下载成功！
✓ 项目配置完成！
✓ Git 仓库初始化成功！

✨ 项目创建成功！

下一步操作:
  cd my-microservice
  go mod tidy
  make build

更多信息请查看 README.md
```

## 🏗️ 项目结构

```
youling-gen/
├── cmd/                    # 命令行入口
│   ├── root.go            # 根命令
│   ├── create.go          # create 命令
│   └── list.go            # list 命令
├── internal/
│   └── generator/         # 项目生成器
│       └── generator.go
├── bin/                   # 构建输出
├── main.go               # 主入口
├── go.mod
├── Makefile
└── README.md
```

## 🛠️ 开发

### 构建项目

```bash
make build
```

### 运行测试

```bash
make test
```

### 格式化代码

```bash
make fmt
```

### 本地运行

```bash
make run
```

## 📋 命令说明

### create

创建新的微服务项目。

```bash
youling-gen create <project-name> [flags]
```

**标志：**
- `-t, --template string`: 指定模板名称（默认: "microservice"）
- `-b, --branch string`: 指定模板分支（默认: "main"）
- `--skip-git`: 跳过 Git 初始化

### list

列出所有可用的项目模板。

```bash
youling-gen list
```

## 🔧 配置

项目创建后会在根目录生成 `.youling.json` 文件，包含项目元数据：

```json
{
  "name": "my-project",
  "description": "项目描述",
  "author": "作者名称",
  "version": "1.0.0",
  "created_at": "2025-11-26T10:00:00Z"
}
```

## 📚 模板说明

### microservice

基于 Go 的微服务模板，包含：

- **架构**: Hertz (HTTP) + gRPC 混合架构
- **功能**: 
  - Protocol Buffers 代码生成
  - 多语言支持 (Go, Python, TypeScript)
  - Wire 依赖注入
  - 分层架构设计
  - 完整的项目结构
  - 配置管理（Viper）
  - 日志系统（Zap）

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

## 🔗 相关链接

- [microservice-template](https://github.com/JiSuanSiWeiShiXun/microservice-template) - 模板项目
- [Cobra](https://github.com/spf13/cobra) - CLI 框架
- [Survey](https://github.com/AlecAivazis/survey) - 交互式提示
