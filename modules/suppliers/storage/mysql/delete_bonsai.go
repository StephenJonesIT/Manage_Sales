package mysql

import (
	"context"
	"manage_sales/modules/suppliers/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deleteStatus := model.BonsaiStatusDeleted
	if err := s.db.Table(model.SupplierItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"trang_thai": deleteStatus.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}
