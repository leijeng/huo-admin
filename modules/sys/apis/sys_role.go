package apis

import (
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service"
	"github.com/leijeng/huo-admin/modules/sys/service/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/leijeng/huo-core/core/base"
)

type SysRoleApi struct {
	base.BaseApi
}

var ApiSysRole = SysRoleApi{}

// QueryPage 获取SysRole列表
// @Summary 获取SysRole列表
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param Authorization header string false "token信息"
// @Param data body dto.SysRoleGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysRole}} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role/page [post]
// @Security Bearer
func (e *SysRoleApi) QueryPage(c *gin.Context) {
	var req dto.SysRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysRole, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerSysRole.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysRole
// @Summary 获取SysRole
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param Authorization header string false "token信息"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysRole} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role/get [post]
// @Security Bearer
func (e *SysRoleApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysRole
	if err := service.SerSysRole.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysRole
// @Summary 创建SysRole
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param Authorization header string false "token信息"
// @Param data body dto.SysRoleDto true "body"
// @Success 200 {object} base.Resp{data=models.SysRole} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role/create [post]
// @Security Bearer
func (e *SysRoleApi) Create(c *gin.Context) {
	var req dto.SysRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysRole
	copier.Copy(&data, req)
	if err := service.SerSysRole.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysRole
// @Summary 更新SysRole
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param Authorization header string false "token信息"
// @Param data body dto.SysRoleDto true "body"
// @Success 200 {object} base.Resp{data=models.SysRole} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role/update [post]
// @Security Bearer
func (e *SysRoleApi) Update(c *gin.Context) {
	var req dto.SysRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysRole
	copier.Copy(&data, req)
	if err := service.SerSysRole.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysRole
// @Summary 删除SysRole
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param Authorization header string false "token信息"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysRole} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-role/del [post]
// @Security Bearer
func (e *SysRoleApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysRole.DelIds(&models.SysRole{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
