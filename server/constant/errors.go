package constant

type ResponseCode int

const (
	CodeParamError ResponseCode = 101
	CodeRedisError ResponseCode = 102
)
