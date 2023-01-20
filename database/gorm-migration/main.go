package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host      = "localhost"
	port      = 3307
	username  = "root"
	password  = "root123"
	schema    = "db"
	charset   = "utf8mb4"
	loc       = "Local"
	parseTime = "True"
)

func main() {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		username, password, host, port, schema, charset, parseTime, loc,
	)
	// Membuka session database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			},
		),
	})
	if err != nil {
		log.Fatalln("gagal membuka koneksi database:", err.Error())
	}

	db.Debug().AutoMigrate(&User{}) // Auto migration table dengan grom
}
