package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewMiniBlogCommand 创建一个 *cobra.Command 对象，用于启动应用程序。
func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: " mb-apiserver",
		// 命令的简短描述
		Short: "A mini blog show best practices for develop a full-featured Go project",
		// 命令的详细描述
		Long: `A mini blog show best practices for develop a full-featured Go project.

The project features include:
• Utilization of a clean architecture;
...
• High-quality code.`,
		// 命令出错时，不打印帮助信息。设置为 true 可以确保命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello MiniBlog!")
			return nil
		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。例如：。/miniblog param1 param2
		Args: cobra.NoArgs,
	}

	return cmd
}
