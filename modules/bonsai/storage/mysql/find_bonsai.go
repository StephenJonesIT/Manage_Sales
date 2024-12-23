package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bonsai/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.BonsaiItem, error) {
	var data model.BonsaiItem

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}