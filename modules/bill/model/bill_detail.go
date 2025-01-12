package model
var EntityDetail = "ChiTietHoaDon"
type ChiTietHoaDon struct { 
	MaSP string `gorm:"column:masp" json:"MaSP"` 
	MaHD string `gorm:"column:mahd" json:"MaHD"` 
	SoLuong int `gorm:"column:so_luong" json:"SoLuong"` 
}

func (ChiTietHoaDon) TableName() string{
	return "chi_tiet_hd"
}