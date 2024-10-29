package service

import "github.com/leijeng/huo-core/core/base"

type SysUser struct {
	*base.BaseService
}

var SerSysUser = SysUser{
	base.NewService("admin"),
}
