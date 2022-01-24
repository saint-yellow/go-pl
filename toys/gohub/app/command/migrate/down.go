package migrate

import "github.com/spf13/cobra"

var DownSubcommand = &cobra.Command{
    Use: "down",
    // 设置别名 migrate down == migrate rollback
    Aliases: []string{"rollback"},
    Short:   "Reverse the up command",
    Run:     runDown,
}

func runDown(cmd *cobra.Command, args []string) {
    migrator().Rollback()
}
