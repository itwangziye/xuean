package apis

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type XaBill struct {
	api.Api
}

// GetPage 获取流水表列表
// @Summary 获取流水表列表
// @Description 获取流水表列表
// @Tags 流水表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.XaBill}} "{"code": 200, "data": [...]}"
// @Router /api/v1/xa-bill [get]
// @Security Bearer
func (e XaBill) GetPage(c *gin.Context) {
	req := dto.XaBillGetPageReq{}
	s := service.XaBill{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.XaBill, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取流水表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取流水表
// @Summary 获取流水表
// @Description 获取流水表
// @Tags 流水表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.XaBill} "{"code": 200, "data": [...]}"
// @Router /api/v1/xa-bill/{id} [get]
// @Security Bearer
func (e XaBill) Get(c *gin.Context) {
	req := dto.XaBillGetReq{}
	s := service.XaBill{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.XaBill

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取流水表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建流水表
// @Summary 创建流水表
// @Description 创建流水表
// @Tags 流水表
// @Accept application/json
// @Product application/json
// @Param data body dto.XaBillInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/xa-bill [post]
// @Security Bearer
func (e XaBill) Insert(c *gin.Context) {
	req := dto.XaBillInsertReq{}
	s := service.XaBill{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建流水表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改流水表
// @Summary 修改流水表
// @Description 修改流水表
// @Tags 流水表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.XaBillUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/xa-bill/{id} [put]
// @Security Bearer
func (e XaBill) Update(c *gin.Context) {
	req := dto.XaBillUpdateReq{}
	s := service.XaBill{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改流水表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除流水表
// @Summary 删除流水表
// @Description 删除流水表
// @Tags 流水表
// @Param data body dto.XaBillDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/xa-bill [delete]
// @Security Bearer
func (e XaBill) Delete(c *gin.Context) {
	s := service.XaBill{}
	req := dto.XaBillDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除流水表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
