package controller

import (
	"goFinal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dbCardItem *gorm.DB

func NewCardItem(router *gin.Engine, gormdb *gorm.DB) {
	dbCardItem = gormdb
	product := router.Group("/cardItem")
	{
		product.POST("/", getItem)
	}

}

type Item struct {
	CartID    uint `json:"cart_id"` // ใช้ JSON tags เพื่อให้สามารถ map ค่าใน JSON ได้
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

func getItem(ctx *gin.Context) {
	var item Item
	err := ctx.ShouldBindJSON(&item) //รับค่าจาก body
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := service.NewCardItemService(dbCardItem)
	ProductSearch, err := service.PostCardItem(item.CartID, item.ProductID, item.Quantity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ส่งผลลัพธ์กลับ
	ctx.JSON(http.StatusOK, ProductSearch)
}
