package models

import (
)

//SysRoleMenu
type SysRoleMenu struct {
    
    SysRoleId int `json:"sysRoleId" gorm:"type:int unsigned;comment:SysRoleId"` // 
    SysMenuId int `json:"sysMenuId" gorm:"type:int unsigned;comment:SysMenuId"` // 
}

const TBSysRoleMenu = "sys_role_menu"

func (SysRoleMenu) TableName() string {
    return TBSysRoleMenu
}

func NewSysRoleMenu() *SysRoleMenu{
    return &SysRoleMenu{}
}

