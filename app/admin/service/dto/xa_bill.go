package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type XaBillGetPageReq struct {
	dto.Pagination `search:"-"`
	XaBillOrder
}

type XaBillOrder struct {
	BillId       string `form:"billId"  search:"type:exact;column:bill_id;table:xa_bill"`
	BillType     string `form:"billType"  search:"type:exact;column:bill_type;table:xa_bill"`
	BillObj      string `form:"billObj"  search:"type:contains;column:bill_obj;table:xa_bill"`
	PayType      string `form:"payType"  search:"type:exact;column:pay_type;table:xa_bill"`
	BillStatus   string `form:"billStatus"  search:"type:exact;column:bill_status;table:xa_bill"`
	OperatorName string `form:"operatorName"  search:"type:exact;column:operator_name;table:xa_bill"`
}

func (m *XaBillGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type XaBillInsertReq struct {
	Id           int    `json:"-" comment:""` //
	BillType     string `json:"billType" comment:"流水类型，1车费结算，2返差"`
	BillObj      string `json:"billObj" comment:"交易对象"`
	PayType      string `json:"payType" comment:"交易类型"`
	Income       string `json:"income" comment:"收入"`
	PayOut       string `json:"payOut" comment:"支出"`
	Remark       string `json:"remark" comment:"备注"`
	OperatorName string `json:"operatorName" comment:"经办人"`
	common.ControlBy
}

func (s *XaBillInsertReq) Generate(model *models.XaBill) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.BillId = "LSGL" + dto.RandStr(10)
	model.BillType = s.BillType
	model.BillObj = s.BillObj
	model.PayType = s.PayType
	model.Income = s.Income
	model.PayOut = s.PayOut

	if s.BillType == "2" {
		model.Income = "0"
	}

	if s.BillType == "1" {
		model.PayOut = "0"
	}
	model.BillStatus = "1"
	model.Remark = s.Remark
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.OperatorName = s.OperatorName
	model.Counted = time.Now().Format("2006-01-02")
}

func (s *XaBillInsertReq) GetId() interface{} {
	return s.Id
}

type XaBillUpdateReq struct {
	Id           int    `uri:"id" comment:""` //
	BillType     string `json:"billType" comment:"流水类型，1车费结算，2返差"`
	BillObj      string `json:"billObj" comment:"交易对象"`
	PayType      string `json:"payType" comment:"交易类型"`
	Income       string `json:"income" comment:"收入"`
	PayOut       string `json:"payOut" comment:"支出"`
	BillStatus   string `json:"billStatus" comment:"状态"`
	Remark       string `json:"remark" comment:"备注"`
	OperatorName string `json:"operatorName" comment:"经办人"`
	common.ControlBy
}

func (s *XaBillUpdateReq) Generate(model *models.XaBill) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	//model.BillType = s.BillType
	model.BillObj = s.BillObj
	model.PayType = s.PayType
	model.Income = s.Income
	model.PayOut = s.PayOut
	if s.BillType == "2" {
		model.Income = "0"
	}

	if s.BillType == "1" {
		model.PayOut = "0"
	}
	model.BillStatus = "1"
	model.Remark = s.Remark
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.OperatorName = s.OperatorName
	model.Counted = time.Now().Format("2006-01-02")
}

func (s *XaBillUpdateReq) GetId() interface{} {
	return s.Id
}

// XaBillGetReq 功能获取请求参数
type XaBillGetReq struct {
	Id int `uri:"id"`
}

func (s *XaBillGetReq) GetId() interface{} {
	return s.Id
}

// XaBillDeleteReq 功能删除请求参数
type XaBillDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *XaBillDeleteReq) GetId() interface{} {
	return s.Ids
}
