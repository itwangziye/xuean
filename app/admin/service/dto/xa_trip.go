package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type XaTripGetPageReq struct {
	dto.Pagination `search:"-"`
	XaTripOrder
}

type XaTripOrder struct {
	TripId       string `form:"tripId"  search:"type:exact;column:trip_id;table:xa_trip"`
	CarLink      string `form:"carLink"  search:"type:contains;column:car_link;table:xa_trip"`
	CarId        string `form:"carId"  search:"type:contains;column:car_id;table:xa_trip"`
	driverName   string `form:"driverName"  search:"type:exact;column:driver_name;table:xa_trip"`
	OperatorName string `form:"operatorName"  search:"type:contains;column:operator_name;table:xa_trip"`
	IsSettle     string `form:"isSettle"  search:"type:exact;column:is_settle;table:xa_trip"`
	IsInvoicing  string `form:"isInvoicing"  search:"type:exact;column:is_Invoicing;table:xa_trip"`
	Counted      string `form:"counted"  search:"type:exact;column:counted;table:xa_trip"`
}

type TotalMoney struct {
	Money1 string `json:"money1" comment:"总金额1"`
	Money2 string `json:"money2" comment:"总金额2"`
	Money3 string `json:"money3" comment:"总金额3"`
}

func (m *XaTripGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type XaTripInsertReq struct {
	Id int `json:"-" comment:""` //
	//TripId       string `json:"tripId" comment:"行程编号"`
	TripName     string `json:"tripName" comment:"行程说明" vd:"len($)>0"`
	CarLink      string `json:"carLink" comment:"用车联系人" vd:"len($)>0"`
	CarId        string `json:"carId" comment:"车牌号" vd:"len($)>0"`
	DriverName   string `json:"driverName" comment:"司机姓名" vd:"len($)>0"`
	OperatorName string `json:"operatorName" comment:"经办人" vd:"len($)>0"`
	PreMoney     string `json:"preMoney" comment:"应付金额" vd:"len($)>0"`
	PayMoney     string `json:"payMoney" comment:"实付金额"`
	IsSettle     string `json:"isSettle" comment:"是否结算，1是，2否" vd:"len($)>0"`
	IsInvoicing  string `json:"isInvoicing" comment:"是否开票，1是，2否" vd:"len($)>0"`
	//TripStatus   string `json:"tripStatus" comment:"状态"`
	InvoiceId string `json:"invoiceId" comment:"发票编号"`
	//BillId       string `json:"billId" comment:"流水编号"`
	InvoiceCompany string `json:"invoiceCompany" comment:"发票单位"`
	Money          string `json:"money" comment:"发票金额"`
	Remark         string `json:"remark" comment:"发票备注"`

	common.ControlBy
}

func (s *XaTripInsertReq) Generate(model *models.XaTrip) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	//model.TripId = s.TripId
	model.TripId = "XCBH" + dto.RandStr(10)
	model.TripName = s.TripName
	model.CarLink = s.CarLink
	model.CarId = s.CarId
	model.DriverName = s.DriverName
	model.OperatorName = s.OperatorName
	model.PreMoney = s.PreMoney
	model.IsSettle = s.IsSettle
	model.PayMoney = "0"

	if s.IsSettle == "1" {
		model.PayMoney = s.PayMoney
	}

	if s.IsInvoicing == "1" {
		model.InvoiceId = s.InvoiceId
	}

	model.IsInvoicing = s.IsInvoicing
	model.TripStatus = "1"
	//model.InvoiceId = s.InvoiceId
	//model.BillId = s.BillId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Counted = time.Now().Format("2006-01-02")
}

func (s *XaTripInsertReq) GetId() interface{} {
	return s.Id
}

type XaTripUpdateReq struct {
	Id           int    `uri:"id" comment:""` //
	TripName     string `json:"tripName" comment:"行程说明" vd:"len($)>0"`
	CarLink      string `json:"carLink" comment:"用车联系人" vd:"len($)>0"`
	CarId        string `json:"carId" comment:"车牌号" vd:"len($)>0"`
	DriverName   string `json:"driverName" comment:"司机姓名" vd:"len($)>0"`
	OperatorName string `json:"operatorName" comment:"经办人" vd:"len($)>0"`
	PreMoney     string `json:"preMoney" comment:"应付金额" vd:"len($)>0"`
	PayMoney     string `json:"payMoney" comment:"实付金额"`
	IsSettle     string `json:"isSettle" comment:"是否结算，1是，2否" vd:"len($)>0"`
	IsInvoicing  string `json:"isInvoicing" comment:"是否开票，1是，2否" vd:"len($)>0"`
	//TripStatus   string `json:"tripStatus" comment:"状态"`
	InvoiceId string `json:"invoiceId" comment:"发票编号"`
	//BillId       string `json:"billId" comment:"流水编号"`
	common.ControlBy
}

func (s *XaTripUpdateReq) Generate(model *models.XaTrip) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.TripName = s.TripName
	model.CarLink = s.CarLink
	model.CarId = s.CarId
	model.DriverName = s.DriverName
	model.OperatorName = s.OperatorName
	model.PreMoney = s.PreMoney
	model.PayMoney = "0"
	model.IsSettle = s.IsSettle
	model.IsInvoicing = s.IsInvoicing

	if s.IsSettle == "1" {
		model.PayMoney = s.PayMoney
	}

	if s.IsInvoicing == "1" {
		model.InvoiceId = s.InvoiceId
	}
	model.TripStatus = "1"
	//model.BillId = s.BillId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Counted = time.Now().Format("2006-01-02")
}

type XaTipUpdateStatusReq struct {
	Id         int    `uri:"id" comment:""` //
	TripStatus string `json:"tripStatus" comment:"状态"`
	common.ControlBy
}

func (s *XaTipUpdateStatusReq) GenerateStatus(model *models.XaTrip) {
	model.Id = s.Id
	model.TripStatus = s.TripStatus
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *XaTripUpdateReq) GetId() interface{} {
	return s.Id
}

func (s *XaTipUpdateStatusReq) GetStatusId() interface{} {
	return s.Id
}

// XaTripGetReq 功能获取请求参数
type XaTripGetReq struct {
	Id int `uri:"id"`
}

func (s *XaTripGetReq) GetId() interface{} {
	return s.Id
}

// XaTripDeleteReq 功能删除请求参数
type XaTripDeleteReq struct {
	Ids []int `json:"ids"`
}

// XaTripListReq 获取指定行程数据参数
type XaTripListReq struct {
	Ids []int `json:"ids"`
}

func (s *XaTripDeleteReq) GetId() interface{} {
	return s.Ids
}

func (s *XaTripListReq) GetId() interface{} {
	return s.Ids
}
