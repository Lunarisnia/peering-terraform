package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Reply struct {
	ID   uint `gorm:"primaryKey;default;auto_random()`
	Text string
}

func main() {
	dsn := "host=gopinger-database-2.cihdebq7rpev.ap-southeast-1.rds.amazonaws.com user=postgres dbname=gopinger password=password1234 port=5432"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Reply{})

	defaultReply := &Reply{Text: "Pong"}
	db.Create(&defaultReply)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pang",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		result := &Reply{}
		db.First(&result)
		if result != nil {
			c.JSON(200, gin.H{
				"reply": result.Text,
			})
		} else {
			c.JSON(404, gin.H{
				"reply": "NOT FOUND!",
			})
		}
	})
	r.Run()
}
