package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zyldgd/redis-z/server/constant"
	"github.com/zyldgd/redis-z/server/entity"
	"github.com/zyldgd/redis-z/server/resource"
	"net/http"
	"time"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, entity.Response{
		Code:  0,
		Error: "",
		Data:  "hello World!",
	})
	//c.String(http.StatusOK, "hello World!")
}

func Set(c *gin.Context) {
	response := &entity.Response{
		Code:  0,
		Error: "",
		Data:  nil,
	}
	status := http.StatusOK
	defer c.JSON(status, response)
	// -----------------------------------
	param := &entity.RedisString{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error = err.Error()
		response.Code = constant.CodeParamError
		return
	}

	rdb := resource.GetRedisClient()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := rdb.Set(ctx, param.Key, param.Value, time.Duration(param.EX)*time.Second)
	fmt.Printf("cmd:%+v \n", *cmd)
}

func Get(c *gin.Context) {
	response := &entity.Response{
		Code:  0,
		Error: "",
		Data:  nil,
	}
	status := http.StatusOK
	defer c.JSON(status, response)
	// -----------------------------------
	param := &entity.RedisString{}
	if err := c.ShouldBindJSON(param); err != nil {
		response.Error = err.Error()
		response.Code = constant.CodeParamError
		return
	}

	rdb := resource.GetRedisClient()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := rdb.Get(ctx, param.Key)
	fmt.Printf("cmd:%+v \n", *cmd)
	if cmd.Err() != nil {
		response.Error = cmd.Err().Error()
		response.Code = constant.CodeRedisError
	}

	response.Data = &entity.RedisString{
		Key:   param.Key,
		Value: cmd.Val(),
	}
}
