package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"time"
)

type XaInvoiceGetPageReq struct {
	dto.Pagination `search:"-"`
	InvoiceId      string `form:"invoiceId"  search:"type:exact;column:invoice_id;table:xa_invoice"`
	TripId         string `form:"tripId"  search:"-"`
	InvoiceCompany string `form:"invoiceCompany"  search:"type:contains;column:invoice_company;table:xa_invoice"`
	Remark         string `form:"remark"  search:"type:contains;column:remark;table:xa_invoice"`
	BeginTime      string `form:"beginTime" search:"type:gte;column:counted;table:xa_invoice" comment:"创建时间"`
	EndTime        string `form:"endTime" search:"type:lte;column:counted;table:xa_invoice" comment:"创建时间"`
	XaInvoiceOrder
}

type XaInvoiceOrder struct {
	CreatedAtOrder string `search:"type:order;column:created_at;table:xa_invoice" form:"createdAtOrder"`
}

func (m *XaInvoiceGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type XaInvoiceInsertReq struct {
	Id             int      `json:"-" comment:""` //
	InvoiceId      string   `json:"invoiceId" comment:"发票号码" vd:"len($)>0"`
	InvoiceCompany string   `json:"invoiceCompany" comment:"发票单位" vd:"len($)>0"`
	Money          string   `json:"money" comment:"发票金额" vd:"len($)>0"`
	Remark         string   `json:"remark" comment:"发票备注"`
	InvoiceStatus  string   `json:"invoiceStatus" comment:"状态"`
	TripId         []string `json:"trip_id" comment:"行程编号"`
	InvoiceDate    string   `json:"invoiceDate" comment:"开票日期"`
	common.ControlBy
}

func (s *XaInvoiceInsertReq) Generate(model *models.XaInvoice) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}

	model.InvoiceId = s.InvoiceId
	model.InvoiceCompany = s.InvoiceCompany
	model.Money = s.Money
	model.Remark = s.Remark
	model.InvoiceStatus = s.InvoiceStatus
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Counted = time.Now().Format("2006-01-02")
	model.InvoiceDate = s.InvoiceDate

	if s.InvoiceDate == "" {
		model.InvoiceDate = time.Now().Format("2006-01-02")
	}
}

func (s *XaInvoiceInsertReq) GetId() interface{} {
	return s.Id
}

type XaInvoiceUpdateReq struct {
	Id             int    `uri:"id" comment:""` //
	InvoiceId      string `json:"invoiceId" comment:"发票号码"`
	InvoiceCompany string `json:"invoiceCompany" comment:"发票单位"`
	Money          string `json:"money" comment:"发票金额"`
	Remark         string `json:"remark" comment:"发票备注"`
	InvoiceStatus  string `json:"invoiceStatus" comment:"状态"`
	InvoiceDate    string `json:"invoiceDate" comment:"开票日期"`
	common.ControlBy
}

func (s *XaInvoiceUpdateReq) Generate(model *models.XaInvoice) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.InvoiceId = s.InvoiceId
	model.InvoiceCompany = s.InvoiceCompany
	model.Money = s.Money
	model.Remark = s.Remark
	model.InvoiceStatus = "1"
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Counted = time.Now().Format("2006-01-02")
	model.InvoiceDate = s.InvoiceDate

	if s.InvoiceDate == "" {
		model.InvoiceDate = time.Now().Format("2006-01-02")
	}
}

func (s *XaInvoiceUpdateReq) GetId() interface{} {
	return s.Id
}

// XaInvoiceGetReq 功能获取请求参数
type XaInvoiceGetReq struct {
	Id int `uri:"id"`
}

func (s *XaInvoiceGetReq) GetId() interface{} {
	return s.Id
}

// XaInvoiceReviewReq 功能获取请求参数
type XaInvoiceReviewReq struct {
	Id            int    `uri:"id"`
	InvoiceStatus string `json:"invoiceStatus" comment:"状态"`
	common.ControlBy
}

func (s *XaInvoiceReviewReq) GetId() interface{} {
	return s.Id
}

func (s *XaInvoiceReviewReq) Generate(model *models.XaInvoice) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.InvoiceStatus = s.InvoiceStatus
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

// XaInvoiceDeleteReq 功能删除请求参数
type XaInvoiceDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *XaInvoiceDeleteReq) GetId() interface{} {
	return s.Ids
}
