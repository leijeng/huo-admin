package apis

import (
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service"
	"github.com/leijeng/huo-admin/modules/sys/service/dto"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/leijeng/huo-core/core/base"
)

type SysMenuApiApi struct {
	base.BaseApi
}

var ApiSysMenuApi = SysMenuApiApi{}

// QueryPage 获取SysMenuApi列表
// @Summary 获取SysMenuApi列表
// @Tags sys-SysMenuApi
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysMenuApiGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysMenuApi}} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu-api/page [post]
// @Security Bearer
func (e *SysMenuApiApi) QueryPage(c *gin.Context) {
	var req dto.SysMenuApiGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysMenuApi, 0, req.GetSize())
	var total int64

	if err := service.SerSysMenuApi.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysMenuApi
// @Summary 获取SysMenuApi
// @Tags sys-SysMenuApi
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysMenuApi} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu-api/get [post]
// @Security Bearer
func (e *SysMenuApiApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMenuApi
	if err := service.SerSysMenuApi.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysMenuApi
// @Summary 创建SysMenuApi
// @Tags sys-SysMenuApi
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysMenuApiDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMenuApi} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu-api/create [post]
// @Security Bearer
func (e *SysMenuApiApi) Create(c *gin.Context) {
	var req dto.SysMenuApiDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMenuApi
	copier.Copy(&data, req)
	if err := service.SerSysMenuApi.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysMenuApi
// @Summary 更新SysMenuApi
// @Tags sys-SysMenuApi
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysMenuApiDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMenuApi} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu-api/update [post]
// @Security Bearer
func (e *SysMenuApiApi) Update(c *gin.Context) {
	var req dto.SysMenuApiDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMenuApi
	copier.Copy(&data, req)
	if err := service.SerSysMenuApi.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysMenuApi
// @Summary 删除SysMenuApi
// @Tags sys-SysMenuApi
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysMenuApi} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-menu-api/del [post]
// @Security Bearer
func (e *SysMenuApiApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysMenuApi.DelIds(&models.SysMenuApi{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
