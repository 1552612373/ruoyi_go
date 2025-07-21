package rescode

type Code int

const (
	Success Code = 0

	// 通用错误码 1xxx
	ErrUnknown        Code = 1000
	ErrInvalidParam   Code = 1001
	ErrNotFound       Code = 1002
	ErrTimeout        Code = 1003
	ErrLimitExceed    Code = 1004
	ErrTooManyRequest Code = 1005

	// 用户相关错误码 2xxx
	ErrCustomize        Code = 2000
	ErrUserNotFound     Code = 2001
	ErrUserAlreadyExist Code = 2002
	ErrUserUnauthorized Code = 2003
	ErrUserForbidden    Code = 2004

	// 系统级错误码 9xxx
	ErrInternalServer Code = 9000
)

var codeTextMap = map[Code]string{
	Success: "成功",

	// 通用错误
	ErrUnknown:        "未知错误",
	ErrInvalidParam:   "参数无效",
	ErrNotFound:       "未找到数据",
	ErrTimeout:        "请求超时",
	ErrLimitExceed:    "超出限制",
	ErrTooManyRequest: "请求过于频繁",

	// 用户相关
	ErrCustomize:        "",
	ErrUserNotFound:     "用户不存在",
	ErrUserAlreadyExist: "用户已存在",
	ErrUserUnauthorized: "登录状态已失效，请重新登录",
	ErrUserForbidden:    "禁止访问",

	// 系统错误
	ErrInternalServer: "系统内部错误",
}

// String 返回错误码对应的字符串描述
func (c Code) String() string {
	msg, ok := codeTextMap[c]
	if !ok {
		return "未知错误"
	}
	return msg
}
