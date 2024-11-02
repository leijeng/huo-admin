package router

import (
	"github.com/leijeng/huo-admin/modules/sys/apis"
	"github.com/leijeng/huo-admin/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysMenuRouter)
}

// 默认需登录认证的路由
func registerSysMenuRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-menu").Use(middleware.AdminJwtHandler()) //.Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysMenu.Get)
		r.POST("/create", apis.ApiSysMenu.Create)
		r.POST("/update", apis.ApiSysMenu.Update)
		r.POST("/page", apis.ApiSysMenu.QueryPage)
		r.POST("/del", apis.ApiSysMenu.Del)
	}
}