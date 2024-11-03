package dto

import (
    "github.com/leijeng/huo-core/core/base"
    "github.com/leijeng/huo-admin/modules/sys/models"
)

type SysUserRoleGetPageReq struct {
	base.ReqPage `query:"-"`
    
}

func (SysUserRoleGetPageReq) TableName() string {
	return models.TBSysUserRole
}


//SysUserRole
type SysUserRoleDto struct {
    
    SysUserId int `json:"sysUserId" form:"sysUserId"` // 
    SysRoleId int `json:"sysRoleId" form:"sysRoleId"` // 
}

