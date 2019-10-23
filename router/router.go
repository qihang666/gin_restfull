package router

import (
	"fmt"
	"gin_restfull/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine  {
	r := gin.Default()
	//rGroup := r.Group("")
	//验证权限中间件
	r.Use(CheckToken())
	//swagger 配置相关
	url := ginSwagger.URL("http://localhost:9000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//用户模块路由
	r.GET("/user",service.Get)
	r.POST("/user",service.Add)
	r.PUT("/user",service.Update)
	r.DELETE("/user",service.Delete)

    r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,"hello")
	})
    return r
}

func CheckToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		fmt.Printf("%s::%s \t %s \t %s =====", time.Now().Format("2006-01-02 15:04:05"), host, url, method)
		context.Next()
		fmt.Println(context.Writer.Status())
	}
}


