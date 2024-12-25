package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/account/model"
)

func (s *accountRepository) CreateAccount(ctx context.Context, data *model.AccountCreate) error{
	if err := s.db.Create(&data).Error; err != nil{
		return common.ErrCannotCreateEntity(model.Entity,err)
	}
	return nil
}