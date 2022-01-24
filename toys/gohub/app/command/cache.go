package command

import (
	subcommand "github.com/saint-yellow/go-pl/toys/gohub/app/command/cache"

	"github.com/spf13/cobra"
)

var Cache = &cobra.Command{
    Use:   "cache",
    Short: "Cache management",
}

func init() {
    // 注册 cache 命令的子命令
    Cache.AddCommand(
        subcommand.Clear,
        subcommand.Forget,
    )
}
