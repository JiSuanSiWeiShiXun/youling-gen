# Shell 自动补全配置指南

`youling-gen` 支持多种 Shell 的自动补全功能，让命令行使用更加便捷。

## 🎯 什么是自动补全？

自动补全允许你在输入命令时按 `Tab` 键自动完成命令、子命令和参数：

```bash
$ youling-gen cr[Tab]
$ youling-gen create    # 自动补全

$ youling-gen create --[Tab]
--branch    --skip-git    --template    # 显示可用选项
```

## 📦 安装配置

### Bash

#### macOS

```bash
# 1. 安装 bash-completion (如果还没安装)
brew install bash-completion

# 2. 生成补全脚本
youling-gen completion bash > $(brew --prefix)/etc/bash_completion.d/youling-gen

# 3. 添加到 ~/.bash_profile
echo 'source $(brew --prefix)/etc/bash_completion' >> ~/.bash_profile

# 4. 重新加载
source ~/.bash_profile
```

#### Linux

```bash
# 1. 生成补全脚本
sudo youling-gen completion bash > /etc/bash_completion.d/youling-gen

# 2. 重新加载
source /etc/bash_completion.d/youling-gen

# 或者添加到用户配置
youling-gen completion bash >> ~/.bashrc
source ~/.bashrc
```

### Zsh

#### macOS / Linux

```bash
# 1. 确保补全功能已启用（添加到 ~/.zshrc）
autoload -Uz compinit
compinit

# 2. 生成补全脚本
mkdir -p ~/.zsh/completion
youling-gen completion zsh > ~/.zsh/completion/_youling-gen

# 3. 添加补全目录到 fpath（添加到 ~/.zshrc）
fpath=(~/.zsh/completion $fpath)

# 4. 重新加载
source ~/.zshrc
```

#### Oh My Zsh

```bash
# 1. 生成补全脚本到 Oh My Zsh 插件目录
youling-gen completion zsh > ~/.oh-my-zsh/completions/_youling-gen

# 2. 重新加载
source ~/.zshrc
```

### Fish

```bash
# 1. 生成补全脚本
youling-gen completion fish > ~/.config/fish/completions/youling-gen.fish

# 2. 重新加载
source ~/.config/fish/config.fish
```

### PowerShell

```powershell
# 1. 生成补全脚本
youling-gen completion powershell > youling-gen.ps1

# 2. 添加到 PowerShell 配置文件
# 找到配置文件路径
$PROFILE

# 添加以下内容到配置文件
. path/to/youling-gen.ps1
```

## 🔍 补全功能演示

### 命令补全

```bash
$ youling-gen [Tab]
completion  create      help        list
```

### 子命令补全

```bash
$ youling-gen create [Tab]
# 显示项目名称提示或当前目录下的文件夹
```

### 标志（Flags）补全

```bash
$ youling-gen create --[Tab]
--branch     # 指定模板分支
--skip-git   # 跳过 Git 初始化
--template   # 指定模板名称
```

### 值补全

```bash
$ youling-gen create -t [Tab]
microservice    # 显示可用模板
```

## 🛠️ 测试补全功能

安装后测试是否正常工作：

```bash
# 测试命令补全
$ youling-gen [Tab][Tab]

# 测试标志补全
$ youling-gen create --[Tab][Tab]

# 测试帮助信息
$ youling-gen completion --help
```

## 📝 支持的 Shell

| Shell       | 支持状态 | 安装命令 |
|-------------|----------|----------|
| Bash        | ✅       | `youling-gen completion bash` |
| Zsh         | ✅       | `youling-gen completion zsh` |
| Fish        | ✅       | `youling-gen completion fish` |
| PowerShell  | ✅       | `youling-gen completion powershell` |

## 🔧 故障排除

### 补全不工作

1. **确认 Shell 类型**
   ```bash
   echo $SHELL
   ```

2. **检查补全脚本是否加载**
   ```bash
   # Bash
   complete -p youling-gen
   
   # Zsh
   which _youling-gen
   ```

3. **重新生成并加载**
   ```bash
   # 重新生成补全脚本
   youling-gen completion bash > ~/.youling-gen-completion.bash
   
   # 手动加载测试
   source ~/.youling-gen-completion.bash
   ```

### macOS Zsh 权限问题

```bash
# 如果遇到 "insecure directories" 错误
chmod -R 755 ~/.zsh/completion
compaudit | xargs chmod g-w
```

### Linux 权限问题

```bash
# 确保补全脚本有执行权限
sudo chmod +x /etc/bash_completion.d/youling-gen
```

## 💡 高级技巧

### 临时启用补全

不想永久安装，只在当前会话启用：

```bash
# Bash
source <(youling-gen completion bash)

# Zsh
source <(youling-gen completion zsh)
```

### 为所有用户安装（需要 root 权限）

```bash
# Bash - 系统级
sudo youling-gen completion bash > /usr/share/bash-completion/completions/youling-gen

# Zsh - 系统级
sudo youling-gen completion zsh > /usr/share/zsh/site-functions/_youling-gen
```

## 📚 相关资源

- [Cobra 自动补全文档](https://github.com/spf13/cobra/blob/master/shell_completions.md)
- [Bash 补全教程](https://www.gnu.org/software/bash/manual/html_node/Programmable-Completion.html)
- [Zsh 补全系统](http://zsh.sourceforge.net/Doc/Release/Completion-System.html)

## 🎉 享受高效的命令行体验！

启用自动补全后，使用 `youling-gen` 会更加流畅和高效！
