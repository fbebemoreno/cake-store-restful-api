package cakecontroller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fbebemoreno/cake-store-restful-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var cakes []models.Cake

	models.DB.Order("title, rating desc").Find(&cakes)
	c.JSON(http.StatusOK, gin.H{"data": cakes})
}

func Show(c *gin.Context) {
	var cake models.Cake
	id := c.Param("id")

	if err := models.DB.First(&cake, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": cake})
}

func Create(c *gin.Context) {
	var cake models.Cake

	if err := c.ShouldBindJSON(&cake); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&cake)
	c.JSON(http.StatusOK, gin.H{"data": cake})
}

func Update(c *gin.Context) {
	var cake models.Cake
	id := c.Param("id")

	if err := c.ShouldBindJSON(&cake); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&cake).Where("id = ?", id).Updates(&cake).UpdateColumns(models.Cake{Updated_At: time.Now(), Deleted_At: time.Now()}).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diperbarui!"})
}

func Delete(c *gin.Context) {
	var cake models.Cake
	id := c.Param("id")

	_id, _ := strconv.ParseInt(id, 10, 64)
	if models.DB.Delete(&cake, _id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat menghapus data!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus!"})
}
