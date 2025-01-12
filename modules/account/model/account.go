package model

import (
	"manage_sales/common"
	"time"
)
var Entity = "Tài khoản"
type AccountItem struct {
	common.LoginRequest
	NgayTao     time.Time 		`json:"NgayTao" gorm:"column:ngay_tao"`
	MaNV 		string 			`json:"MaNV,omitempty" gorm:"column:manv"`
	Loai 		*AccountType 	`json:"LoaiTaiKhoan" gorm:"column:loai"`
}

func(AccountItem) TableName() string{
	return "tai_khoan"
}

type AccountCreate struct{
	common.LoginRequest
	Loai *AccountType   `json:"LoaiTaiKhoan,omitempty" gorm:"column:loai"`
	MaNV string 		`json:"MaNV,omitempty" gorm:"column:manv"`
}

func(AccountCreate) TableName() string{
	return AccountItem{}.TableName()
}