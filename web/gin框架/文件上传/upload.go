package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//创建路由
	r := gin.Default()

	//处理multipart forms提交文件时默认的内存限制是32 MiB
	r.MaxMultipartMemory = 8 //router.MaxMultipartMemory = 8 << 20  // 8 MiB

	//文件上传 单
	r.POST("/upload", func(c *gin.Context) {
		//获取文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "请上传文件"})
			return
		}

		/*fmt.Println("文件名：", file.Filename)

		fmt.Println("文件大小：", file.Size)

		fmt.Println("文件后缀：", path.Ext(file.Filename))

		fmt.Println("文件类型：", file.Header.Get("Content-Type"))*/

		//保存文件路径
		dst := fmt.Sprintf("D:\\GoProject\\goStudy\\web\\gin\\upload\\%s", file.Filename)

		//保存上传文件
		if err = c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "文件保存失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "文件上传成功"})
	})

	//文件上传 多
	r.POST("/uploads", func(c *gin.Context) {
		//获取解析后表单
		if form, err := c.MultipartForm(); err == nil {
			//获取文件
			files, ok := form.File["files"]
			if !ok {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "请上传文件"})
				return
			}

			//循环保存文件
			for _, file := range files {
				//保存文件路径
				dst := fmt.Sprintf("D:\\GoProject\\goStudy\\web\\gin\\upload\\%s", file.Filename)

				//保存上传文件
				if err = c.SaveUploadedFile(file, dst); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"msg": "文件保存失败"})
					return
				}
			}

			c.JSON(http.StatusOK, gin.H{"msg": "文件上传成功"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "系统错误，请稍后再试"})
		}
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
