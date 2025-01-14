package mysql

import (
    "context"
    "manage_sales/modules/import_slip/model"
    bonsai "manage_sales/modules/bonsai/model"
    report "manage_sales/modules/report/model"
    "gorm.io/gorm"
)

func (sql *sqlStore) UpdateImportSlipById(ctx context.Context, cond map[string]interface{}, updateData *model.PhieuNhapUpdate) error {
    tx := sql.db.Begin()

    if len(cond) == 0 {
        tx.Rollback()
        return gorm.ErrMissingWhereClause
    }

    // Cập nhật phiếu nhập
    if err := tx.Model(&model.PhieuNhap{}).Where(cond).Updates(updateData).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Nếu TrangThai = "Done", cập nhật chi phí trong bảng báo cáo
    if updateData.TrangThai != nil && *updateData.TrangThai == model.Done {
        var report report.Report
        if err := tx.Where("ma_bao_cao = ?", updateData.MaBaoCao).First(&report).Error; err != nil {
            tx.Rollback()
            return err
        }
        report.ChiPhi += updateData.TongTien
        if err := tx.Where("ma_bao_cao = ?", updateData.MaBaoCao).Save(&report).Error; err != nil {
            tx.Rollback()
            return err
        }
    }

    updateMaSPs := make(map[string]bool)
    for _, data := range updateData.ChiTiet {
        updateMaSPs[data.MaSP] = true
    }

    var currentDetails []model.ChiTietPhieuNhap
    if err := tx.Where("mapn = ?", updateData.MaPN).Find(&currentDetails).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Xóa các chi tiết không có trong danh sách cập nhật
    for _, currentDetail := range currentDetails {
        if _, exists := updateMaSPs[currentDetail.MaSP]; !exists {
            if err := tx.Where("masp = ? AND mapn = ?", currentDetail.MaSP, updateData.MaPN).Delete(&model.ChiTietPhieuNhap{}).Error; err != nil {
                tx.Rollback()
                return err
            }
        }
    }

    // Cập nhật hoặc thêm mới các chi tiết phiếu nhập
    for _, data := range updateData.ChiTiet {
        var existingDetail model.ChiTietPhieuNhap
        if err := tx.Where("masp = ? AND mapn = ?", data.MaSP, updateData.MaPN).First(&existingDetail).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                data.MaPN = updateData.MaPN
                if err := tx.Create(&data).Error; err != nil {
                    tx.Rollback()
                    return err
                }
            } else {
                tx.Rollback()
                return err
            }
        } else {
            existingDetail.SoLuong = data.SoLuong
            if err := tx.Where("masp = ? AND mapn = ?", data.MaSP, updateData.MaPN).Updates(&existingDetail).Error; err != nil {
                tx.Rollback()
                return err
            }
        }

        // Nếu TrangThai = "Done", cập nhật bảng sản phẩm
        if updateData.TrangThai != nil && *updateData.TrangThai == model.Done {
            var sanPham bonsai.BonsaiItem
            if err := tx.Where("masp = ?", data.MaSP).First(&sanPham).Error; err != nil {
                tx.Rollback()
                return err
            }
            sanPham.SoLuong += data.SoLuong
            if err := tx.Save(&sanPham).Error; err != nil {
                tx.Rollback()
                return err
            }
        }
    }

    // Commit giao dịch nếu không có lỗi
    if err := tx.Commit().Error; err != nil {
        return err
    }

    return nil
}
