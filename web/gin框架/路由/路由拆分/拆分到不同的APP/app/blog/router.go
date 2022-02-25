package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(r *gin.Engine) {
	blogGroup := r.Group("blog")
	{
		blogGroup.GET("find", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop find"})
		})

		blogGroup.POST("upload", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "shop upload"})
		})
	}
}
