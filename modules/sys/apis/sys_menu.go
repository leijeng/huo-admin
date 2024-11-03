package apis

import (
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service"
	"github.com/leijeng/huo-admin/modules/sys/service/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/leijeng/huo-core/core/base"
)

type SysMenuApi struct {
	base.BaseApi
}

var ApiSysMenu = SysMenuApi{}

// QueryPage 获取SysMenu列表
// @Summary 获取SysMenu列表
// @Tags sys-SysMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysMenuGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysMenu}} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu/page [post]
// @Security Bearer
func (e *SysMenuApi) QueryPage(c *gin.Context) {
	var req dto.SysMenuGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysMenu, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerSysMenu.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysMenu
// @Summary 获取SysMenu
// @Tags sys-SysMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu/get [post]
// @Security Bearer
func (e *SysMenuApi) Get(c *gin.Context) {
	var req dto.SysMenuGetReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMenu
	menuData, err := service.SerSysMenu.Get(&req, &data)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, menuData)
}

// Create 创建SysMenu
// @Summary 创建SysMenu
// @Tags sys-SysMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysMenuInsertReq true "body"
// @Success 200 {object} base.Resp{data=models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu/create [post]
// @Security Bearer
func (e *SysMenuApi) Create(c *gin.Context) {
	var req dto.SysMenuInsertReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMenu
	copier.Copy(&data, req)
	if err := service.SerSysMenu.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysMenu
// @Summary 更新SysMenu
// @Tags sys-SysMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysMenuUpdateReq true "body"
// @Success 200 {object} base.Resp{data=models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu/update [post]
// @Security Bearer
func (e *SysMenuApi) Update(c *gin.Context) {
	var req dto.SysMenuUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMenu
	copier.Copy(&data, req)
	if err := service.SerSysMenu.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysMenu
// @Summary 删除SysMenu
// @Tags sys-SysMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu/del [post]
// @Security Bearer
func (e *SysMenuApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysMenu.DelIds(&models.SysMenu{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// AddApis 添加关联路由
// @Summary 添加关联路由
// @Tags sys-SysMenuApi
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.AddMenuApiDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMenuApi} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu/addApis [post]
// @Security Bearer
func (e *SysMenuApi) AddApis(c *gin.Context) {
	var req dto.AddMenuApiDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysMenuApi.AddApis(req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
