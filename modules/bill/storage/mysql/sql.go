package mysql

import "gorm.io/gorm"

type billRepository struct{
	db *gorm.DB
}
func  NewBillRepository(db *gorm.DB) *billRepository{
	return &billRepository{db: db}
}