package main

import (
	"Go_FinalExam/dto"
	"Go_FinalExam/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var l_db2 *gorm.DB

func getLandmark(c *gin.Context) {
	landmarks := []models.Landmark{}
	l_db2.Find(&landmarks)

	landmarks_dto := []dto.Landmark{}
	copier.Copy(&landmarks_dto, &landmarks)
	c.JSON(200, landmarks_dto)
}

func postLandmark(c *gin.Context) {
	landmarks_dto := []dto.Landmark{}
	c.ShouldBindJSON(&landmarks_dto)

	landmark := models.Landmark{}
	copier.Copy(&landmark, &landmarks_dto)

	l_db2.Create(&landmark)
	c.JSON(200, "Success")
}

func getLandmarkbyid(c *gin.Context) {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	landmark := models.Landmark{}
	l_db2.Preload("Countrydata").Where("Idx = ?", idx).Find(&landmark)
	// landmark_dto := dto.Landmark{}
	// copier.Copy(&landmark_dto, &landmark)
	c.JSON(200, landmark)
}
func getLandmarkbyname(c *gin.Context) {
	name := c.Query("name")
	landmark := []models.Landmark{}
	name = "%" + name + "%"
	l_db2.Preload("Countrydata").Where("Name like ?", name).Find(&landmark)
	// landmark_dto := dto.Landmark{}
	// copier.Copy(&landmark_dto, &landmark)
	c.JSON(200, landmark)
}

func putLandmark(c *gin.Context) {
	landmark := []models.Landmark{}
	c.ShouldBindJSON(&landmark)
	for _, l := range landmark {
		l_db2.Model(models.Landmark{}).Where("Idx = ?", l.Idx).Update("Name", "จองแล้วนะจ๊ะ")
	}
	c.JSON(200, "Success")
}

func putsLandmark(c *gin.Context) {
	_landmark := models.Landmark{}
	c.ShouldBindJSON(&_landmark)
	l_db2.Model(models.Landmark{}).Where("Idx = ?", _landmark.Idx).Updates(&_landmark)
	c.JSON(200, "Success")
}
