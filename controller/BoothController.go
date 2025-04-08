package controller

import (
	"ApiBooth/config"
	"ApiBooth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBoothSlametRiyadi(c *gin.Context) {
	trans_date := c.Query("trans_date")

	if trans_date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch ID dan Trx Date Harap diisi!!!"})
	}

	var results []model.SlametRiyadi
	if err := config.DB.Where("no_transaksi LIKE ? AND trans_date = ?", "%S035A%", trans_date).Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan!!!"})
		return
	}

	c.JSON(http.StatusOK, results)
	return
}
