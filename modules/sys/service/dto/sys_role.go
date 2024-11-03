package dto

import (
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-core/core/base"
)

type SysRoleGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int    `json:"status" form:"status" query:"column:status"` //状态：1 有效 2 无效

}

func (SysRoleGetPageReq) TableName() string {
	return models.TBSysRole
}

// SysRole
type SysRoleDto struct {
	Id        int    `json:"id"`                         //主键
	Name      string `json:"name" form:"name"`           //角色名称
	RoleKey   string `json:"roleKey" form:"roleKey"`     //角色代码
	RoleSort  int    `json:"roleSort" form:"roleSort"`   //排序
	Status    int    `json:"status" form:"status"`       //状态：1 有效 2 无效
	TeamId    int    `json:"teamId" form:"teamId"`       //团队id
	Remark    string `json:"remark" form:"remark"`       //备注
	CreatedBy int    `json:"createdBy" form:"createdBy"` //创建人id
	UpdatedBy int    `json:"updatedBy" form:"updatedBy"` //修改人id
}

type AddRoleMenuDto struct {
	RoleId  int   `json:"roleId"`
	MenuIds []int `json:"menuIds"`
}
