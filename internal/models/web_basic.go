package models

import "gorm.io/gorm"

type WebBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`     // 网站唯一标识
	Name     string `gorm:"column:name;type:varchar(255);" json:"name"`            // 名称
	Domain   string `gorm:"column:domain;type:varchar(255);" json:"domain"`        // 域名
	Dir      string `gorm:"column:dir;type:varchar(255);" json:"dir"`              // 静态文件目录
	ConfPath string `gorm:"column:conf_path; type:varchar(255);" json:"conf_path"` // 配置文件位置
}

func (table WebBasic) TableName() string {
	return "web_basic"
}
