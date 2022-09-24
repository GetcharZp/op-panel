package models

import "gorm.io/gorm"

type WebBasic struct {
	gorm.Model
	Name   string `gorm:"column:name;type:varchar(255);" json:"name"`     // 名称
	Domain string `gorm:"column:domain;type:varchar(255);" json:"domain"` // 域名
	Dir    string `gorm:"column:dir;type:varchar(255);" json:"dir"`       // 静态文件目录
}

func (table WebBasic) TableName() string {
	return "web_basic"
}
