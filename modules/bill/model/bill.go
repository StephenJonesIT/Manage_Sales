package model

import "manage_sales/common"

var Entity = "HoaDon"

type HoaDonRequest struct {
	common.HoaDon
	ChiTietHD []ChiTietHoaDon `gorm:"foreignKey:MaHD;references:MaHD" json:"ChiTietHD"`
}

