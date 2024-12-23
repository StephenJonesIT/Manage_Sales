package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/employee/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetEmployee(ctx context.Context, cond map[string]interface{}) (*model.Employee, error){
	var data model.Employee

	if err := s.db.Where(cond).Find(&data).Error; err!=nil{
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}