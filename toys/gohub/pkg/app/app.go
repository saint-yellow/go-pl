// Package app 应用信息
package app

import "github.com/saint-yellow/go-pl/toys/gohub/pkg/config"

func IsLocal() bool {
    return config.Get("app.env") == "local"
}

func IsProduction() bool {
    return config.Get("app.env") == "production"
}

func IsTesting() bool {
    return config.Get("app.env") == "testing"
}
