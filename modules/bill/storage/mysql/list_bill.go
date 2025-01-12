package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bill/model"
)

func (sql *billRepository) ListBill(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]model.HoaDon, error) {
	var result []model.HoaDon
	db := sql.db.Where("trang_thai <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.TrangThai; v != "" {
			db = db.Where("trang_thai = ?", v)
		}
	}

	if err := db.Table(model.HoaDon{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("ngay_lap_hd desc").
		Preload("ChiTiet").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
