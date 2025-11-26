package generator

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type ProjectConfig struct {
	ProjectName  string
	TemplateName string
	Branch       string
	SkipGit      bool
	Description  string
	Author       string
	UseGit       bool
}

type ProjectMetadata struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Version     string    `json:"version"`
	CreatedAt   time.Time `json:"created_at"`
}

type Generator struct {
	spinner *spinner.Spinner
}

func New() *Generator {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("cyan")
	return &Generator{
		spinner: s,
	}
}

func (g *Generator) Generate(config *ProjectConfig) error {
	// 检查项目目录是否已存在
	targetDir, err := filepath.Abs(config.ProjectName)
	if err != nil {
		return fmt.Errorf("无法获取绝对路径: %w", err)
	}

	if _, err := os.Stat(targetDir); !os.IsNotExist(err) {
		return fmt.Errorf("目录 %s 已存在", config.ProjectName)
	}

	// 询问用户配置
	if err := g.promptUserConfig(config); err != nil {
		return err
	}

	// 下载模板
	if err := g.downloadTemplate(config, targetDir); err != nil {
		// 清理失败的目录
		os.RemoveAll(targetDir)
		return err
	}

	// 配置项目
	if err := g.configureProject(config, targetDir); err != nil {
		return err
	}

	// 初始化 Git
	if config.UseGit && !config.SkipGit {
		if err := g.initGit(targetDir); err != nil {
			color.Yellow("⚠️  Git 初始化失败: %v", err)
		}
	}

	return nil
}

func (g *Generator) promptUserConfig(config *ProjectConfig) error {
	questions := []*survey.Question{
		{
			Name: "description",
			Prompt: &survey.Input{
				Message: "请输入项目描述:",
				Default: "一个基于 Go 的微服务项目",
			},
		},
		{
			Name: "author",
			Prompt: &survey.Input{
				Message: "请输入项目作者:",
				Default: "",
			},
		},
	}

	if !config.SkipGit {
		questions = append(questions, &survey.Question{
			Name: "useGit",
			Prompt: &survey.Confirm{
				Message: "是否初始化 Git 仓库?",
				Default: true,
			},
		})
	}

	answers := struct {
		Description string
		Author      string
		UseGit      bool
	}{}

	if err := survey.Ask(questions, &answers); err != nil {
		return err
	}

	config.Description = answers.Description
	config.Author = answers.Author
	config.UseGit = answers.UseGit

	return nil
}

func (g *Generator) downloadTemplate(config *ProjectConfig, targetDir string) error {
	g.spinner.Suffix = " 正在下载模板..."
	g.spinner.Start()
	defer g.spinner.Stop()

	// 使用 git clone 下载模板
	repoURL := fmt.Sprintf("https://github.com/JiSuanSiWeiShiXun/microservice-template.git")

	cmd := exec.Command("git", "clone", "--depth", "1", "--branch", config.Branch, repoURL, targetDir)
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("下载模板失败: %w", err)
	}

	// 删除 .git 目录
	gitDir := filepath.Join(targetDir, ".git")
	if err := os.RemoveAll(gitDir); err != nil {
		return fmt.Errorf("删除 .git 目录失败: %w", err)
	}

	g.spinner.Stop()
	color.Green("✓ 模板下载成功！")

	return nil
}

func (g *Generator) configureProject(config *ProjectConfig, targetDir string) error {
	g.spinner.Suffix = " 正在配置项目..."
	g.spinner.Start()
	defer g.spinner.Stop()

	// 更新 go.mod
	if err := g.updateGoMod(config, targetDir); err != nil {
		return err
	}

	// 更新 README.md
	if err := g.updateReadme(config, targetDir); err != nil {
		return err
	}

	// 创建项目元数据文件
	if err := g.createMetadata(config, targetDir); err != nil {
		return err
	}

	g.spinner.Stop()
	color.Green("✓ 项目配置完成！")

	return nil
}

func (g *Generator) updateGoMod(config *ProjectConfig, targetDir string) error {
	goModPath := filepath.Join(targetDir, "go.mod")

	content, err := os.ReadFile(goModPath)
	if err != nil {
		// 如果没有 go.mod，跳过
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("读取 go.mod 失败: %w", err)
	}

	// 替换 module 名称
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "module ") {
			lines[i] = fmt.Sprintf("module %s", config.ProjectName)
			break
		}
	}

	newContent := strings.Join(lines, "\n")
	if err := os.WriteFile(goModPath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("写入 go.mod 失败: %w", err)
	}

	return nil
}

func (g *Generator) updateReadme(config *ProjectConfig, targetDir string) error {
	readmePath := filepath.Join(targetDir, "README.md")

	content, err := os.ReadFile(readmePath)
	if err != nil {
		// 如果没有 README，创建一个简单的
		if os.IsNotExist(err) {
			newReadme := fmt.Sprintf("# %s\n\n%s\n\n## 作者\n\n%s\n",
				config.ProjectName, config.Description, config.Author)
			return os.WriteFile(readmePath, []byte(newReadme), 0644)
		}
		return fmt.Errorf("读取 README.md 失败: %w", err)
	}

	// 更新 README 内容
	lines := strings.Split(string(content), "\n")
	var newLines []string

	// 添加新的标题和描述
	newLines = append(newLines, fmt.Sprintf("# %s", config.ProjectName))
	newLines = append(newLines, "")
	newLines = append(newLines, config.Description)
	newLines = append(newLines, "")
	if config.Author != "" {
		newLines = append(newLines, "## 作者")
		newLines = append(newLines, "")
		newLines = append(newLines, config.Author)
		newLines = append(newLines, "")
	}

	// 跳过原来的第一行标题
	skipFirst := true
	for _, line := range lines {
		if skipFirst && strings.HasPrefix(line, "#") {
			skipFirst = false
			continue
		}
		if !skipFirst {
			newLines = append(newLines, line)
		}
	}

	newContent := strings.Join(newLines, "\n")
	if err := os.WriteFile(readmePath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("写入 README.md 失败: %w", err)
	}

	return nil
}

func (g *Generator) createMetadata(config *ProjectConfig, targetDir string) error {
	metadata := ProjectMetadata{
		Name:        config.ProjectName,
		Description: config.Description,
		Author:      config.Author,
		Version:     "1.0.0",
		CreatedAt:   time.Now(),
	}

	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化元数据失败: %w", err)
	}

	metadataPath := filepath.Join(targetDir, ".youling.json")
	if err := os.WriteFile(metadataPath, data, 0644); err != nil {
		return fmt.Errorf("写入元数据失败: %w", err)
	}

	return nil
}

func (g *Generator) initGit(targetDir string) error {
	g.spinner.Suffix = " 正在初始化 Git 仓库..."
	g.spinner.Start()
	defer g.spinner.Stop()

	cmd := exec.Command("git", "init")
	cmd.Dir = targetDir
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Run(); err != nil {
		return err
	}

	g.spinner.Stop()
	color.Green("✓ Git 仓库初始化成功！")

	return nil
}
