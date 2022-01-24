package cache

func init() {
	// 设置 cache forget 命令的选项
	Forget.Flags().StringVarP(&cacheKey, "key", "k", "", "KEY of the cache")
	Forget.MarkFlagRequired("key")
}