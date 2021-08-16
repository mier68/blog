package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部出现错误")
	InvalidParams             = NewError(10000001, "无效的参数")
	NotFound                  = NewError(10000002, "资源找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的key和secret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，token生成失败")
	ToomanyRequests           = NewError(10000007, "请求过多")
)
