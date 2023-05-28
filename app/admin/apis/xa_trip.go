package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type XaTrip struct {
	api.Api
}

// GetPage 获取行程表列表
// @Summary 获取行程表列表
// @Description 获取行程表列表
// @Tags 行程表
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.XaTrip}} "{"code": 200, "data": [...]}"
// @Router /api/v1/xa-trip [get]
// @Security Bearer
func (e XaTrip) GetPage(c *gin.Context) {
	req := dto.XaTripGetPageReq{}

	s := service.XaTrip{}
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
	list := make([]models.XaTrip, 0)
	var count int64

	var TotalMoney = dto.TotalMoney{}

	err = s.GetPage(&req, p, &list, &count, &TotalMoney)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取行程表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取行程表
// @Summary 获取行程表
// @Description 获取行程表
// @Tags 行程表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.XaTrip} "{"code": 200, "data": [...]}"
// @Router /api/v1/xa-trip/{id} [get]
// @Security Bearer
func (e XaTrip) Get(c *gin.Context) {
	req := dto.XaTripGetReq{}
	s := service.XaTrip{}
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
	var object models.XaTrip

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取行程表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建行程表
// @Summary 创建行程表
// @Description 创建行程表
// @Tags 行程表
// @Accept application/json
// @Product application/json
// @Param data body dto.XaTripInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/xa-trip [post]
// @Security Bearer
func (e XaTrip) Insert(c *gin.Context) {
	req := dto.XaTripInsertReq{}
	s := service.XaTrip{}
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
		e.Error(500, err, fmt.Sprintf("创建行程失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改行程表
// @Summary 修改行程表
// @Description 修改行程表
// @Tags 行程表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.XaTripUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/xa-trip/{id} [put]
// @Security Bearer
func (e XaTrip) Update(c *gin.Context) {
	req := dto.XaTripUpdateReq{}
	s := service.XaTrip{}
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
		e.Error(500, err, fmt.Sprintf("修改行程失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除行程表
// @Summary 删除行程表
// @Description 删除行程表
// @Tags 行程表
// @Param data body dto.XaTripDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/xa-trip [delete]
// @Security Bearer
func (e XaTrip) Delete(c *gin.Context) {
	s := service.XaTrip{}
	req := dto.XaTripDeleteReq{}
	fmt.Println(req)

	fmt.Println(req.GetId())
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
		e.Error(500, err, fmt.Sprintf("删除行程失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Review 审核行程表
// @Summary 审核行程表
// @Description 审核行程表
// @Tags 行程表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.XaTripUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/xa-trip/{id} [put]
// @Security Bearer
func (e XaTrip) Review(c *gin.Context) {
	req := dto.XaTipUpdateStatusReq{}
	s := service.XaTrip{}
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
		e.Error(500, err, fmt.Sprintf("审核行程失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetStatusId(), "审核成功")
}

func (e *XaTrip) GetTripList(c *gin.Context) {
	req := dto.XaTripListReq{}

	fmt.Println(req)
	s := service.XaTrip{}
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

	p := actions.GetPermissionFromContext(c)
	list := make([]models.XaTrip, 0)

	err = s.GetTripList(&req, p, &list)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取行程表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(list, "获取成功")
}
