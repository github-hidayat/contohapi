package controllers

import (
	"net/http"
	"pustaka-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MahasiswaInput struct {
	Nim  string `json:"nim"`
	Nama string `json:"nama"`
}

//Tampil data mahasiswa
func MahasiswaTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{
		"data":    mhs,
		"message": "data ok",
	})
}

// Tambah data mahasiswa
func MahasiswaTambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi input/masukkan
	var dataInput MahasiswaInput

	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses input
	mhs := models.Mahasiswa{
		Nim:  dataInput.Nim,
		Nama: dataInput.Nama,
	}

	db.Create(&mhs)

	c.JSON(http.StatusOK, gin.H{
		"data": mhs,
		"message": "data add ok",
		})
}

// Ubah data mahasiswa
func MahasiswaUbah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var mhs models.Mahasiswa
	if(c.Param("nim") == ""){
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Nim harus diisi"})
		return
	}
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Data mahasiswa tidak ditemukan"})
		return
	}

	//validasi input/masukkan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses ubah data
	db.Model(&mhs).Update(dataInput)

	c.JSON(http.StatusOK, gin.H{
		"data": mhs,
		"message": "data edit ok",
		})
}

// Hapus data mahasiswa
func MahasiswaHapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek dulu datanya
	var mhs models.Mahasiswa
	if(c.Param("nim") == ""){
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Nim harus diisi"})
		return
	}

	if  err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Data mahasiswa tidak ditemukan"})
		return
	}

	//proses hapus data
	db.Delete(&mhs)

	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"message": "data delete ok",
		})
}
