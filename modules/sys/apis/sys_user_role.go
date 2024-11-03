package apis

import (
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service"
	"github.com/leijeng/huo-admin/modules/sys/service/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/leijeng/huo-core/core/base"
)

type SysUserRoleApi struct {
	base.BaseApi
}

var ApiSysUserRole = SysUserRoleApi{}

// QueryPage 获取SysUserRole列表
// @Summary 获取SysUserRole列表
// @Tags sys-SysUserRole
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysUserRoleGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysUserRole}} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user-role/page [post]
// @Security Bearer
func (e *SysUserRoleApi) QueryPage(c *gin.Context) {
	var req dto.SysUserRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysUserRole, 0, req.GetSize())
	var total int64

	if err := service.SerSysUserRole.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysUserRole
// @Summary 获取SysUserRole
// @Tags sys-SysUserRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysUserRole} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user-role/get [post]
// @Security Bearer
func (e *SysUserRoleApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUserRole
	if err := service.SerSysUserRole.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysUserRole
// @Summary 创建SysUserRole
// @Tags sys-SysUserRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysUserRoleDto true "body"
// @Success 200 {object} base.Resp{data=models.SysUserRole} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user-role/create [post]
// @Security Bearer
func (e *SysUserRoleApi) Create(c *gin.Context) {
	var req dto.SysUserRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUserRole
	copier.Copy(&data, req)
	if err := service.SerSysUserRole.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysUserRole
// @Summary 更新SysUserRole
// @Tags sys-SysUserRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysUserRoleDto true "body"
// @Success 200 {object} base.Resp{data=models.SysUserRole} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user-role/update [post]
// @Security Bearer
func (e *SysUserRoleApi) Update(c *gin.Context) {
	var req dto.SysUserRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUserRole
	copier.Copy(&data, req)
	if err := service.SerSysUserRole.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysUserRole
// @Summary 删除SysUserRole
// @Tags sys-SysUserRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysUserRole} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user-role/del [post]
// @Security Bearer
func (e *SysUserRoleApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysUserRole.DelIds(&models.SysUserRole{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
