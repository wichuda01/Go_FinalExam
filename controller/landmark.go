package controller

import (
	"Go_FinalExam/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewLandmarkController(router *gin.Engine, gormdb *gorm.DB) {
	db = gormdb
	landmark := router.Group("/landmark")
	{
		landmark.GET("/", allLandmark)
		// landmark.GET("/search", getLandmarkByName)
		// landmark.POST("/add", addLandmark)
		// landmark.GET("/:lid", getLandmarkByID)
		// landmark.GET("/user/:uid", getLandmarkByuid)
	}
}

func allLandmark(ctx *gin.Context) {
	serv := services.NewLandmarkService(db)
	landmarks, err := serv.GetAllLandmark()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, landmarks)
}
