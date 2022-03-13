package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// ErrorUsernameUsed code = 1000 用户模块err
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExists  = 1003
	ErrorTokenNotExist  = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007

	//code = 2000 文章模块err
	//code = 3000 分类模块err

)

var CodeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	//code = 1000 用户模块err
	ErrorUsernameUsed:   "用户名已被使用",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExists:  "用户不存在",
	ErrorTokenNotExist:  "token不存在",
	ErrorTokenRuntime:   "token过期",
	ErrorTokenWrong:     "token错误",
	ErrorTokenTypeWrong: "token未过期",

	//code = 2000 文章模块err
	//code = 3000 分类模块err
}

func GetErrorMsg(code int) string {
	return CodeMsg[code]
}
