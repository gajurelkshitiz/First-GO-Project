package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"context"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	// create a gin router with default middleware
	r := gin.Default()

	// Define a simple GET Endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK,gin.H{
			"message" : "pong",
		})
	})

	// HomePage for minimalistic app
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to minimalistic app.",
		})
	})

	// ** Just for practising: sqlite database connection using - gorm
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// create
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})

	// Read
	// product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)  // find product with integer primary key
	product, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // find product with code D42

	// Update - update product's price to 200
	err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)


	fmt.Println("## Product after read:", product)



	// server starts on localhost:8080 by default
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}