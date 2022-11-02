package controllers

import (
	"net/http"
	"rest-api-golang/models"

	"github.com/gin-gonic/gin"
)

func CariBuku(c *gin.Context) {
	var books []models.Buku
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

type InputTambahBuku struct {
	Judul    string `json:"judul" binding:"required"`
	Penulis  string `json:"penulis" binding:"required"`
	Penerbit string `json:"penerbit" binding:"required"`
	Harga    uint64 `json:"harga" binding:"required"`
}

func TambahBuku(c *gin.Context) {
	var input InputTambahBuku
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Buku{Judul: input.Judul, Penulis: input.Penulis, Penerbit: input.Penerbit, Harga: input.Harga}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
