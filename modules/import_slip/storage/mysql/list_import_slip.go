package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/import_slip/model"
)

func (sql *sqlStore) ListImportSlip(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]model.PhieuNhap, error) {
	var result []model.PhieuNhap
	db := sql.db.Where("trang_thai <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.TrangThai; v != "" {
			db = db.Where("trang_thai = ?", v)
		}
	}

	if err := db.Table(model.PhieuNhap{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("ngay_lap_pn desc").
		Preload("ChiTiet").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
