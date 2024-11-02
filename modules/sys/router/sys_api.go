package router

import (
	"github.com/leijeng/huo-admin/modules/sys/apis"
	"github.com/leijeng/huo-admin/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysApiRouter)
}

// 默认需登录认证的路由
func registerSysApiRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-api").Use(middleware.AdminJwtHandler()) //.Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysApi.Get)
		r.POST("/create", apis.ApiSysApi.Create)
		r.POST("/update", apis.ApiSysApi.Update)
		r.POST("/page", apis.ApiSysApi.QueryPage)
		r.POST("/del", apis.ApiSysApi.Del)
	}
}