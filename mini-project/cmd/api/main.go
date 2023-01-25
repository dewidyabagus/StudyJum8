package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dewidyabagus/go-hexagonal/cmd/api/routes"
	"github.com/dewidyabagus/go-hexagonal/config"
	"github.com/dewidyabagus/go-hexagonal/config/db"
	"github.com/dewidyabagus/go-hexagonal/utils/logger"

	"github.com/gin-gonic/gin"
)

//Mode application
const (
	production = "production"
)

func main() {
	cfg := config.LoadConfigAPI("./config")

	mlog, err := logger.New("info")
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Ketergantungan yang digunakan
	db, err := db.NewMySQL(&cfg.Database)
	if err != nil {
		log.Fatalln("open mysql failed:", err.Error())
	}

	// Memasukan ketergantungan yang digunakan untuk kebutuhan health check
	depends := &dependencies{db: db}

	// Setup service REST API
	if cfg.Env == production {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	routing := &routes.HandlerConfig{
		Config:      cfg,
		HealthCheck: depends,
	}
	prepareService(routing, depends)
	routing.CreateRouting(r)

	svc := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}

	go func() {
		//Service connections
		if err := svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("listen and serve failed:", err.Error())
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 30 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	mlog.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := svc.Shutdown(ctx); err != nil {
		log.Fatalln("shutting down failed:", err.Error())
	}

	mlog.Info("server exiting")
}
