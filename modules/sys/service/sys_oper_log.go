package service

import "github.com/leijeng/huo-core/core/base"

type SysOperaLogService struct {
	*base.BaseService
}

var SerSysOperaLog = SysOperaLogService{
	base.NewService("sys"),
}
