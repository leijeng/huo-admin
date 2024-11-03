package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leijeng/huo-admin/common/middleware"
	"github.com/leijeng/huo-admin/modules/sys/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysUserRouter)
}

// 默认需登录认证的路由
func registerSysUserRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-user").Use(middleware.AdminJwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysUser.Get)
		r.POST("/create", apis.ApiSysUser.Create)
		r.POST("/update", apis.ApiSysUser.Update)
		r.POST("/page", apis.ApiSysUser.QueryPage)
		r.POST("/del", apis.ApiSysUser.Del)
		r.POST("/logout", apis.ApiSysUser.Logout)
		r.POST("/addRoles", apis.ApiSysUser.AddRoles)
	}
	rv := v1.Group("sys-user")
	{
		rv.POST("/login", apis.ApiSysUser.Login)
	}
}
