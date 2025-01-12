package mysql

import (
	"context"
	"fmt"
	"manage_sales/modules/import_slip/model"
	report "manage_sales/modules/report/model"
	"time"

	"gorm.io/gorm"
)

func (sql *sqlStore) CreateImportSlip(ctx context.Context, newPhieuNhap *model.PhieuNhapCreate) error {
	tx := sql.db.Begin()

	// Validate TongTien before any database operations
	if newPhieuNhap.TongTien < 0 {
		tx.Rollback()
		return fmt.Errorf("invalid TongTien")
	}

	var report report.Report
	// Check if the report exists or create it if not
	if err := tx.Where("ma_bao_cao = ?", newPhieuNhap.MaBaoCao).First(&report).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			now := time.Now()
			report.MaBaoCao = newPhieuNhap.MaBaoCao
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

	// Insert the import slip
	if err := tx.Create(&newPhieuNhap).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert or update details for the import slip
	for _, detail := range newPhieuNhap.ChiTiet {
		detail.MaPN = newPhieuNhap.MaPN
		// Use ON DUPLICATE KEY UPDATE to update the quantity if the record exists
		query := `
			INSERT INTO chi_tiet_pn (masp, mapn, so_luong) 
			VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE so_luong = so_luong + VALUES(so_luong);
		`
		if err := tx.Exec(query, detail.MaSP, detail.MaPN, detail.SoLuong).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
