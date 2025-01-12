package mysql

import (
	"context"
	"manage_sales/modules/account/model"
)

func (sql *accountRepository) DeleteAccountByID(ctx context.Context, cond map[string]interface{}) error {

	if err := sql.db.Where(cond).Delete(&model.AccountItem{}).Error; err != nil{
		return err
	}

	return nil
}