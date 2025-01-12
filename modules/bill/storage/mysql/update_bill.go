package mysql

import (
	"context"
	"manage_sales/modules/bill/model"
    report "manage_sales/modules/report/model"
    bonsai "manage_sales/modules/bonsai/model"
	"gorm.io/gorm"
)

func (sql *billRepository) UpdateBillById(ctx context.Context, cond map[string]interface{}, updateData *model.UpdateHoaDon) error{
	tx := sql.db.Begin()

    if len(cond) == 0 {
        tx.Rollback()
        return gorm.ErrMissingWhereClause
    }

    // Cập nhật phiếu nhập
    if err := tx.Model(&model.HoaDon{}).Where(cond).Updates(updateData).Error; err != nil {
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
        report.DoanhThu += updateData.TongTien
        if err := tx.Where("ma_bao_cao = ?", updateData.MaBaoCao).Save(&report).Error; err != nil {
            tx.Rollback()
            return err
        }
    }

    updateMaSPs := make(map[string]bool)
    for _, data := range updateData.ChiTiet {
        updateMaSPs[data.MaSP] = true
    }

    var currentDetails []model.ChiTietHoaDon
    if err := tx.Where("mahd = ?", updateData.MaHD).Find(&currentDetails).Error; err != nil {
        tx.Rollback()
        return err
    }

    // Xóa các chi tiết không có trong danh sách cập nhật
    for _, currentDetail := range currentDetails {
        if _, exists := updateMaSPs[currentDetail.MaSP]; !exists {
            var sanPham bonsai.BonsaiItem
            if err := tx.Where("masp = ?", currentDetail.MaSP).First(&sanPham).Error; err != nil {
                tx.Rollback()
                return err
            }
            sanPham.SoLuong += currentDetail.SoLuong
            if err := tx.Save(&sanPham).Error; err != nil {
                tx.Rollback()
                return err
            }

            if err := tx.Where("masp = ? AND mahd = ?", currentDetail.MaSP, updateData.MaHD).Delete(&model.ChiTietHoaDon{}).Error; err != nil {
                tx.Rollback()
                return err
            }
        }
    }

    // Cập nhật hoặc thêm mới các chi tiết phiếu nhập
    for _, data := range updateData.ChiTiet {
        var existingDetail model.ChiTietHoaDon
        //Tìm 
        if err := tx.Where("masp = ? AND mahd = ?", data.MaSP, updateData.MaHD).First(&existingDetail).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                data.MaHD = updateData.MaHD
                if err := tx.Create(&data).Error; err != nil {
                    tx.Rollback()
                    return err
                }

                var sanPham bonsai.BonsaiItem
                if err := tx.Where("masp = ?", data.MaSP).First(&sanPham).Error; err != nil {
                    tx.Rollback()
                    return err
                }
                sanPham.SoLuong -= data.SoLuong
                if err := tx.Save(&sanPham).Error; err != nil {
                    tx.Rollback()
                    return err
                }
            } else {
                tx.Rollback()
                return err
            }
        } else {
            existingDetail.SoLuong = data.SoLuong
            if err := tx.Where("masp = ? AND mahd = ?", data.MaSP, updateData.MaHD).Updates(&existingDetail).Error; err != nil {
                tx.Rollback()
                return err
            }
        }

        // Nếu TrangThai = "Cancel", cập nhật bảng sản phẩm
        if updateData.TrangThai != nil && *updateData.TrangThai == model.Cancel {
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