package mysql

import (
	"context"
	"manage_sales/modules/employee/model"
)

func (s *sqlStore) DeleteEmployee(ctx context.Context, cond map[string]interface{}) error{
	deleteEmployee := model.EmployeeStatusDeleted

	if err := s.db.Table(model.Employee{}.TableName()).Where(cond).Updates(map[string] interface{}{
		"trang_thai":deleteEmployee.String(),
		}).Error; err != nil{
			return err
		} 
	return nil
}