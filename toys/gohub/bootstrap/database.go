package bootstrap

import (
	"errors"
	"fmt"
	"time"

	"github.com/saint-yellow/go-pl/toys/gohub/pkg/config"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/database"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupDatabase 初始化数据库和 ORM
func setupDatabase() {

    var dbConfig gorm.Dialector
    switch config.Get("database.connection") {
    case "mysql":
        // 构建 DSN 信息
        dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
            config.Get("database.mysql.username"),
            config.Get("database.mysql.password"),
            config.Get("database.mysql.host"),
            config.Get("database.mysql.port"),
            config.Get("database.mysql.database"),
            config.Get("database.mysql.charset"),
        )
        dbConfig = mysql.New(mysql.Config{
            DSN: dsn,
        })
    case "postgresql":
        // 构建 DSN 信息
        dsn := fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
            config.Get("database.postgresql.username"),
            config.Get("database.postgresql.password"),
            config.Get("database.postgresql.database"),
            config.Get("database.postgresql.port"),
        )
        dbConfig = postgres.New(postgres.Config{
            DSN: dsn,

        })
    case "sqlite":
        // 初始化 sqlite
        database := config.Get("database.sqlite.database")
        dbConfig = sqlite.Open(database)
    default:
        panic(errors.New("database connection not supported"))
    }

    // 连接数据库，并设置 GORM 的日志模式
    database.Connect(dbConfig, logger.NewGormLogger())

    // 设置最大连接数
    database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
    // 设置最大空闲连接数
    database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
    // 设置每个链接的过期时间
    database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
}
