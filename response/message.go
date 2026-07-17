package response

// Message 系统消息
type Message struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	// 参数校验
	ParamsMissing          = &Message{Code: 200101, Msg: "参数缺失"}
	ParamsInvalid          = &Message{Code: 200102, Msg: "参数无效"}
	ParamsMissingORInvalid = &Message{Code: 200103, Msg: "参数缺失或无效"}

	// 数据相关
	DataNone             = &Message{Code: 200201, Msg: "数据不存在"}
	DataExists           = &Message{Code: 200202, Msg: "数据已存在"}
	DataAccountExists    = &Message{Code: 200203, Msg: "账户已存在"}
	DataAccountNotExists = &Message{Code: 200204, Msg: "账户不存在"}
	DataPasswordError    = &Message{Code: 200205, Msg: "密码错误"}

	// 系统错误
	SysError = &Message{Code: 200501, Msg: "系统错误"}
)
