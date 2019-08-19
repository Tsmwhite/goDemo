package main

import (
	"github.com/gin-gonic/gin"
	"z3/controllers"
	//"net/http"
)

func main (){
	router := gin.Default()
	router.NoRoute(Handle404)
	router.POST("/login",controllers.LoginAuth)
	router.POST("/register",controllers.Register)

	userCenter := router.Group("/user")
	userCenter.Use(AuthVerify())
	{
		userCenter.GET("/info")
		userCenter.POST("/headimg",controllers.ChangeHeadimg)
		userCenter.POST("/password",controllers.ChangePassword)
		userCenter.POST("/changeinfo")
	}

	router.Run(":8080")
}

func Handle404(c *gin.Context) {
	controllers.ResError("Not Fund 404",c)
}

//登录权限验证
func AuthVerify() gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {
			if catch := recover();catch != nil{
				controllers.ResError(catch,c)
			}
		}()
		header  := c.Request.Header
		if header == nil || header["User_key"] == nil || header["User_key"][0] == ""{
			panic("暂无权限进行此操作")
		}
		err := controllers.ResUserKeyMsg(header["User_key"][0])
		if err != ""{
			panic(err)
		}
	}
}