package mysql

import (
	"context"
	"manage_sales/modules/bonsai/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.BonsaiItemUpdate) error {

	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}