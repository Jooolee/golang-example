package controller

import (
	"net/http"

	"golang-example/pkg"

	"github.com/gin-gonic/gin"
)

// SeckillQueueHandler 处理将秒杀请求添加到队列的 HTTP 请求
func SeckillQueueHandler(c *gin.Context) {
	productID := c.Query("product_id")
	userID := c.Query("user_id")

	if productID == "" || userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing product_id or user_id"})
		return
	}

	message, statusCode, err := pkg.AddSeckillRequestToQueue(productID, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.JSON(statusCode, gin.H{"message": message})
}
