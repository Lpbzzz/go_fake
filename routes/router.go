package routes

import (
	"fake_twitter/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	fmt.Println(utils.HttpPort)
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "hello world",
			})
		})
	}
	r.Run(utils.HttpPort)
}
