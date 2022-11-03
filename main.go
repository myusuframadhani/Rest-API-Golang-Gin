package main

import (
	"rest-api-golang/controllers"
	"rest-api-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// Find all books
	r.GET("/books", controllers.CariBuku)

	// Add some book
	r.POST("/books", controllers.TambahBuku)

	// Find book by ID
	r.GET("/books/:id", controllers.CariBukuSatuan)

	// Update book
	r.PATCH("/books/:id", controllers.UpdateBuku)

	// Delete a book
	r.DELETE("/books/:id", controllers.HapusBuku)

	r.Run()
}
