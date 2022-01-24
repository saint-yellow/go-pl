package command

import (
	subcommand "github.com/saint-yellow/go-pl/toys/gohub/app/command/migrate"
	"github.com/spf13/cobra"
)

var Migrate = &cobra.Command{
    Use:   "migrate",
    Short: "Run database migration",
    // 所有 migrate 下的子命令都会执行以下代码
}

func init() {
    Migrate.AddCommand(
        subcommand.Up,
        subcommand.Down,
        subcommand.Reset,
        subcommand.Refresh,
        subcommand.Fresh,
    )
}
