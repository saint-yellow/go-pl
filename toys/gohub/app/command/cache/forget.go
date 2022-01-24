package cache

import (
	"fmt"

	"github.com/saint-yellow/go-pl/toys/gohub/pkg/cache"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/console"
	"github.com/spf13/cobra"
)

// forget 命令的选项
var cacheKey string

var Forget = &cobra.Command{
    Use:   "forget",
    Short: "Delete redis key, example: cache forget cache-key",
    Run:   runCacheForget,
}

func runCacheForget(cmd *cobra.Command, args []string) {
    cache.Forget(cacheKey)
    console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}
