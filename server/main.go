package main

import (
	"main/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dns := "root:@tcp(127.0.0.1:3306)/news?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.News{})
	if err != nil {
		panic("failed to migrate database")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

func main() {
	db := GetDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, // Substitua PORTA_DO_CLIENTE pela porta do seu cliente Svelte
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		var news []models.News
		db.Find(&news)
		c.JSON(200, news)
	})

	r.POST("/news", func(c *gin.Context) {
		var news models.News
		c.BindJSON(&news)
		db.Create(&news)
		c.JSON(200, news)
	})

	r.Run()
}
