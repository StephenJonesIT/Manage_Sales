package model
var EntityDetail = "ChiTietHoaDon"
type ChiTietHoaDon struct { 
	MaSP string `gorm:"column:masp" json:"masp"` 
	MaHD string `gorm:"index;column:mahd" json:"mahd"` 
	SoLuong int `gorm:"column:so_luong" json:"SoLuong"` 
}

func (ChiTietHoaDon) TableName() string{
	return "chi_tiet_hd"
}