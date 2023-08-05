package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	user := &User{}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"id":   user.ID,
	})
}

func getData(c *gin.Context) {
	id := c.Query("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "Invalid ID",
		})
		return
	}

	i, _ := strconv.Atoi(id)

	user := &User{UserID: uint(i)}

	db.First(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user.Data,
	})
}

func setData(c *gin.Context) {
	id := c.Query("id")

	jsonData, err := c.GetRawData()
	if err != nil {
		// Handle error
	}

	if len(id) == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "Invalid ID",
		})
		return
	}

	i, _ := strconv.Atoi(id)
	user := &User{UserID: uint(i)}

	db.First(&user)

	user.Data = string(jsonData)

	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
