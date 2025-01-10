package mysql

import (
	"context"
	"manage_sales/modules/import_slip/model"

	"gorm.io/gorm"
)

func (sql *sqlStore) UpdateImportSlipById(ctx context.Context, cond map[string]interface{}, updateData *model.PhieuNhapUpdate) error{
	tx := sql.db.Begin()

	if err := tx.Where(cond).Updates(&updateData).Error; err != nil{
		tx.Rollback()
		return err
	}

	for _, data := range updateData.ChiTiet {
		var existingDetail model.ChiTietPhieuNhap
		if  err:= tx.Where("masp = ? AND mapn = ?", data.MaSP, updateData.MaPN).First(&existingDetail).Error; err != nil{
			if err == gorm.ErrRecordNotFound {
				data.MaPN = updateData.MaPN
				if err:= tx.Create(&data).Error; err!=nil{
					tx.Rollback()
					return err
				}
			}else{
				tx.Rollback()
				return err
			}
		}else{
			existingDetail.SoLuong = data.SoLuong
			if err := tx.Where("masp = ? AND mapn = ?", data.MaSP, updateData.MaPN).Updates(&existingDetail).Error; err!=nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}