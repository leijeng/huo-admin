package middleware

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/leijeng/huo-core/common/utils"
	"github.com/leijeng/huo-core/common/utils/ips"
	"github.com/leijeng/huo-core/core"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		var body string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			bf := bytes.NewBuffer(nil)
			wt := bufio.NewWriter(bf)
			_, err := io.Copy(wt, c.Request.Body)
			if err != nil {
				core.Log.Warn("copy body error", zap.Error(err))
				err = nil
			}
			wt.Flush()
			rb, _ := io.ReadAll(bf)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(rb))
			body = string(rb)
		}

		c.Next()

		writeLog(startTime, body, c)
	}
}

func writeLog(startTime time.Time, body string, c *gin.Context) {
	// 结束时间
	if c.Request.Method == http.MethodOptions {
		return
	}
	cost := time.Since(startTime)
	if cost.Milliseconds() < 200 {
		core.Log.Info("request", zap.String("ip", ips.GetIP(c)), zap.String("method", c.Request.Method), zap.String("path", c.Request.RequestURI),
			zap.Duration("cost", cost), zap.String("userAgent", c.Request.UserAgent()), zap.String("query", c.Request.URL.RawQuery),
			zap.String("body", body), zap.String("source", core.Cfg.Server.Name), zap.String("reqId", utils.GetReqId(c)))
		//,zap.String("error", strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n")))
	} else {
		core.Log.Warn("request", zap.String("ip", ips.GetIP(c)), zap.String("method", c.Request.Method), zap.String("path", c.Request.RequestURI),
			zap.Duration("cost", cost), zap.String("userAgent", c.Request.UserAgent()), zap.String("query", c.Request.URL.RawQuery),
			zap.String("body", body), zap.String("source", core.Cfg.Server.Name), zap.String("reqId", utils.GetReqId(c)))
		//,zap.String("error", strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n")))
	}
}
