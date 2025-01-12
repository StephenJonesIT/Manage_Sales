package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/account/model"
)

func (s *accountRepository) ListItem(
	ctx context.Context,
	paging *common.Paging,
	morekeys ...string,
) ([]model.AccountItem, error) {

	var result []model.AccountItem

	if err := s.db.Table(model.AccountItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := s.db.Order("ngay_tao desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}