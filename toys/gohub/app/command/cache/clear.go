package cache

import (
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/cache"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var Clear = &cobra.Command{
    Use:   "clear",
    Short: "Clear cache",
    Run:   runCacheClear,
}

func runCacheClear(cmd *cobra.Command, args []string) {
    cache.Flush()
    console.Success("Cache cleared.")
}
