package controller

import (
	"Go_FinalExam/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewProductController(router *gin.Engine, gormdb *gorm.DB) {
	db = gormdb
	landmark := router.Group("/product")
	{
		landmark.GET("/", allProducts)
		// landmark.GET("/search", getLandmarkByName)
		// landmark.POST("/add", addLandmark)
		// landmark.GET("/:lid", getLandmarkByID)
		// landmark.GET("/user/:uid", getLandmarkByuid)
	}
}

func allProducts(ctx *gin.Context) {
	serv := services.NewProductService(db)
	landmarks, err := serv.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, landmarks)
}
