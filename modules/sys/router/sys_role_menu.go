package router

import (
	"github.com/leijeng/huo-admin/modules/sys/apis"
	"github.com/leijeng/huo-admin/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysRoleMenuRouter)
}

// 默认需登录认证的路由
func registerSysRoleMenuRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-role-menu").Use(middleware.AdminJwtHandler()) //.Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysRoleMenu.Get)
		r.POST("/create", apis.ApiSysRoleMenu.Create)
		r.POST("/update", apis.ApiSysRoleMenu.Update)
		r.POST("/page", apis.ApiSysRoleMenu.QueryPage)
		r.POST("/del", apis.ApiSysRoleMenu.Del)
	}
}