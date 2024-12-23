package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/employee/model"
)

func (s *sqlStore) CreateEmployee(context context.Context, data *model.EmployeeCreate) error{
	if err:= s.db.Create(&data).Error; err != nil{
		return common.ErrDB(err)
	}
	return nil
}