package dto

import (
    "github.com/leijeng/huo-core/core/base"
    "github.com/leijeng/huo-admin/modules/sys/models"
)

type SysApiGetPageReq struct {
	base.ReqPage `query:"-"`
    SortOrder  string `json:"-" query:"type:order;column:id"`
    Status int `json:"status" form:"status" query:"column:status"` //
    
}

func (SysApiGetPageReq) TableName() string {
	return models.TBSysApi
}


//SysApi
type SysApiDto struct {
    
    Id int `json:"id"` //主键
    Title string `json:"title" form:"title"` // 
    Method string `json:"method" form:"method"` // 
    Path string `json:"path" form:"path"` // 
    PermType int `json:"permType" form:"permType"` // 
    Status int `json:"status" form:"status"` // 
}

