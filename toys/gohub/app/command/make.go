package command

import (
	"github.com/spf13/cobra"
    subcommand "github.com/saint-yellow/go-pl/toys/gohub/app/command/make"
)

// MakeCommand 说明 cobra 命令
var Make = &cobra.Command{
    Use:   "make",
    Short: "Generate file and code",
}

func init() {
    // 注册 make 的子命令
    Make.AddCommand(
        subcommand.CMD,
		subcommand.Model,
        subcommand.APIController,
        subcommand.Request,
        subcommand.Migration,
    )
}

