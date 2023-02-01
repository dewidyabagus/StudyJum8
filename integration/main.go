package main

import (
	"log"
	"mini-project/integration/handler/shipper"
	"mini-project/integration/httpclient"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// type FormCountry struct {
// 	ID    int `form:"id"`
// 	Limit int `form:"limit"`
// 	Page  int `form:"page"`
// }

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("file .env not found")
	}
	var (
		baseUrl = os.Getenv("SHIPPER_BASE_URL")
		token   = os.Getenv("SHIPPER_TOKEN")
	)

	shipperClient := httpclient.NewShipperAggrClient(baseUrl, token)
	shipperHandler := shipper.NewController(shipperClient)

	// ctx, cancel := context.WithTimeout(context.TODO(), 100*time.Millisecond)
	// defer cancel()

	r := gin.Default()

	r.GET("/v1/countries", shipperHandler.GetCountries)
	r.GET("/v1/provinces", shipperHandler.GetProvinces)

	r.Run("0.0.0.0:4002")

	// shipperClient := NewShipperAggrClient(baseUrl, token)

	// Handler pada package yang sama
	// r.GET("/v1/countries", func(ctx *gin.Context) {
	// 	params := new(FormCountry)
	// 	if err := ctx.BindQuery(params); err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// 	fmt.Println("Country ID:", params.ID)
	// 	res, err := shipperClient.GetCountries(ctx, params.ID, params.Limit, params.Page)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusOK, res)
	// })

}
