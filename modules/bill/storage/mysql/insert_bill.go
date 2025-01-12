package mysql

import (
	"context"
	"manage_sales/modules/bill/model"
	bonsai "manage_sales/modules/bonsai/model"
	report "manage_sales/modules/report/model"
	"time"

	"gorm.io/gorm"
)

func (sql *billRepository) CreateBill(ctx context.Context, newHoaDon *model.CreateHoaDon) error {
	tx := sql.db.Begin()

	var report report.Report
	if err := tx.Where("ma_bao_cao = ?", newHoaDon.MaBaoCao).Find(&report).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			now := time.Now()
			report.MaBaoCao = newHoaDon.MaBaoCao
			report.NgayTao = &now

			if err := tx.Create(&report).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Create(&newHoaDon).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, detail := range newHoaDon.ChiTiet {
		detail.MaHD = newHoaDon.MaHD
            
        query := `
		    INSERT INTO chi_tiet_hd (masp, mahd, so_luong) 
			VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE so_luong = so_luong + VALUES(so_luong);
		`
		if err := tx.Exec(query, detail.MaSP, detail.MaHD, detail.SoLuong).Error; err != nil {
			tx.Rollback()
			return err
		}

		var product bonsai.BonsaiItem
		if err := tx.Where("masp = ?", detail.MaSP).Find(&product).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				tx.Rollback()
				return err
			} else {
				tx.Rollback()
				return err
			}
		}

		product.SoLuong -= detail.SoLuong
		if err := tx.Where("masp = ?", product.MaSP).Updates(product).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
