package dto

import (
    "github.com/leijeng/huo-core/core/base"
    "github.com/leijeng/huo-admin/modules/sys/models"
)

type SysMenuApiGetPageReq struct {
	base.ReqPage `query:"-"`
    
}

func (SysMenuApiGetPageReq) TableName() string {
	return models.TBSysMenuApi
}


//SysMenuApi
type SysMenuApiDto struct {
    
    SysMenuId int `json:"sysMenuId" form:"sysMenuId"` // 
    SysApiId int `json:"sysApiId" form:"sysApiId"` // 
}

