package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/customer/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetCustomer(ctx context.Context, cond map[string]interface{}) (*model.Customer, error){

	var data model.Customer

	if err := s.db.Where(cond).Find(&data).Error; err!=nil{
		if err == gorm.ErrRecordNotFound{
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}