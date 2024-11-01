package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leijeng/huo-admin/common/consts"
	docs "github.com/leijeng/huo-admin/docs"
	"github.com/leijeng/huo-core/core"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	//routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// InitRouter 路由初始化
func InitRouter() {
	r := core.GetGinEngine()
	if core.Cfg.Server.Mode != core.ModeProd.String() {
		fmt.Printf("%s %s  \r\n", docs.SwaggerInfo.Title, docs.SwaggerInfo.Version)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	noCheckRoleRouter(r)
}

// noCheckRoleRouter 无需认证的路由
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group(consts.ApiRoot + "/sys")

	for _, f := range routerNoCheckRole {
		f(v)
	}
}
