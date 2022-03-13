package routes

import (
	v1 "fake_twitter/api/v1"
	"fake_twitter/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	fmt.Println(utils.HttpPort)
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")

	{
		//	用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("user/get", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		//	分类模块的路由接口

		//	分类模块的路由接口

		//	文章模块的路由接口

	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
