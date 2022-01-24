package migrate

import "github.com/spf13/cobra"

var Up = &cobra.Command{
    Use:   "up",
    Short: "Run unmigrated migrations",
    Run:   runUp,
}

func runUp(cmd *cobra.Command, args []string) {
    migrator().Up()
}
