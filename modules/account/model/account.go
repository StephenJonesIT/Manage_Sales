package model

import (
	"manage_sales/common"
	"time"
)
var Entity = "Tài khoản"
type AccountItem struct {
	common.LoginRequest
	NgayTao     time.Time `json:NgayTao gorm:"column:ngay_tao;"`
	Loai 		*AccountType `json:"LoaiTaiKhoan" gorm:"column:loai;"`
}

func(AccountItem) TableName() string{
	return "tai_khoan"
}

type AccountCreate struct{
	common.LoginRequest
	Loai *AccountType `json:"LoaiTaiKhoan, omitempty" gorm:"column:loai;"`
}

func(AccountCreate) TableName() string{
	return AccountItem{}.TableName()
}