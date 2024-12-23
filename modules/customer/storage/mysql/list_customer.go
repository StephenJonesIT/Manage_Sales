package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/customer/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]model.Customer, error) {

	var result []model.Customer
	db := s.db
	if f := filter; f != nil {
		if v := f.LoaiKH; v != "" {
			db = db.Where("loai_kh = ?", v)
		}
	}
	if err := s.db.Table(model.Customer{}.TableName()).Count(&paging.Total).Error; err != nil {
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