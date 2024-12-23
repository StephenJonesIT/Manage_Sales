package mysql

import (
	"context"
	"manage_sales/modules/bonsai/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deleteStatus := model.BonsaiStatusNgung
	if err := s.db.Table(model.BonsaiItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"trang_thai": deleteStatus.String(),
	}).Error; err != nil {
		return err
	}
	return nil
}
