package models

import (
)

//SysMenuApi
type SysMenuApi struct {
    
    SysMenuId int `json:"sysMenuId" gorm:"type:int unsigned;comment:SysMenuId"` // 
    SysApiId int `json:"sysApiId" gorm:"type:int unsigned;comment:SysApiId"` // 
}

const TBSysMenuApi = "sys_menu_api"

func (SysMenuApi) TableName() string {
    return TBSysMenuApi
}

func NewSysMenuApi() *SysMenuApi{
    return &SysMenuApi{}
}

