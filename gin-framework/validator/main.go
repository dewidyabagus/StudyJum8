package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

// Gin model binding and validation
// Source:
// - https://gin-gonic.com/docs/examples/binding-and-validation/
// - https://gin-gonic.com/docs/examples/custom-validators/

// Instance validator
var (
	validate *validator.Validate
)

// Struktur transaksi yang diikat dengan tag validasi data
type Transaction struct {
	ID       uint64  `json:"id" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Date     string  `json:"date" binding:"datetime=2006-01-02 15:04:05,today"`
	Price    float64 `json:"price" binding:"required"`
	Quantity float32 `json:"quantity" binding:"required"`
}

// Struktur untuk respose
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Handler untuk transaksi
func TransactionHandler(ctx *gin.Context) {
	var trx Transaction
	if err := ctx.ShouldBindJSON(&trx); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			Response{
				// Response dengan pesan error ketika gagal binding data
				// atau data yang dikirimkan tidak valid
				Message: translateError(err).Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, Response{Message: "success create trx", Data: trx})
}

// Main services
func main() {
	r := gin.Default()

	// Register validation function dengan tag today
	var ok bool
	if validate, ok = binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("today", todayValFunc)
	}

	r.POST("/transaction", TransactionHandler)
	r.Run(":8001")
}
