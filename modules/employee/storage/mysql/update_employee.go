package mysql

import (
	"context"
	"manage_sales/modules/employee/model"
)

func (s *sqlStore) UpdateEmployee(ctx context.Context, cond map[string]interface{}, dataUpdate *model.EmployeeUpdate) error {

	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}