package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mpay/controllers/payments"
	"github.com/mpay/database"
	"github.com/mpay/models"

	_ "github.com/go-sql-driver/mysql"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Payment{})
}

func main() {
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	db := database.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	v1 := r.Group("/payments")
	payments.Router(v1)
	testAuth := r.Group("/api/ping")
	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
