package main

import (
	"rest-api-golang/controllers"
	"rest-api-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.CariBuku)
	r.POST("/books", controllers.TambahBuku)

	r.Run()
}
