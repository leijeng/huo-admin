package service

import (
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service/dto"
	"github.com/leijeng/huo-core/core"
	"github.com/leijeng/huo-core/core/base"
)

type SysRoleMenuService struct {
	*base.BaseService
}

var SerSysRoleMenu = SysRoleMenuService{
	base.NewService("sys"),
}

// AddMenus 角色添加菜单
func (s *SysRoleMenuService) AddMenus(data dto.AddRoleMenuDto) error {
	tx := s.DB().Model(&models.SysRoleMenu{}).Begin()

	var sysRoleMenuData models.SysRoleMenu

	err := tx.Where("sys_role_id = ?", data.RoleId).Delete(&sysRoleMenuData).Error

	if err != nil {
		core.Log.Error(" AddMenus err " + err.Error())
		tx.Rollback()
		return err
	}

	var sysRoleMenuAddData []models.SysRoleMenu

	for _, menuId := range data.MenuIds {

		var tmp models.SysRoleMenu
		tmp.SysRoleId = data.RoleId
		tmp.SysMenuId = menuId

		sysRoleMenuAddData = append(sysRoleMenuAddData, tmp)
	}

	err = tx.Create(&sysRoleMenuAddData).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
