package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leijeng/huo-admin/common/middleware"
	"github.com/leijeng/huo-admin/modules/sys/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysMenuApiRouter)
}

// 默认需登录认证的路由
func registerSysMenuApiRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-menu-api").Use(middleware.AdminJwtHandler()) //.Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysMenuApi.Get)
		r.POST("/create", apis.ApiSysMenuApi.Create)
		r.POST("/update", apis.ApiSysMenuApi.Update)
		r.POST("/page", apis.ApiSysMenuApi.QueryPage)
		r.POST("/del", apis.ApiSysMenuApi.Del)
	}
}
