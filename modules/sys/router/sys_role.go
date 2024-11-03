package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leijeng/huo-admin/common/middleware"
	"github.com/leijeng/huo-admin/modules/sys/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysRoleRouter)
}

// 默认需登录认证的路由
func registerSysRoleRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-role").Use(middleware.AdminJwtHandler()) //.Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysRole.Get)
		r.POST("/create", apis.ApiSysRole.Create)
		r.POST("/update", apis.ApiSysRole.Update)
		r.POST("/page", apis.ApiSysRole.QueryPage)
		r.POST("/del", apis.ApiSysRole.Del)
		r.POST("/addMenus", apis.ApiSysRole.AddMenus)
	}
}
