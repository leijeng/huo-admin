package service

import (
	"github.com/leijeng/huo-core/core/base"
)

type SysRoleService struct {
	*base.BaseService
}

var SerSysRole = SysRoleService{
	base.NewService("sys"),
}

