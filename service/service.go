package service

import (
	"gin_restfull/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Result struct {
	Code int
	Message string
    Data  interface{}
}

var  FailResult Result = Result{Code:2,Message:"错误",Data:nil}

// @Summary 删除用户信息
// @Id 1
// @Tags User
// @version 1.0
// @Accept application/x-json-stream
// @Param user body model.User true "用户"
// @Success 200 object service.Result 成功后返回值
// @Router /user [delete]
func Delete(context *gin.Context) {
	var use model.User
	if err := context.ShouldBindJSON(&use);err==nil{
		log.Println(use)
		use.DeleteUser()
		context.JSON(http.StatusOK,use)
	}else{
		//log.Println(err.Error())
		log.Println(use)
		context.JSON(http.StatusBadRequest,FailResult)
	}
}
// @Summary 新增用户信息
// @Id 1
// @Tags User
// @version 1.0
// @Accept application/x-json-stream
// @Param user body model.User true "用户"
// @Success 200 object service.Result 成功后返回值
// @Router /user [post]
func Add(context *gin.Context) {
	var use model.User
	if err := context.ShouldBindJSON(&use);err==nil{
		log.Println(use)
		use.AddUser()
		context.JSON(http.StatusOK,use)
	}else{
		log.Println(use)
		context.JSON(http.StatusBadRequest,FailResult)
	}
}
// @Summary 更新用户信息
// @Id 1
// @Tags User
// @version 1.0
// @Accept application/x-json-stream
// @Param user body model.User true "用户"
// @Success 200 object service.Result 成功后返回值
// @Router /user [put]
func Update(c *gin.Context)  {
	var use model.User
	if err := c.ShouldBindJSON(&use);err==nil{
		log.Println(use)
		use.UpdateUser()
		c.JSON(http.StatusOK,use)
	}else{
		//log.Println(err.Error())
		log.Println(use)
		c.JSON(http.StatusBadRequest,FailResult)
	}
}
// @Summary 获取用户信息
// @Id 1
// @Tags User
// @version 1.0
// @Accept application/x-json-stream
// @Param user body model.User true "用户"
// @Success 200 object service.Result 成功后返回值
// @Router /user [get]
func Get(c *gin.Context)  {
	var use model.User
	if err := c.ShouldBindJSON(&use);err!=nil{
		c.JSON(http.StatusBadRequest,Result{2,"参数错误",err.Error()})
		return
	}
	u,err :=use.GetUser()
	if err!=nil{
		c.JSON(http.StatusBadRequest,Result{2,"参数错误",err.Error()})
		return
	}
	c.JSON(http.StatusOK,u)

}