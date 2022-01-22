package bootstrap

import (
	"embed"

	"github.com/saint-yellow/go-pl/toys/goblog/pkg/view"
)

// SetupTemplate 初始化模板
func SetupTemplate(tplFS embed.FS) {
	view.TplFS = tplFS
}