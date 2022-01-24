package migrate

import (
	"github.com/saint-yellow/go-pl/toys/gohub/database/migrations"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/migrate"
)

func migrator() *migrate.Migrator {
    // 注册 database/migrations 下的所有迁移文件
    migrations.Initialize()
    // 初始化 migrator
    return migrate.NewMigrator()
}
