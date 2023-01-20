package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var i int
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Aku sudah ditutup")
				return

			case <-ticker.C:
				log.Println("Waktunya di eksekusi [", i, "]")
				i++
			}
		}
	}()

	fmt.Println("Jumlah Goroutine:", runtime.NumGoroutine())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	cancel()

	time.Sleep(time.Second)
}
