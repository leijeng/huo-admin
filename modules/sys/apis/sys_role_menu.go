package apis

import (
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service"
	"github.com/leijeng/huo-admin/modules/sys/service/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/leijeng/huo-core/core/base"
)

type SysRoleMenuApi struct {
	base.BaseApi
}

var ApiSysRoleMenu = SysRoleMenuApi{}

// QueryPage 获取SysRoleMenu列表
// @Summary 获取SysRoleMenu列表
// @Tags sys-SysRoleMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysRoleMenuGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysRoleMenu}} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role-menu/page [post]
// @Security Bearer
func (e *SysRoleMenuApi) QueryPage(c *gin.Context) {
	var req dto.SysRoleMenuGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysRoleMenu, 0, req.GetSize())
	var total int64

	if err := service.SerSysRoleMenu.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysRoleMenu
// @Summary 获取SysRoleMenu
// @Tags sys-SysRoleMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysRoleMenu} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role-menu/get [post]
// @Security Bearer
func (e *SysRoleMenuApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysRoleMenu
	if err := service.SerSysRoleMenu.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysRoleMenu
// @Summary 创建SysRoleMenu
// @Tags sys-SysRoleMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysRoleMenuDto true "body"
// @Success 200 {object} base.Resp{data=models.SysRoleMenu} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role-menu/create [post]
// @Security Bearer
func (e *SysRoleMenuApi) Create(c *gin.Context) {
	var req dto.SysRoleMenuDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysRoleMenu
	copier.Copy(&data, req)
	if err := service.SerSysRoleMenu.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysRoleMenu
// @Summary 更新SysRoleMenu
// @Tags sys-SysRoleMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysRoleMenuDto true "body"
// @Success 200 {object} base.Resp{data=models.SysRoleMenu} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role-menu/update [post]
// @Security Bearer
func (e *SysRoleMenuApi) Update(c *gin.Context) {
	var req dto.SysRoleMenuDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysRoleMenu
	copier.Copy(&data, req)
	if err := service.SerSysRoleMenu.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysRoleMenu
// @Summary 删除SysRoleMenu
// @Tags sys-SysRoleMenu
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysRoleMenu} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role-menu/del [post]
// @Security Bearer
func (e *SysRoleMenuApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysRoleMenu.DelIds(&models.SysRoleMenu{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
