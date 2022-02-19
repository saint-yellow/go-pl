package main

import (
	"fmt"
	"os"

	"github.com/saint-yellow/go-pl/toys/gohub/app/command"
	"github.com/saint-yellow/go-pl/toys/gohub/bootstrap"
	btsConfig "github.com/saint-yellow/go-pl/toys/gohub/config"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/config"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/console"
	"github.com/spf13/cobra"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	// 应用的主入口，默认调用 cmd.CmdServe 命令
    var rootCmd = &cobra.Command{
        Use:   config.Get("app.name"),
        Short: "A simple forum project",
        Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

        // rootCmd 的所有子命令都会执行以下代码
        PersistentPreRun: func(cmd *cobra.Command, args []string) {

            // 配置初始化，依赖命令行 --env 参数
            config.InitConfig(command.Env)

            bootstrap.Setup()
        },
    }

    // 注册子命令
    rootCmd.AddCommand(
        command.ServeCommand,
        command.KeyCommand,
        command.PlayCommand,
        command.Make,
        command.Migrate,
        command.Seed,
        command.Cache,
    )

    // 配置默认运行 Web 服务
    command.RegisterDefaultCmd(rootCmd, command.ServeCommand)

    // 注册全局参数，--env
    command.RegisterGlobalFlags(rootCmd)

    // 执行主命令
    if err := rootCmd.Execute(); err != nil {
        console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
    }
}
