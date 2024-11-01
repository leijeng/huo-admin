package models

import (
            "gorm.io/gorm"
    "github.com/leijeng/huo-admin/common/utils"
)

//SysRole
type SysRole struct {
    
    Id int `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
    Name string `json:"name" gorm:"type:varchar(255);comment:角色名称"` //角色名称 
    RoleKey string `json:"roleKey" gorm:"type:varchar(255);comment:角色代码"` //角色代码 
    RoleSort int `json:"roleSort" gorm:"type:int;comment:排序"` //排序 
    Status int `json:"status" gorm:"type:tinyint(1);comment:状态：1 有效 2 无效"` //状态：1 有效 2 无效 
    TeamId int `json:"teamId" gorm:"type:int;comment:团队id"` //团队id 
    Remark string `json:"remark" gorm:"type:varchar(255);comment:备注"` //备注 
    CreatedBy int `json:"createdBy" gorm:"type:int;comment:创建人id"` //创建人id 
    UpdatedBy int `json:"updatedBy" gorm:"type:int;comment:修改人id"` //修改人id 
    CreatedAt utils.LocalTime `json:"createdAt" gorm:"type:datetime;comment:CreatedAt"` // 
    UpdatedAt utils.LocalTime `json:"updatedAt" gorm:"type:datetime;comment:UpdatedAt"` // 
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`     //删除时间
}

const TBSysRole = "sys_role"

func (SysRole) TableName() string {
    return TBSysRole
}

func NewSysRole() *SysRole{
    return &SysRole{}
}

