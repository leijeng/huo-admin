package apis

import (
	"errors"
	"github.com/leijeng/huo-admin/common/consts"
	"github.com/leijeng/huo-admin/common/utils"
	"github.com/leijeng/huo-admin/modules/sys/models"
	"github.com/leijeng/huo-admin/modules/sys/service"
	"github.com/leijeng/huo-admin/modules/sys/service/dto"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/leijeng/huo-core/core/base"
)

type SysUserApi struct {
	base.BaseApi
}

var ApiSysUser = SysUserApi{}

// QueryPage 获取SysUser列表
// @Summary 获取SysUser列表
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysUserGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysUser}} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user/page [post]
// @Security Bearer
func (e *SysUserApi) QueryPage(c *gin.Context) {
	var req dto.SysUserGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysUser, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerSysUser.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysUser
// @Summary 获取SysUser
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user/get [post]
// @Security Bearer
func (e *SysUserApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUser
	if err := service.SerSysUser.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysUser
// @Summary 创建SysUser
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysUserDto true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user/create [post]
// @Security Bearer
func (e *SysUserApi) Create(c *gin.Context) {
	var req dto.SysUserDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUser
	copier.Copy(&data, req)
	if err := service.SerSysUser.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysUser
// @Summary 更新SysUser
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysUserDto true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user/update [post]
// @Security Bearer
func (e *SysUserApi) Update(c *gin.Context) {
	var req dto.SysUserDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUser
	copier.Copy(&data, req)
	if err := service.SerSysUser.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysUser
// @Summary 删除SysUser
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user/del [post]
// @Security Bearer
func (e *SysUserApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysUser.DelIds(&models.SysUser{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// AddRoles 添加角色
// @Summary 添加角色
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Param data body dto.SysAddRoleDto true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user/addRoles [post]
// @Security Bearer
func (e *SysUserApi) AddRoles(c *gin.Context) {
	var req dto.SysAddRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	if len(req.RoleIds) < 0 {
		e.Error(c, errors.New("role empty"))
		return
	}

	if err := service.SerSysUser.AddRoles(req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// Login 用户登录
// @Summary 用户登录
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserLoginReq true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user/login [post]
// @Security Bearer
func (e *SysUserApi) Login(c *gin.Context) {
	var req dto.SysUserLoginReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	var userCondition = make(map[string]interface{})
	userCondition["username"] = req.Username

	var userData models.SysUser

	if err := service.SerSysUser.GetByMap(userCondition, &userData); err != nil {
		e.Error(c, err)
		return
	}

	data, err := service.SerSysUser.LoginPwd(req, userData)
	if err != nil {
		e.Error(c, err)
		return
	}

	e.Ok(c, data)
}

// Logout 用户登出
// @Summary 用户登出
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param authorization header string false "token信息"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /v2/admin/sys/sys-user/logout [post]
// @Security Bearer
func (e *SysUserApi) Logout(c *gin.Context) {
	client := utils.InitRedis()
	cacheKey := consts.LogoutJwtKey + utils.Md5(c.GetHeader("Authorization"))
	client.SetNX(c, cacheKey, 1, time.Second*86400*3)
	e.Ok(c)
}
