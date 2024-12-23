package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/suppliers/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.SupplierItemCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}