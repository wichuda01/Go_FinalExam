package controller

import (
	"Go_FinalExam/models"
	"Go_FinalExam/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCartItemController(router *gin.Engine, gormdb *gorm.DB) {
	db = gormdb
	cartitem := router.Group("/cartitem")
	{
		cartitem.POST("/add", addCartItem)
	}
}

func addCartItem(ctx *gin.Context) {
	serv := services.NewCartItemService(db)
	item := models.CartItem{}
	ctx.ShouldBindJSON(&item)
	res := serv.AddcartItem(item)
	if res.Error != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"insertfail": res.Error})
	}
	ctx.JSON(http.StatusOK, gin.H{"insertsuccess": res.RowsAffected})
}
