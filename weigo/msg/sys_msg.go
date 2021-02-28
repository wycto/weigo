package msg

import "wycto/weigo"

var (
	//参数校验
	ParamsMissing          = &weigo.Message{Code: 200101, Msg: "参数缺失"}
	ParamsInvalid          = &weigo.Message{Code: 200102, Msg: "参数无效"}
	ParamsMissingORInvalid = &weigo.Message{Code: 200102, Msg: "参数缺失或无效"}

	//数据库信息信息
	DataNone             = &weigo.Message{Code: 200201, Msg: "数据不存在"}
	DataExists           = &weigo.Message{Code: 200202, Msg: "数据已存在"}
	DataAccountExists    = &weigo.Message{Code: 200203, Msg: "账户已存在"}
	DataAccountNotExists = &weigo.Message{Code: 200204, Msg: "账户不存在"}
	DataPasswordError    = &weigo.Message{Code: 200204, Msg: "密码错误"}

	//业务逻辑信息

	//系统错误
	SysError = &weigo.Message{Code: 200501, Msg: "系统错误"}
)
