package router

import (
	"github.com/gin-gonic/gin"
	"webDemo/server"
)

func Start() error {
	router := gin.New()
	router.Use(gin.Recovery())

	//添加业务路由，根据需要选择需要的中间件

	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", server.SayHello)

		v1.POST("/hello1", server.SqlI1)
		v1.POST("/hello2", server.SqlI2)
		v1.POST("/hello3", server.SqlI3)
		v1.POST("/hello4", server.SqlI4)
		v1.POST("/hello5", server.SqlI5)
		v1.POST("/hello6", server.SqlI6)
		v1.POST("/hello7", server.SqlI7)
		v1.POST("/hello8", server.SqlI8)
		v1.POST("/hello9", server.SqlI9)
		v1.POST("/hello10", server.SqlI10)
		//v1.POST("/hello11", server.SqlI11)
		//v1.POST("/hello12", server.SqlI12)
		//v1.POST("/hello13", server.SqlI13)
		v1.POST("/hello15", server.SqlI15)
		v1.POST("/hello17", server.SqlI17)

	}

	return router.Run(":8888")
}
