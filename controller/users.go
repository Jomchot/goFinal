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
		user.GET("/id", getUserById)
		user.GET("/email", getUserByEmail)
		user.POST("", insertUser)
		user.PATCH("", updateUserById)
		user.PATCH("/email", updateUserByEmail)
	}
}

func getAllUsers(ctx *gin.Context) {
	service := service.NewUsersService(db)
	getUser := service.GetAllUsers()
	ctx.JSON(http.StatusOK, getUser)
}

func getUserById(ctx *gin.Context) {
	idx := ctx.Query("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	service := service.NewUsersService(db)
	getUser, error := service.GetUserById(id)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, getUser)
}

func getUserByEmail(ctx *gin.Context) {
	service := service.NewUsersService(db)
	email := ctx.Query("email")

	user, err := service.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func updateUserByEmail(ctx *gin.Context) {
	user := model.User{}
	err := ctx.ShouldBindJSON(&user) //รับค่าจาก body
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service := service.NewUsersService(db)
	updatedUser, err := service.UpdateUserByEmail(user.Email, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func updateUserById(ctx *gin.Context) {
	user := model.User{}

	idx := ctx.Query("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}
	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}
	service := service.NewUsersService(db)
	updatedUser, err := service.UpdateUserByID(id, user)
	ctx.JSON(http.StatusOK, updatedUser)
}

func insertUser(ctx *gin.Context) {
	service := service.NewUsersService(db)
	user := model.User{}
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
