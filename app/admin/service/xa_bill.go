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

type XaBill struct {
	service.Service
}

// GetPage 获取XaBill列表
func (e *XaBill) GetPage(c *dto.XaBillGetPageReq, p *actions.DataPermission, list *[]models.XaBill, count *int64, money *dto.TotalMoneyBill) error {
	var err error
	var data models.XaBill

	var listTemp = make([]models.XaBill, 0)

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(&listTemp).Limit(-1).Offset(-1).
		Count(count).Error

	for _, value := range listTemp {
		billId := value.BillId

		var tripList []models.XaTrip

		e.Orm.Table("xa_trip").Where("bill_id = ?", billId).Find(&tripList)

		tripIds := ""
		if len(tripList) > 0 {
			for _, v := range tripList {
				tripIds = v.TripId + ","
			}
		}
		value.TripId = tripIds

		*list = append(*list, value)
	}

	e.Orm.Table("xa_bill").Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		actions.Permission(data.TableName(), p),
	).Pluck("sum(income) as income, sum(pay_out) as pay_out, sum(income-pay_out) as profit", &money)

	if err != nil {
		e.Log.Errorf("XaBillService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取XaBill对象
func (e *XaBill) Get(d *dto.XaBillGetReq, p *actions.DataPermission, model *models.XaBill) error {
	var data models.XaBill

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetXaBill error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建XaBill对象
func (e *XaBill) Insert(c *dto.XaBillInsertReq) error {

	if c.BillType == "1" {
		// 修改行程信息
		if len(c.TripId) == 0 {
			return errors.New("请选择行程信息")
		}

		var tripList []models.XaTrip
		e.Orm.Table("xa_trip").Where("trip_id in ?", c.TripId).Where("is_settle != ?", "2").Find(&tripList)
		if len(tripList) == 0 {
			return errors.New("选择的行程信息不存在或已结算")
		}

	}

	var err error
	var data models.XaBill
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("XaBillService Insert error:%s \r\n", err)
		return err
	}

	if c.BillType == "1" {
		err = e.Orm.Table("xa_trip").Where("trip_id in ?", c.TripId).
			Update("bill_id", data.BillId).
			Update("is_settle", "2").Error

		if err != nil {
			e.Log.Errorf("XaInvoiceService Insert error:%s \r\n", err)
			return err
		}
	}
	return nil
}

// Update 修改XaBill对象
func (e *XaBill) Update(c *dto.XaBillUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.XaBill{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	if data.BillType != c.BillType {
		return errors.New("流水类型不能修改")
	}

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("XaBillService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除XaBill
func (e *XaBill) Remove(d *dto.XaBillDeleteReq, p *actions.DataPermission) error {
	var data models.XaBill

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveXaBill error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
