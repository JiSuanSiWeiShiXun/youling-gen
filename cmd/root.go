package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
)

var rootCmd = &cobra.Command{
	Use:   "youling-gen",
	Short: "Youling 项目生成器",
	Long: color.CyanString(`
┬ ┬┌─┐┬ ┬┬  ┬┌┐┌┌─┐   ┌─┐┌─┐┌┐┌
└┬┘│ ││ ││  │││││ ┬───│ ┬├┤ │││
 ┴ └─┘└─┘┴─┘┴┘└┘└─┘   └─┘└─┘┘└┘

基于 microservice-template 的 Go 微服务脚手架工具
`),
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red("错误: %v", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(fmt.Sprintf("youling-gen version %s\n", version))
}
