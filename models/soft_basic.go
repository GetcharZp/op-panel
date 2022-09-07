package models

import "gorm.io/gorm"

type SoftBasic struct {
	gorm.Model
	Name           string `gorm:"column:name;type:varchar(255);" json:"name"`
	Desc           string `gorm:"column:desc;type:varchar(255);" json:"desc"`
	ShellInstall   string `gorm:"column:shell_install;type:varchar(255);" json:"shell_install"`
	ShellUninstall string `gorm:"column:shell_uninstall;type:varchar(255);" json:"shell_uninstall"`
	ShellStart     string `gorm:"column:shell_start;type:varchar(255);" json:"shell_start"`
	ShellStop      string `gorm:"column:shell_stop;type:varchar(255);" json:"shell_stop"`
	ShellRestart   string `gorm:"column:shell_restart;type:varchar(255);" json:"shell_restart"`
	State          int    `gorm:"column:state;type:tinyint(1);" json:"state"` // 0-未安装 1-运行中 -1-停止中
}

func (table SoftBasic) TableName() string {
	return "soft_basic"
}
