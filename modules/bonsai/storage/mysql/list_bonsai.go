package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bonsai/model"
)

func (s *sqlStore) ListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]model.BonsaiItem, error) {

	var result []model.BonsaiItem

	db := s.db.Where("trang_thai <> ?", "Ngưng")

	if f := filter; f != nil {
		if v := f.TrangThai; v != "" {
			db = db.Where("trang_thai = ?", v)
		}
	}

	if err := db.Table(model.BonsaiItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("masp desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}