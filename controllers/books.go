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

func TambahBuku(c *gin.Context) {
	input := InputTambahBuku{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Buku{Judul: input.Judul, Penulis: input.Penulis, Penerbit: input.Penerbit}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CariBukuSatuan(c *gin.Context) {
	var book models.Buku

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Buku tidak ditemukan!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBuku(c *gin.Context) {
	// Get model if exist
	var book models.Buku
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	// Validate input
	var input InputUpdateBuku
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func HapusBuku(c *gin.Context) {

	var book models.Buku

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
