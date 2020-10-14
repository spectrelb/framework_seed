package Response

type MyCode int64

const (
	CodeSuccess         MyCode = 200
	CodeInvalidParams   MyCode = 1001
	CodeServerBusy      MyCode = 1002
	CodeLoginFail       MyCode = 1003
)

var msgFlags = map[MyCode]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeServerBusy:      "服务繁忙",
	CodeLoginFail:		"登录失败",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
