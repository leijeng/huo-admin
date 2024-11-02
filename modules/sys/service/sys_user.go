package service

import (
	"errors"
	"github.com/leijeng/huo-admin/common/codes"
	"github.com/leijeng/huo-admin/common/utils"
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service/dto"
	"github.com/leijeng/huo-core/core"
	"github.com/leijeng/huo-core/core/base"
	"github.com/leijeng/huo-core/core/errs"
	"time"
)

type SysUser struct {
	*base.BaseService
}

var SerSysUser = SysUser{
	base.NewService("sys"),
}

func (e *SysUser) LoginPwd(req dto.SysUserLoginReq, userData models.SysUser) (res dto.LoginOK, err error) {

	if !userData.CompPwd(req.Password) {
		return res, errors.New("password error")
	}

	return e.generateToken(userData)
}

func (e *SysUser) generateToken(userData models.SysUser) (res dto.LoginOK, err error) {

	exp := time.Now().Add(time.Duration(core.Cfg.JWT.Expires) * time.Minute)
	claims := utils.NewClaims(userData.Id, exp, core.Cfg.JWT.Issuer, core.Cfg.JWT.Subject)
	claims.Phone = userData.Phone
	claims.Nickname = userData.Nickname

	token, err := utils.Generate(claims, core.Cfg.JWT.SignKey)

	lok := dto.LoginOK{}

	if err != nil {
		return lok, errs.Err(codes.FAILURE, "", err)
	}

	lok.Expire = utils.LocalTime(exp)
	lok.AccessToken = token

	if userData.Nickname != "" {
		lok.Username = userData.Nickname
	} else if userData.Username != "" {
		lok.Username = userData.Username
	} else if userData.Phone != "" {
		lok.Username = userData.Phone
	} else if userData.Email != "" {
		lok.Username = userData.Email
	}

	claims.ExpiresAt(exp.Add(time.Hour * 24 * 7))
	refT, _ := utils.Generate(claims, core.Cfg.JWT.SignKey)
	lok.RefreshToken = refT
	return lok, nil
}
