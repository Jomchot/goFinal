package controller

import (
	"goFinal/model"
	"goFinal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dbProduct *gorm.DB

func NewProduct(router *gin.Engine, gormdb *gorm.DB) {
	dbProduct = gormdb
	product := router.Group("/product")
	{
		product.GET("", getAllUsers)
		product.PUT("/password", updatePassord)
	}

}

func getProduct(ctx *gin.Context) {
	user := model.Customer{}
	err := ctx.ShouldBindJSON(&user) //รับค่าจาก body
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email := user.Email
	service := service.NewUsersService(dbProduct)

	userAuth, err := service.PostAuthLogin(email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userAuth)
}
