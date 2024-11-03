package models

import (
)

//SysUserRole
type SysUserRole struct {
    
    SysUserId int `json:"sysUserId" gorm:"type:int unsigned;comment:SysUserId"` // 
    SysRoleId int `json:"sysRoleId" gorm:"type:int unsigned;comment:SysRoleId"` // 
}

const TBSysUserRole = "sys_user_role"

func (SysUserRole) TableName() string {
    return TBSysUserRole
}

func NewSysUserRole() *SysUserRole{
    return &SysUserRole{}
}

