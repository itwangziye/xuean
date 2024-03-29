package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type XaBillGetPageReq struct {
	dto.Pagination `search:"-"`
	BillId         string `form:"billId"  search:"type:exact;column:bill_id;table:xa_bill"`
	TripId         string `form:"tripId"  search:"-"`
	BillType       string `form:"billType"  search:"type:exact;column:bill_type;table:xa_bill"`
	BillObj        string `form:"billObj"  search:"type:contains;column:bill_obj;table:xa_bill"`
	PayType        string `form:"payType"  search:"type:exact;column:pay_type;table:xa_bill"`
	BillStatus     string `form:"billStatus"  search:"type:exact;column:bill_status;table:xa_bill"`
	OperatorName   string `form:"operatorName"  search:"type:contains;column:operator_name;table:xa_bill"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:xa_bill"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:counted;table:xa_bill" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:counted;table:xa_bill" comment:"创建时间"`
	XaBillOrder
}

type XaBillOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:xa_bill" form:"createdAtOrder"`
}

func (m *XaBillGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type TotalMoneyBill struct {
	Income float64 `json:"Income" comment:"总收入"`
	PayOut float64 `json:"PayOut" comment:"总支出"`
	Profit float64 `json:"Profit" comment:"净利润"`
}

type XaBillInsertReq struct {
	Id           int      `json:"-" comment:""` //
	BillType     string   `json:"billType" comment:"流水类型，1车费结算，2返差"`
	BillObj      string   `json:"billObj" comment:"交易对象"`
	PayType      string   `json:"payType" comment:"交易类型"`
	Income       string   `json:"income" comment:"收入"`
	PayOut       string   `json:"payOut" comment:"支出"`
	Remark       string   `json:"remark" comment:"备注"`
	OperatorName string   `json:"operatorName" comment:"经办人"`
	TripId       []string `json:"trip_id" comment:"行程编号"`
	BillDate     string   `json:"billDate" comment:"收款日期"`
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
	model.BillDate = s.BillDate

	if s.BillDate == "" {
		model.BillDate = time.Now().Format("2006-01-02")
	}
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
	BillDate     string `json:"billDate" comment:"收款日期"`
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
	model.BillDate = s.BillDate

	if s.BillDate == "" {
		model.BillDate = time.Now().Format("2006-01-02")
	}
}

type XaBillUpdateStatusReq struct {
	Id         int    `uri:"id" comment:""` //
	BillStatus string `json:"BillStatus" comment:"状态"`
	common.ControlBy
}

func (s *XaBillUpdateStatusReq) GetStatusId() interface{} {
	return s.Id
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

func (s *XaBillUpdateStatusReq) GenerateStatus(model *models.XaBill) {
	model.Id = s.Id
	model.BillStatus = s.BillStatus
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}
