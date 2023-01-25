package main

import (
	"log"

	"github.com/dewidyabagus/go-hexagonal/config"
	"github.com/dewidyabagus/go-hexagonal/config/db"
	"github.com/dewidyabagus/go-hexagonal/modules/database/user"
	"github.com/dewidyabagus/go-hexagonal/utils/logger"
)

func main() {
	cfg := config.LoadConfigAPI("./config")

	mlog, err := logger.New("info")
	if err != nil {
		log.Fatalln(err.Error())
	}

	db, err := db.NewMySQL(&cfg.Database)
	if err != nil {
		log.Fatalln("open mysql failed:", err.Error())
	}

	mlog.Info("start migration model ...")

	// Add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user.User{})

	mlog.Info("success migration model ...")
}
