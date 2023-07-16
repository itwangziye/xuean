package models

import (
	"go-admin/common/models"
)

type XaInvoice struct {
	models.Model

	InvoiceId      string `json:"invoiceId" gorm:"type:varchar(50);comment:发票号码"`
	InvoiceCompany string `json:"invoiceCompany" gorm:"type:varchar(50);comment:发票单位"`
	Money          string `json:"money" gorm:"type:decimal(19,2);comment:发票金额"`
	Remark         string `json:"remark" gorm:"type:varchar(500);comment:发票备注"`
	InvoiceStatus  string `json:"invoiceStatus" gorm:"type:tinyint(4);comment:状态"`
	Counted        string `json:"counted" gorm:"type:varchar(50);comment:创建日期"`
	TripId         string `json:"tripId" gorm:"-"`
	InvoiceDate    string `json:"invoiceDate" gorm:"type:varchar(50);comment:开票日期"`
	models.ModelTime
	models.ControlBy
}

func (XaInvoice) TableName() string {
	return "xa_invoice"
}

func (e *XaInvoice) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *XaInvoice) GetId() interface{} {
	return e.Id
}
