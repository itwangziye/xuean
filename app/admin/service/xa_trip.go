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

type XaTrip struct {
	service.Service
}

// GetPage 获取XaTrip列表
func (e *XaTrip) GetPage(c *dto.XaTripGetPageReq, p *actions.DataPermission, list *[]models.XaTrip, count *int64, money *dto.TotalMoney) error {
	var err error
	var data models.XaTrip
	c.CreatedAtOrder = "desc"
	var listTemp = make([]models.XaTrip, 0)

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(&listTemp).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("XaTripService GetPage error:%s \r\n", err)
		return err
	}

	for _, value := range listTemp {
		tripDate := value.TripDate

		if len(tripDate) >= 10 {
			value.TripDate = tripDate[0:10]
		}
		*list = append(*list, value)
	}

	e.Orm.Table("xa_trip").Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		actions.Permission(data.TableName(), p),
	).Where("deleted_at is null").Pluck("sum(pre_money) as money1, sum(pay_money) as money2, sum(if (is_settle=1, pre_money, 0)) as money3", &money)

	return nil
}

// Get 获取XaTrip对象
func (e *XaTrip) Get(d *dto.XaTripGetReq, p *actions.DataPermission, model *models.XaTrip) error {
	var data models.XaTrip

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetXaTrip error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建XaTrip对象
func (e *XaTrip) Insert(c *dto.XaTripInsertReq) error {
	var err error
	var data models.XaTrip
	c.Generate(&data)

	if len(c.TripMark) > 500 {
		return errors.New("行程备注不能超过500字")
	}

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("XaTripService Insert error:%s \r\n", err)
		return err
	}

	// 流水增加
	if c.IsSettle == "2" {
		var xaBillData models.XaBill
		xaBillData.BillId = "LSGL" + cDto.RandStr(10)
		xaBillData.BillType = "1"
		xaBillData.BillObj = c.InvoiceCompany
		xaBillData.Income = c.PayMoney
		xaBillData.PayOut = "0"
		xaBillData.BillStatus = "1"
		xaBillData.PayType = "1"
		xaBillData.OperatorName = c.OperatorName
		xaBillData.Counted = data.Counted

		err3 := e.Orm.Table("xa_bill").Create(&xaBillData).Error
		if err3 != nil {
			e.Log.Errorf("XaTripService Insert error:%s \r\n", err3)
			return err3
		}
	}

	//发票增加
	if c.IsInvoicing == "2" {
		var checkXaInvoice models.XaInvoice
		e.Orm.Table("xa_invoice").Where("invoice_id=?", c.InvoiceId).Find(&checkXaInvoice)

		if checkXaInvoice.InvoiceId != "" {
			return errors.New("该发票号码已添加，请勿重复操作")
		}
		var xaInvoiceData models.XaInvoice
		xaInvoiceData.InvoiceId = c.InvoiceId
		xaInvoiceData.InvoiceCompany = c.InvoiceCompany
		xaInvoiceData.Remark = c.Remark
		xaInvoiceData.Money = c.Money
		xaInvoiceData.CreateBy = c.CreateBy
		xaInvoiceData.InvoiceStatus = "1"
		xaInvoiceData.Counted = data.Counted

		err2 := e.Orm.Table("xa_invoice").Create(&xaInvoiceData).Error
		if err2 != nil {
			e.Log.Errorf("XaTripService Insert error:%s \r\n", err2)
			return err2
		}
	}
	return nil
}

// Update 修改XaTrip对象
func (e *XaTrip) Update(c *dto.XaTripUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.XaTrip{}

	if len(c.TripMark) > 500 {
		return errors.New("行程备注不能超过500字")
	}

	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("XaTripService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除XaTrip
func (e *XaTrip) Remove(d *dto.XaTripDeleteReq, p *actions.DataPermission) error {
	var data models.XaTrip

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveXaTrip error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// UpdateStatus 修改XaTrip对象
func (e *XaTrip) UpdateStatus(c *dto.XaTipUpdateStatusReq, p *actions.DataPermission) error {
	if c.TripStatus != "2" && c.TripStatus != "3" {
		return errors.New("操作失败，参数错误")
	}
	var err error
	var data = models.XaTrip{}
	var model models.XaTrip

	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetStatusId())
	c.GenerateStatus(&model)

	if data.Id == 0 {
		return errors.New("操作失败，行程信息不存在或已被删除")
	}

	msg := "确认"
	if c.TripStatus == "3" {
		msg = "审核失败"
	}

	if data.TripStatus == c.TripStatus {
		return errors.New("操作失败，行程已" + msg)
	}

	update := e.Orm.Model(&model).Where("id = ?", &model.Id).Updates(&model)

	if update.Error != nil {
		e.Log.Errorf("XaTripService Save error:%s \r\n", err)
		return err

	}
	return nil
}

func (e *XaTrip) GetTripList(c *dto.XaTripListReq, p *actions.DataPermission, list *[]models.XaTrip) error {

	var ids = c.GetId()

	var data models.XaTrip
	var err error

	err = e.Orm.Model(&data).Where("id in ?", ids).Find(&list).Error

	if err != nil {
		e.Log.Errorf("获取行程信息失败 %s \r\n", err)
		return err
	}
	return nil
}
