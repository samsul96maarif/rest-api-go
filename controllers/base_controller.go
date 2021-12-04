package controllers

import (
	"github.com/gin-gonic/gin"
)

type BaseController interface {
	Fetch(ctx *gin.Context)
	List(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
