package migrate

import "github.com/spf13/cobra"

var Reset = &cobra.Command{
    Use:   "reset",
    Short: "Rollback all database migrations",
    Run:   runReset,
}

func runReset(cmd *cobra.Command, args []string) {
    migrator().Reset()
}
