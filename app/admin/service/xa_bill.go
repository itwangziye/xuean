package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
	"strconv"

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
	c.CreatedAtOrder = "desc"

	if c.TripId != "" {
		var tripInfo models.XaTrip

		e.Orm.Table("xa_trip").Where("trip_id=?", c.TripId).Find(&tripInfo)
		c.BillId = tripInfo.BillId

		if tripInfo.Id == 0 {
			c.BillId = "0"
		}
	}

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

		if tripIds != "" {
			length := len(tripIds) - 1
			tripIds = tripIds[0:length]
		}
		value.TripId = tripIds

		*list = append(*list, value)
	}

	e.Orm.Table("xa_bill").Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		actions.Permission(data.TableName(), p),
	).Where("deleted_at is null").Pluck("sum(income) as income, sum(pay_out) as pay_out, sum(income-pay_out) as profit", &money)

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

	var totalMoney float64
	var tripList []models.XaTrip
	if c.BillType == "1" {
		// 修改行程信息
		if len(c.TripId) == 0 {
			return errors.New("请选择行程信息")
		}

		e.Orm.Table("xa_trip").Where("trip_id in ?", c.TripId).Where("is_settle != ?", "2").Find(&tripList)
		if len(tripList) == 0 {
			return errors.New("选择的行程信息不存在或已结算")
		}
		for _, value := range tripList {
			if value.TripStatus == "1" {
				return errors.New("选择的行程【" + value.TripId + "】尚未审核")
			}
			if value.TripStatus == "3" {
				return errors.New("选择的行程【" + value.TripId + "】未审核通过")
			}

			preMoney, _ := strconv.ParseFloat(value.PreMoney, 64)
			totalMoney = totalMoney + preMoney
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

		income, _ := strconv.ParseFloat(c.Income, 64)

		for _, val := range tripList {
			var payMoney float64
			preMoney, _ := strconv.ParseFloat(val.PreMoney, 64)

			payMoney = (preMoney / totalMoney) * income

			err = e.Orm.Table("xa_trip").Where("id = ?", val.Id).
				Update("bill_id", data.BillId).
				Update("is_settle", "2").
				Update("pay_money", payMoney).Error

			if err != nil {
				e.Log.Errorf("XaInvoiceService Insert error:%s \r\n", err)
				return err
			}
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

	// 同步更新行程实付金额
	var tripList []models.XaTrip

	e.Orm.Table("xa_trip").Where("bill_id = ?", data.BillId).Find(&tripList)

	if len(tripList) > 0 {
		var totalMoney float64
		for _, value := range tripList {
			var preMoney float64
			preMoney, _ = strconv.ParseFloat(value.PreMoney, 64)
			totalMoney = totalMoney + preMoney
		}

		for _, val := range tripList {
			var payMoney float64
			var preMoney float64
			preMoney, _ = strconv.ParseFloat(val.PreMoney, 64)
			income, _ := strconv.ParseFloat(data.Income, 64)
			payMoney = (preMoney / totalMoney) * income

			_ = e.Orm.Table("xa_trip").Where("id = ?", val.Id).
				Update("pay_money", payMoney).Error
		}
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

// UpdateStatus 修改XaBill对象
func (e *XaBill) UpdateStatus(c *dto.XaBillUpdateStatusReq, p *actions.DataPermission) error {
	if c.BillStatus != "2" && c.BillStatus != "3" {
		return errors.New("操作失败，参数错误")
	}
	var err error
	var data = models.XaBill{}
	var model models.XaBill

	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetStatusId())
	c.GenerateStatus(&model)

	if data.Id == 0 {
		return errors.New("操作失败，流水信息不存在或已被删除")
	}

	msg := "确认"
	if c.BillStatus == "3" {
		msg = "审核失败"
	}

	if data.BillStatus == c.BillStatus {
		return errors.New("操作失败，流水已" + msg)
	}

	update := e.Orm.Model(&model).Where("id = ?", &model.Id).Updates(&model)

	if update.Error != nil {
		e.Log.Errorf("XaBillService Save error:%s \r\n", err)
		return err

	}
	return nil
}
