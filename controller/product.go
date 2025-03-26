package controller

import (
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
		product.GET("/search", getProduct)
	}

}

type Search struct {
	Description string  `json:"description"`
	MinPrice    float64 `json:"minPrice"`
	MaxPrice    float64 `json:"maxPrice"`
}

func getProduct(ctx *gin.Context) {
	var search Search
	err := ctx.ShouldBindJSON(&search) //รับค่าจาก body
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := service.NewProductService(dbProduct)
	ProductSearch, err := service.GetProduct(search.Description, search.MinPrice, search.MaxPrice)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ProductSearch)
}
