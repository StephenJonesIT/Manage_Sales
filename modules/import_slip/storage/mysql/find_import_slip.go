package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/import_slip/model"

	"gorm.io/gorm"
)

func (sql *sqlStore) GetImportSlipById(ctx context.Context, cond map[string]interface{}) (*model.PhieuNhap, error){
	var data model.PhieuNhap
	if err := sql.db.Where(cond).Preload("ChiTiet").Find(&data).Error; err!=nil{
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
	}
	return &data, nil
}