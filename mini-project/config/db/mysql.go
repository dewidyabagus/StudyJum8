package db

import (
	"fmt"

	"github.com/dewidyabagus/go-hexagonal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const charset = "utf8mb4"
const parseTime = "True"

func NewMySQL(cfg *config.Database) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Schema, charset, parseTime, cfg.Loc,
	)

	// Untuk konfigurasi ORM GORM
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
	}

	if db, err = gorm.Open(mysql.Open(dsn), gormConfig); err != nil {
		return
	}

	return
}
