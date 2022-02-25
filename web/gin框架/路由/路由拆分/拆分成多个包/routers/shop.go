package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadShop(r *gin.Engine) {
	shopGroup := r.Group("shop")
	{
		shopGroup.GET("find", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop find"})
		})

		shopGroup.POST("upload", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop upload"})
		})
	}
}
