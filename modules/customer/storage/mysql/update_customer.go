package mysql

import (
	"context"
	"manage_sales/modules/customer/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.CustomerUpdate) error {

	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}