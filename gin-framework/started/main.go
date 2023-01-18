package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Instance engine gin dengan default config
	r := gin.Default()
	// Routing
	r.GET("/hello", welcomeController)
	// Menjalankan services di :8001
	r.Run(":8001")

	// http.ServeMux
	http.HandleFunc("/net/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "<h1><center>Hello World</center></h1>")
	})

	http.ListenAndServe(":8002", nil)
}

// Function controller
func welcomeController(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World!")
}
