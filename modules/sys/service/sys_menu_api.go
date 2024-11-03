package service

import (
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service/dto"
	"github.com/leijeng/huo-core/core"
	"github.com/leijeng/huo-core/core/base"
)

type SysMenuApiService struct {
	*base.BaseService
}

var SerSysMenuApi = SysMenuApiService{
	base.NewService("sys"),
}

func (s *SysMenuApiService) AddApis(data dto.AddMenuApiDto) error {

	tx := s.DB().Model(&models.SysMenuApi{}).Begin()

	var sysMenuApiData models.SysMenuApi

	err := tx.Where("sys_menu_id = ?", data.MenuId).Delete(&sysMenuApiData).Error

	if err != nil {
		core.Log.Error(" AddApis err " + err.Error())
		tx.Rollback()
		return err
	}

	var sysMenuApiAddData []models.SysMenuApi

	for _, apiId := range data.ApiIds {

		var tmp models.SysMenuApi
		tmp.SysMenuId = data.MenuId
		tmp.SysApiId = apiId

		sysMenuApiAddData = append(sysMenuApiAddData, tmp)
	}

	err = tx.Create(&sysMenuApiAddData).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
