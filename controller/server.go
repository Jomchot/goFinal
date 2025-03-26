package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartServer() {
	viper.SetConfigName("config") //กำหนดชื่อไฟล์
	viper.AddConfigPath(".")      //กำหนดให้ Viper ค้นหาไฟล์คอนฟิกในโฟลเดอร์ปัจจุบัน (. คือโฟลเดอร์เดียวกับที่โปรแกรมกำลังรัน)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", viper.GetString("mysql.dsn"))
	dsn := viper.GetString("mysql.dsn")
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "message API Is Working....")
	})

	// เรียกใช้ controller
	NewUsers(router, db)
	NewProduct(router, db)
	NewCardItem(router, db)

	router.Run()
}
