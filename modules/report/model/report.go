package model

import "time"

type Report struct {
	MaBaoCao 	string 		`json:"MaBaoCao" gorm:"column:ma_bao_cao"`
	NgayTao  	*time.Time 	`json:"NgayTao" gorm:"column:ngay_tao"`
	DoanhThu 	int 		`json:"DoanhThu" gorm:"column:doanh_thu"`
	ChiPhi 		int 		`json:"ChiPhi" gorm:"column:chi_phi"`
}

func(Report) TableName() string{
	return "bao_cao"
}