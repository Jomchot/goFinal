package controller

import (
	"goFinal/model"
	"goFinal/service"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewUsers(router *gin.Engine, gormdb *gorm.DB) {
	db = gormdb
	user := router.Group("/user")
	{
		user.GET("", getAllUsers)
		user.POST("", insertUser)
		user.PUT("/password", updatePassord)
	}
	auth := router.Group("/auth")
	{
		auth.POST("/login", postUserByEmail)
	}
}

func getAllUsers(ctx *gin.Context) {
	service := service.NewUsersService(db)
	getUser := service.GetAllUsers()
	ctx.JSON(http.StatusOK, getUser)
}

func postUserByEmail(ctx *gin.Context) {
	user := model.Customer{}
	err := ctx.ShouldBindJSON(&user) //รับค่าจาก body
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	email := user.Email
	service := service.NewUsersService(db)

	userAuth, err := service.PostAuthLogin(email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userAuth)
}

func updatePassord(ctx *gin.Context) {
	user := model.Customer{}
	err := ctx.ShouldBindJSON(&user) //รับค่าจาก body
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// email := user.Email
	service := service.NewUsersService(db)
	updatedUser, err := service.UpdatePasswordByEmail(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func insertUser(ctx *gin.Context) {
	service := service.NewUsersService(db)
	user := model.Customer{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	affectedRow, err := service.InsertUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "affectedRow: " + strconv.FormatInt(affectedRow, 10),
	})
}
