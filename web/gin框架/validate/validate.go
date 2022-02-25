package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type Person struct {
	Age      int       `form:"age" json:"age" binding:"required,gt=10"`
	Name     string    `form:"name" json:"name" binding:"required"`
	Birthday time.Time `form:"birthday" json:"birthday" binding:"timing" time_format:"2006-01-02" time_utc:"1"`
}

var trans ut.Translator

// InitTrans 初始化一个翻译器函数
func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎属性，实现定制
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册一个获取json的tag自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
	zhT := zh.New()              // 中文翻译器
	enT := en.New()              // 英文翻译器
	uni := ut.New(enT, zhT, enT) // 配置默认翻译器、以及可支持的翻译器
	trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) error", locale)
	}
	switch locale {
	case "en":
		_ = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		_ = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		_ = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	return
}

// RemoveTopStruct 去掉结构体名称前缀
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// 自定义验证 时间必须小于当前时间
func timing(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok && time.Now().After(date) {
		return true
	}
	return false
}

func main() {
	//创建路由
	r := gin.Default()

	// 注册验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("timing", timing); err != nil {
			log.Fatalln(err)
			return
		}
	}

	// 调用翻译器函数
	if err := InitTrans("zh"); err != nil {
		log.Fatalln(err)
		return
	}

	// post请求
	r.POST("validate", func(c *gin.Context) {
		//申明变量类型
		var person Person

		// 绑定参数
		if err := c.ShouldBind(&person); err != nil {
			// 如果错误类型为validate错误则翻译错误信息
			if vErr, ok := err.(validator.ValidationErrors); !ok {
				// 如果错误不能转化，不能进行翻译，就返回错误信息
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				// 如果错误能进行翻译，就返回翻译后的错误信息
				c.JSON(http.StatusBadRequest, gin.H{"error": RemoveTopStruct(vErr.Translate(trans))})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"person：": person})
	})

	//监听端口，默认为8080
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err)
	}
}
