package dto

import (
	"github.com/leijeng/huo-admin/common/utils"
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-core/core/base"
)

type SysUserGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int    `json:"status" form:"status" query:"column:status"` //

}

func (SysUserGetPageReq) TableName() string {
	return models.TBSysUser
}

// SysUser
type SysUserDto struct {
	Id       int             `json:"id"`                       //主键
	Username string          `json:"username" form:"username"` //
	Phone    string          `json:"phone" form:"phone"`       //
	Email    string          `json:"email" form:"email"`       //
	Password string          `json:"password" form:"password"` //
	Nickname string          `json:"nickname" form:"nickname"` //
	Avatar   string          `json:"avatar" form:"avatar"`     //
	Bio      string          `json:"bio" form:"bio"`           //
	Birthday utils.LocalTime `json:"birthday" form:"birthday"` //
	Gender   int             `json:"gender" form:"gender"`     //
	RoleId   int             `json:"roleId" form:"roleId"`     //
	Remark   string          `json:"remark" form:"remark"`     //
	LockTime utils.LocalTime `json:"lockTime" form:"lockTime"` //
	Status   int             `json:"status" form:"status"`     //
	IsLogin  int             `json:"isLogin" form:"isLogin"`   //
}

type SysUserLoginReq struct {
	Username string `json:"username" form:"username"` //
	Password string `json:"password" form:"password"` //
}

type LoginOK struct {
	AccessToken  string          `json:"accessToken"`  //返回token
	Expire       utils.LocalTime `json:"expires"`      //有效期
	Username     string          `json:"username"`     //token有效期
	Roles        []string        `json:"roles"`        //角色
	RefreshToken string          `json:"refreshToken"` //刷新token
}
