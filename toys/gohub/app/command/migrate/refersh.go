package migrate

import "github.com/spf13/cobra"

var Refresh = &cobra.Command{
    Use:   "refresh",
    Short: "Reset and re-run all migrations",
    Run:   runRefresh,
}

func runRefresh(cmd *cobra.Command, args []string) {
    migrator().Refresh()
}
