package models

import "gorm.io/gorm"

type ConfigBasic struct {
	gorm.Model
	Key   string `gorm:"column:key; type:varchar(255);" json:"key"`
	Value string `gorm:"column:value; type:varchar(255);" json:"value"`
}

func (table ConfigBasic) TableName() string {
	return "config_basic"
}
