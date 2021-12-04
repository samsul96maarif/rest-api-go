package main

import (
	"github.com/samsul96maarif/rest-api-go/controllers"
	"github.com/samsul96maarif/rest-api-go/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err.Error())
	}
	r := gin.Default()
	productController := controllers.RegisterProductController(db)
	r.GET("/", productController.Fetch)
	r.Run()
}
