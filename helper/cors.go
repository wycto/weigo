package helper

import "github.com/wycto/weigo/config"

// IsOriginAllowed 检查 origin 是否在允许列表中
// Debug 模式且未配置 AllowOrigins 时允许所有来源
func IsOriginAllowed(origin string) bool {
	cfg := config.GlobalConfig
	if cfg == nil {
		return false
	}

	if cfg.App.Debug && len(cfg.App.AllowOrigins) == 0 {
		return true
	}

	if len(cfg.App.AllowOrigins) == 0 {
		return false
	}

	for _, allowed := range cfg.App.AllowOrigins {
		if allowed == "*" || allowed == origin {
			return true
		}
	}
	return false
}
