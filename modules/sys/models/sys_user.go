package models

import (
	"github.com/leijeng/huo-admin/common/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	Id        int             `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Username  string          `json:"username" gorm:"type:varchar(32);comment:用户名"`                    //用户名
	Phone     string          `json:"phone" gorm:"type:varchar(11);comment:手机号"`                       //手机号
	Email     string          `json:"email" gorm:"type:varchar(128);comment:邮箱"`                       //邮箱
	Password  string          `json:"-" gorm:"type:varchar(128);comment:密码"`                           //密码
	Nickname  string          `json:"nickname" gorm:"type:varchar(128);comment:昵称"`                    //昵称
	Avatar    string          `json:"avatar" gorm:"type:varchar(255);comment:头像"`                      //头像
	Bio       string          `json:"bio" gorm:"type:varchar(255);comment:签名"`                         //签名
	Birthday  utils.LocalDate `json:"birthday" gorm:"type:date;comment:生日 格式 yyyy-MM-dd"`              //生日 格式 yyyy-MM-dd
	Gender    int             `json:"gender" gorm:"type:tinyint(1);comment:性别 1男 2女 3未知"`              //性别 1男 2女 3未知
	RoleId    int             `json:"roleId" gorm:"type:mediumint;comment:角色ID"`                       //角色ID
	Remark    string          `json:"remark" gorm:"type:varchar(255);comment:备注"`                      //备注
	LockTime  utils.LocalTime `json:"lockTime" gorm:"type:datetime(3);comment:锁定结束时间"`                 //锁定结束时间
	Status    int             `json:"status" gorm:"type:tinyint;comment:状态 1正常 "`                      //状态 1正常
	UpdateBy  int             `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	CreatedAt utils.LocalTime `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                  //创建时间
	UpdatedAt utils.LocalTime `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                //最后更新时间
	IsLogin   int             `json:"isLogin" gorm:"type:tinyint;comment:是否允许登录 1 允许 2 禁止"`            //是否允许登录
}

const TBSysUser = "sys_user"

func (SysUser) TableName() string {
	return TBSysUser
}

func NewSysUser() *SysUser {
	return &SysUser{}
}

// 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) GenPwd(pwd string) (enPwd string, err error) {
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return
	} else {
		enPwd = string(hash)
	}
	return
}

func (e *SysUser) CompPwd(srcPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(srcPwd)); err != nil {
		return false
	}
	return true
}
