package mysql

import (
	"context"
	"fmt"
	"manage_sales/modules/import_slip/model"
)

func (sql *sqlStore) CreateImportSlip(ctx context.Context, newPhieuNhap *model.PhieuNhapCreate) error {
	
	tx := sql.db.Begin()
	if err := tx.Create(&newPhieuNhap).Error; err!=nil {
		tx.Rollback()
		return err
	}

	if newPhieuNhap.TongTien < 0 {
		tx.Rollback()
		return fmt.Errorf("invalid TongTien")
	}

	for _, detail := range newPhieuNhap.ChiTiet {
		detail.MaPN = newPhieuNhap.MaPN
		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}