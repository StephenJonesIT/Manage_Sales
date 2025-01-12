package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/suppliers/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
	paging *common.Paging,
	morekeys ...string,
) ([]model.SupplierItem, error) {

	var result []model.SupplierItem

	db := s.db.Where("trang_thai <> ?", "Deleted")

	if err := db.Table(model.SupplierItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("ngay_tao desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}