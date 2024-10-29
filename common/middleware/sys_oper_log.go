package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/leijeng/huo-admin/common/utils"
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service"
	"github.com/leijeng/huo-core/core"
	"github.com/leijeng/huo-core/core/base"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"time"
)

type LogResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w LogResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w LogResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func SysOperaLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// 处理请求
		var body string

		lw := &LogResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		if c.Request.Method != http.MethodOptions {
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
				rb, err := io.ReadAll(bf)
				if err != nil {
					core.Log.Warn("ReadAll body error", zap.Error(err))
					err = nil
				}
				log.Println("body2", string(rb))
				c.Request.Body = io.NopCloser(bytes.NewBuffer(rb))
				body = string(rb)
			default:
				body = c.Request.URL.RawQuery
			}

			c.Writer = lw
		}

		c.Next()

		if c.Request.Method == http.MethodOptions {
			return
		}
		cost := time.Since(startTime)
		// 记录日志
		var operlog models.SysOperaLog
		operlog.Title = "-"
		operlog.Status = c.Writer.Status()
		operlog.RequestMethod = c.Request.Method
		operlog.OperUrl = c.Request.URL.Path
		operlog.OperParam = body
		body = lw.body.String()
		if len(body) > 1024 {
			res := base.Resp{}
			json.Unmarshal([]byte(body), &res)
			res.Data = "太长了，只记录前1024个字符"
			b, _ := json.Marshal(&res)
			body = string(b)
		}
		operlog.JsonResult = body
		operlog.OperIp = c.ClientIP()
		operlog.OperName = utils.GetNickname(c)
		operlog.OperTime = utils.LocalTime(startTime)
		operlog.LatencyTime = cost.String()
		operlog.UserAgent = c.Request.UserAgent()
		operlog.CreateBy = c.GetInt("a_userid")

		log.Println("------------operlog------------", operlog.CreateBy)
		apis := GetApis()
		for _, api := range apis {
			if api.Method == operlog.RequestMethod && KeyMatch2(operlog.OperUrl, api.Path) {
				operlog.Title = api.Title
				break
			}
		}

		if operlog.CreateBy > 0 {
			err := service.SerSysOperaLog.Create(&operlog)
			if err != nil {
				core.Log.Error("copy body error", zap.Error(err))
			}
		}

	}
}
