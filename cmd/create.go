package cmd

import (
	"fmt"

	"github.com/JiSuanSiWeiShiXun/youling-gen/internal/generator"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	templateName string
	branch       string
	skipGit      bool
)

var createCmd = &cobra.Command{
	Use:   "create <project-name>",
	Short: "创建新的微服务项目",
	Long:  "基于 microservice-template 创建新的 Go 微服务项目",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		color.Cyan("\n🚀 欢迎使用 Youling 项目生成器！\n")

		// 创建项目配置
		config := &generator.ProjectConfig{
			ProjectName:  projectName,
			TemplateName: templateName,
			Branch:       branch,
			SkipGit:      skipGit,
		}

		// 执行生成
		gen := generator.New()
		if err := gen.Generate(config); err != nil {
			color.Red("\n❌ 创建项目失败: %v\n", err)
			return
		}

		// 显示成功信息
		color.Green("\n✨ 项目创建成功！\n")
		color.Cyan("下一步操作:\n")
		fmt.Printf("  cd %s\n", projectName)
		fmt.Println("  go mod tidy")
		fmt.Println("  make build")
		fmt.Println("\n更多信息请查看 README.md\n")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&templateName, "template", "t", "microservice", "指定模板名称")
	createCmd.Flags().StringVarP(&branch, "branch", "b", "main", "指定模板分支")
	createCmd.Flags().BoolVar(&skipGit, "skip-git", false, "跳过 Git 初始化")
}
