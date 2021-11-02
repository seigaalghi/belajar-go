package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/belajar-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()

	var tulisan []models.Buah

	tulisan = append(tulisan, models.Buah{
		Nama: "Nanas",
	})
	tulisan = append(tulisan, models.Buah{
		Nama: "Mangga",
	})

	fmt.Println(tulisan)

	dsn := "root:@tcp(127.0.0.1:3306)/coba_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	server.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	hasil, _, _ := sum(10, 5)
	fmt.Println(hasil)

	server.Run(":5000")
}

func sum(a int, b int) (int, string, bool) {
	return a + b, "oke", true
}
