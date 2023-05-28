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

type XaInvoice struct {
	api.Api
}

// GetPage 获取发票表列表
// @Summary 获取发票表列表
// @Description 获取发票表列表
// @Tags 发票表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.XaInvoice}} "{"code": 200, "data": [...]}"
// @Router /api/v1/xa-invoice [get]
// @Security Bearer
func (e XaInvoice) GetPage(c *gin.Context) {
	req := dto.XaInvoiceGetPageReq{}

	fmt.Println(req.PageSize)
	s := service.XaInvoice{}
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
	list := make([]models.XaInvoice, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取发票表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取发票表
// @Summary 获取发票表
// @Description 获取发票表
// @Tags 发票表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.XaInvoice} "{"code": 200, "data": [...]}"
// @Router /api/v1/xa-invoice/{id} [get]
// @Security Bearer
func (e XaInvoice) Get(c *gin.Context) {
	req := dto.XaInvoiceGetReq{}
	s := service.XaInvoice{}
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
	var object models.XaInvoice

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取发票表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建发票表
// @Summary 创建发票表
// @Description 创建发票表
// @Tags 发票表
// @Accept application/json
// @Product application/json
// @Param data body dto.XaInvoiceInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/xa-invoice [post]
// @Security Bearer
func (e XaInvoice) Insert(c *gin.Context) {
	req := dto.XaInvoiceInsertReq{}
	s := service.XaInvoice{}
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
		e.Error(500, err, fmt.Sprintf("创建发票表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改发票表
// @Summary 修改发票表
// @Description 修改发票表
// @Tags 发票表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.XaInvoiceUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/xa-invoice/{id} [put]
// @Security Bearer
func (e XaInvoice) Update(c *gin.Context) {
	req := dto.XaInvoiceUpdateReq{}
	s := service.XaInvoice{}
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
		e.Error(500, err, fmt.Sprintf("修改发票表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除发票表
// @Summary 删除发票表
// @Description 删除发票表
// @Tags 发票表
// @Param data body dto.XaInvoiceDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/xa-invoice [delete]
// @Security Bearer
func (e XaInvoice) Delete(c *gin.Context) {
	s := service.XaInvoice{}
	req := dto.XaInvoiceDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除发票表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Review 发票审核
// @Summary 发票审核
// @Description 发票审核
// @Tags 发票表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.XaInvoiceUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "确认成功"}"
// @Router /api/v1/xa-invoice/review/{id} [put]
// @Security Bearer
func (e XaInvoice) Review(c *gin.Context) {
	req := dto.XaInvoiceReviewReq{}
	s := service.XaInvoice{}
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

	err = s.UpdateStatus(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("确认发票表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "确认成功")
}
