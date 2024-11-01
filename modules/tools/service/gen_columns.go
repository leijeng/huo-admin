package service

import (
	"github.com/leijeng/huo-admin/modules/tools/models"
	"github.com/leijeng/huo-core/core/base"
	"gorm.io/gorm"
)

type GenColumnsService struct {
	*base.BaseService
}

var SerGenColumns = GenColumnsService{
	base.NewService("sys"),
}

func (e *GenColumnsService) GetList(tx *gorm.DB, exclude bool, tableId int) ([]models.GenColumns, error) {
	var doc []models.GenColumns

	table := tx
	if table == nil {
		table = e.DB()
	}
	if err := table.Where("table_id = ?", tableId).Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}
