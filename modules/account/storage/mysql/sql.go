package mysql

import "gorm.io/gorm"

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepository{
	return &accountRepository{db : db}
}