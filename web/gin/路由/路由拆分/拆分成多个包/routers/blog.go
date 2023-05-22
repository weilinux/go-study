package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadBlog(r *gin.Engine) {
	blogGroup := r.Group("blog")
	{
		blogGroup.GET("find", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "blog find"})
		})

		blogGroup.POST("upload", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "blog upload"})
		})
	}
}
