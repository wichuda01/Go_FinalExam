package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCartController(router *gin.Engine, gormdb *gorm.DB) {
	db = gormdb
	cart := router.Group("/cart")
	{
		cart.GET("/add", addCart)
	}
}

func addCart(ctx *gin.Context) {
	panic("unimplemented")
}
