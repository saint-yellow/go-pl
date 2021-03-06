package bootstrap

import (
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/config"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/logger"
)

// setupLogger 初始化 Logger
func setupLogger() {

    logger.InitLogger(
        config.GetString("log.filename"),
        config.GetInt("log.max_size"),
        config.GetInt("log.max_backup"),
        config.GetInt("log.max_age"),
        config.GetBool("log.compress"),
        config.GetString("log.type"),
        config.GetString("log.level"),
    )
}
