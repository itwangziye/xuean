package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type XaInvoice struct {
	service.Service
}

// GetPage 获取XaInvoice列表
func (e *XaInvoice) GetPage(c *dto.XaInvoiceGetPageReq, p *actions.DataPermission, list *[]models.XaInvoice, count *int64) error {
	var err error
	var data models.XaInvoice

	if c.TripId != "" {
		var tripInfo models.XaTrip

		e.Orm.Table("xa_trip").Where("trip_id=?", c.TripId).Find(&tripInfo)
		c.InvoiceId = tripInfo.InvoiceId
	}
	var listTemp = make([]models.XaInvoice, 0)

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(&listTemp).Count(count).Error
	for _, value := range listTemp {
		invoiceId := value.InvoiceId

		var tripList []models.XaTrip

		e.Orm.Table("xa_trip").Where("invoice_id = ?", invoiceId).Find(&tripList)

		tripIds := ""
		if len(tripList) > 0 {
			for _, v := range tripList {
				tripIds = v.TripId + ","
			}
		}
		value.TripId = tripIds

		*list = append(*list, value)
	}
	if err != nil {
		e.Log.Errorf("XaInvoiceService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取XaInvoice对象
func (e *XaInvoice) Get(d *dto.XaInvoiceGetReq, p *actions.DataPermission, model *models.XaInvoice) error {
	var data models.XaInvoice

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetXaInvoice error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建XaInvoice对象
func (e *XaInvoice) Insert(c *dto.XaInvoiceInsertReq) error {
	// 修改行程信息
	if len(c.TripId) == 0 {
		return errors.New("请选择行程信息")
	}

	var tripList []models.XaTrip
	e.Orm.Table("xa_trip").Where("trip_id in ?", c.TripId).Where("is_invoicing != ?", "2").Find(&tripList)
	if len(tripList) == 0 {
		return errors.New("选择的行程信息不存在或已出票")
	}

	var checkXaInvoice models.XaInvoice

	e.Orm.Table("xa_invoice").Where("invoice_id=?", c.InvoiceId).Find(&checkXaInvoice)

	if checkXaInvoice.InvoiceId != "" {
		return errors.New("该发票号码已添加，请勿重复添加")
	}

	var err error
	var data models.XaInvoice
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("XaInvoiceService Insert error:%s \r\n", err)
		return err
	}

	err = e.Orm.Table("xa_trip").Where("trip_id in ?", c.TripId).
		Update("invoice_id", c.InvoiceId).
		Update("is_invoicing", "2").Error

	if err != nil {
		e.Log.Errorf("XaInvoiceService Insert error:%s \r\n", err)
		return err
	}

	return nil
}

// Update 修改XaInvoice对象
func (e *XaInvoice) Update(c *dto.XaInvoiceUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.XaInvoice{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("XaInvoiceService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除XaInvoice
func (e *XaInvoice) Remove(d *dto.XaInvoiceDeleteReq, p *actions.DataPermission) error {
	var data models.XaInvoice

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveXaInvoice error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *XaInvoice) UpdateStatus(d *dto.XaInvoiceReviewReq, p *actions.DataPermission) error {
	if d.InvoiceStatus != "2" && d.InvoiceStatus != "3" {
		return errors.New("操作失败，参数错误")
	}

	var err error
	var data = models.XaInvoice{}
	var model models.XaInvoice

	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, d.GetId())

	d.Generate(&model)

	if data.Id == 0 {
		return errors.New("操作失败，发票信息不存在")
	}

	if data.InvoiceStatus == d.InvoiceStatus {
		return errors.New("操作失败，发票已审核")
	}

	update := e.Orm.Model(&model).Where("id=?", d.Id).Updates(&model)

	if update.Error != nil {
		e.Log.Errorf("XaInvoiceService Review error:%s \r\n", err)
		return err

	}
	return nil
}
