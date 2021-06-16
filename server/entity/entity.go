package entity

import "github.com/zyldgd/redis-z/server/constant"

type Response struct {
	Code  constant.ResponseCode `json:"code"`
	Error string                `json:"error"`
	Data  interface{}           `json:"data"`
}

// -------------------------------------------------

type RedisString struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	EX    int64  `json:"ex"`
}
