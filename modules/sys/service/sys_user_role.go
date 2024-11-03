package service

import (
	"github.com/leijeng/huo-core/core/base"
)

type SysUserRoleService struct {
	*base.BaseService
}

var SerSysUserRole = SysUserRoleService{
	base.NewService("sys"),
}

