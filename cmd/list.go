package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出可用的项目模板",
	Long:  "显示所有可用的项目模板列表",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("\n📋 可用模板:\n\n")

		templates := []struct {
			name        string
			description string
		}{
			{
				name:        "microservice",
				description: "基于 Go 的微服务模板，包含 Hertz (HTTP) + gRPC 混合架构",
			},
		}

		for _, tmpl := range templates {
			color.White("  - %s", tmpl.name)
			color.HiBlack("    %s\n", tmpl.description)
		}

		fmt.Println()
		color.Yellow("使用方法:")
		fmt.Println("  youling-gen create <project-name> -t microservice")
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
