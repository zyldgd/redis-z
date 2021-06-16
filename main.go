package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zyldgd/redis-z/server/handler"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/hw", handler.HelloWorld)
	redisGroup := r.Group("/redis")
	{
		redisGroup.POST("/set", handler.Set)
		redisGroup.GET("/get", handler.Get)
	}

	if err := r.Run(":8000"); err != nil {
		fmt.Printf("server run err:%+v \n", err)
	}
}
