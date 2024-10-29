package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/leijeng/huo-core/config"
)

func InitMiddleware(r *gin.Engine, cfg *config.AppCfg) {

	r.Use(LoggerToFile())
	r.Use(SysOperaLogHandler()) // 操作日志

	if cfg.Cors.Enable {
		r.Use(CorsByRules(&cfg.Cors))
	}

	r.Use(CustomError)

	r.Use(ReqId)
}
