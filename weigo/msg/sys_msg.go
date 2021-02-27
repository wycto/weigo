package msg

import "wycto/weigo"

var (
	//参数校验
	ParamsMissing = &weigo.Message{Code: 200101, Msg: "参数缺失"}
	ParamsInvalid = &weigo.Message{Code: 200102, Msg: "参数无效"}

	//数据库信息信息

	//业务逻辑信息
)
