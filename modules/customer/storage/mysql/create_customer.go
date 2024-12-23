package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/customer/model"
)

func(s *sqlStore) CreateItem(ctx context.Context, data *model.CustomerCreate) error{
	if err := s.db.Create(&data).Error; err != nil{
		return common.ErrCannotCreateEntity(model.EntityName,err)
	}
	return nil
}