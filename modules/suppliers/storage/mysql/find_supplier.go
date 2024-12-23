package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/suppliers/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.SupplierItem, error) {
	var data model.SupplierItem

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}