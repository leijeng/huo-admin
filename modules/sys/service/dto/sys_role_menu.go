package dto

import (
    "github.com/leijeng/huo-core/core/base"
    "github.com/leijeng/huo-admin/modules/sys/models"
)

type SysRoleMenuGetPageReq struct {
	base.ReqPage `query:"-"`
    
}

func (SysRoleMenuGetPageReq) TableName() string {
	return models.TBSysRoleMenu
}


//SysRoleMenu
type SysRoleMenuDto struct {
    
    SysRoleId int `json:"sysRoleId" form:"sysRoleId"` // 
    SysMenuId int `json:"sysMenuId" form:"sysMenuId"` // 
}

