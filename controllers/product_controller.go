package controllers

import (
	"database/sql"
	"net/http"

	"github.com/samsul96maarif/rest-api-go/database"
	"github.com/samsul96maarif/rest-api-go/services"

	"github.com/gin-gonic/gin"
)

type ProductControllerInterface interface {
	BaseController
}

type productController struct {
	service services.ProductServiceInterface
}

func NewController(service services.ProductServiceInterface) ProductControllerInterface {
	return &productController{
		service: service,
	}
}

func (controller productController) List(c *gin.Context)     {}
func (controller productController) Delete(c *gin.Context)   {}
func (controller productController) FindById(c *gin.Context) {}
func (controller productController) Update(c *gin.Context)   {}
func (controller productController) Store(c *gin.Context)    {}

func (controller productController) Fetch(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	tem, err := controller.service.Fetch()
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": tem})
}

func RegisterProductController(db *sql.DB) ProductControllerInterface {
	service := services.RegisterProductService(db)
	return NewController(service)
}
