package models

import (
	"go-admin/common/models"
)

type XaTrip struct {
	models.Model

	TripId       string `json:"tripId" gorm:"type:varchar(50);comment:行程编号"`
	TripName     string `json:"tripName" gorm:"type:varchar(20);comment:行程说明"`
	CarLink      string `json:"carLink" gorm:"type:varchar(20);comment:用车联系人"`
	CarId        string `json:"carId" gorm:"type:varchar(20);comment:车牌号"`
	DriverName   string `json:"driverName" gorm:"type:varchar(20);comment:司机姓名"`
	OperatorName string `json:"operatorName" gorm:"type:varchar(20);comment:经办人"`
	PreMoney     string `json:"preMoney" gorm:"type:decimal(19,2);comment:应付金额"`
	PayMoney     string `json:"payMoney" gorm:"type:decimal(19,2);comment:实付金额"`
	IsSettle     string `json:"isSettle" gorm:"type:tinyint(4);comment:是否结算，1是，2否"`
	IsInvoicing  string `json:"isInvoicing" gorm:"type:tinyint(4);comment:是否开票，1是，2否"`
	TripStatus   string `json:"tripStatus" gorm:"type:tinyint(4);comment:状态"`
	InvoiceId    string `json:"invoiceId" gorm:"type:varchar(50);comment:发票编号"`
	BillId       string `json:"billId" gorm:"type:varchar(50);comment:流水编号"`
	Counted      string `json:"counted" gorm:"type:varchar(50);comment:创建时间"`

	models.ModelTime
	models.ControlBy
}

func (XaTrip) TableName() string {
	return "xa_trip"
}

func (e *XaTrip) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *XaTrip) GetId() interface{} {
	return e.Id
}
