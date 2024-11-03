package router

import (
	"github.com/leijeng/huo-admin/modules/sys/apis"
	"github.com/leijeng/huo-admin/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysUserRoleRouter)
}

// 默认需登录认证的路由
func registerSysUserRoleRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-user-role").Use(middleware.AdminJwtHandler()) //.Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysUserRole.Get)
		r.POST("/create", apis.ApiSysUserRole.Create)
		r.POST("/update", apis.ApiSysUserRole.Update)
		r.POST("/page", apis.ApiSysUserRole.QueryPage)
		r.POST("/del", apis.ApiSysUserRole.Del)
	}
}