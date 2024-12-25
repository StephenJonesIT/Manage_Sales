package mysql

import (
	"context"
	"manage_sales/modules/account/model"
)

func (sql *accountRepository) FindAccount(context context.Context, cond map[string]interface{}) (*model.AccountItem, error){
	var data model.AccountItem
	if err := sql.db.Where(cond).First(&data).Error; err != nil{
		return nil, err 
	}
	return &data, nil
}
