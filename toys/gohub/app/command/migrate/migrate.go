package migrate

import (
	"github.com/saint-yellow/go-pl/toys/gohub/database/migrations"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/migrate"
	"github.com/spf13/cobra"
)

var MigrateCommand = &cobra.Command{
    Use:   "migrate",
    Short: "Run database migration",
    // 所有 migrate 下的子命令都会执行以下代码
}

func init() {
    MigrateCommand.AddCommand(
        UpSubcommand,
    )
}

func migrator() *migrate.Migrator {
    // 注册 database/migrations 下的所有迁移文件
    migrations.Initialize()
    // 初始化 migrator
    return migrate.NewMigrator()
}
