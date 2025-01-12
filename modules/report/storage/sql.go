package storage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewSQLReport(db *gorm.DB) *sqlStore{
	return &sqlStore{db:db}
}