package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Contoh menggunakan kasus user
type User struct {
	ID    uint64 `json:"id,omitempty"`
	Hash  string `json:"hash,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUsers(ctx *gin.Context) {
	users := []User{{Name: "John Wick", Email: "john.wick@email.com"}}
	ctx.JSON(http.StatusOK, users)
}

// Retrieve data menggunakan URL Params
func getUserWithID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Bad Request", "error": err.Error()},
		)
		return
	}
	user := User{ID: uint64(id), Name: "john wick", Email: "john.wick@gmail.com"}
	ctx.JSON(http.StatusOK, user)
}

// Retrieve data menggunakan Query Params
func getUserWithQueryParams(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	users := []User{{ID: 1, Name: "Heru Suksano", Email: "heru@example.com"}}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"users": users,
			"page":  page,
			"limit": limit,
		},
	)
}

// Retrieve data menggunakan form value
func postUser1(ctx *gin.Context) {
	name := ctx.Request.FormValue("name")
	email := ctx.Request.FormValue("email")

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"message": "sukses menambahkan data",
			"user": User{
				Name:  name,
				Email: email,
			},
		},
	)
}
func main() {
	r := gin.Default()

	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserWithID)
	r.GET("/users/search", getUserWithQueryParams)
	r.POST("/users", postUser1)

	r.Run(":8001")
}
