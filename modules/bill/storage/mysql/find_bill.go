package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bill/model"

	"gorm.io/gorm"
)

func (sql *billRepository) GetBillById(ctx context.Context, cond map[string]interface{}) (*model.HoaDon, error){
	var data model.HoaDon
	if err := sql.db.Where(cond).Preload("ChiTiet").Find(&data).Error; err!=nil{
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
	}
	return &data, nil
}