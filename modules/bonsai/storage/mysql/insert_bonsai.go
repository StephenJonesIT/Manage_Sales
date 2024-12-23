package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bonsai/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.BonsaiItemCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}