package migrations

import (
	"database/sql"

	"github.com/saint-yellow/go-pl/toys/gohub/app/models"
	"github.com/saint-yellow/go-pl/toys/gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
    type Link struct {
        models.BaseModel

        Name string `gorm:"type:varchar(255);not null"`
        URL  string `gorm:"type:varchar(255);default:null"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&Link{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&Link{})
    }

    migrate.Add("2022_01_24_223219_add_links_table", up, down)
}
