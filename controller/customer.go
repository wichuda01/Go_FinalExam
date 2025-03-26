package controller

import (
	"Go_FinalExam/dto"
	"Go_FinalExam/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// // รับข้อมูล Login body
// var req struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

func NewCustomerController(router *gin.Engine, gormdb *gorm.DB) {
	db = gormdb
	customer := router.Group("/customer")
	{
		customer.GET("/", allCustomers)
		customer.POST("/login", login)
		customer.POST("/edit:cid", editProfile)
		customer.POST("/changepassword", changePassword)

	}
}

func changePassword(ctx *gin.Context) {

}

func editProfile(ctx *gin.Context) {
	serv := services.NewCustomerService(db)
	var profileData struct {
		Address string `json:"address" binding:"required"`
	}
	cid, err := strconv.Atoi(ctx.Param("cid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	userID := uint(cid)
	ctx.ShouldBindJSON(&profileData)

	res := serv.EditProfile(userID, profileData.Address)
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"UpdateFail": res.Error})
	}
	ctx.JSON(http.StatusOK, gin.H{"UpdateSuccess": res.RowsAffected})
}

func login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	serv := services.NewCustomerService(db)
	customer, err := serv.Login(req.Email, req.Password)
	if err != nil {
		switch err.Error() {
		case "user not found":
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		case "invalid password":
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Success case
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    customer,
	})
}

func allCustomers(ctx *gin.Context) {
	serv := services.NewCustomerService(db)
	customers, err := serv.GetAllCustomer()
	//dto no pssword
	customers_dto := []dto.CustomerDtoElement{}
	copier.Copy(&customers_dto, &customers)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(200, customers_dto)

}
