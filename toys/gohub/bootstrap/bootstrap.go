package bootstrap

func Setup() {
	// 初始化 Logger
    setupLogger()

	// 初始化数据库
	setupDatabase()

	// 初始化 Redis
	setupRedis()

	// 初始化缓存
	setupCache()
}