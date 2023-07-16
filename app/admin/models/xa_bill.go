package models

import (
	"go-admin/common/models"
)

type XaBill struct {
	models.Model

	BillId       string `json:"billId" gorm:"type:varchar(50);comment:流水编号"`
	BillType     string `json:"billType" gorm:"type:tinyint(4);comment:流水类型，1车费结算，2返差"`
	BillObj      string `json:"billObj" gorm:"type:varchar(50);comment:交易对象"`
	PayType      string `json:"payType" gorm:"type:tinyint(4);comment:交易类型"`
	Income       string `json:"income" gorm:"type:decimal(19,2);comment:收入"`
	PayOut       string `json:"payOut" gorm:"type:decimal(19,2);comment:支出"`
	BillStatus   string `json:"billStatus" gorm:"type:tinyint(4);comment:状态"`
	Remark       string `json:"remark" gorm:"type:varchar(500);comment:备注"`
	OperatorName string `json:"operatorName" gorm:"type:varchar(50);comment:经办人"`
	Counted      string `json:"counted" gorm:"type:varchar(50);comment:创建日期"`
	TripId       string `json:"tripId" gorm:"-"`
	BillDate     string `json:"billDate" gorm:"type:varchar(50);comment:收款日期"`
	models.ModelTime
	models.ControlBy
}

func (XaBill) TableName() string {
	return "xa_bill"
}

func (e *XaBill) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *XaBill) GetId() interface{} {
	return e.Id
}
