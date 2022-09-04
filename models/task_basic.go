package models

import "gorm.io/gorm"

type TaskBasic struct {
	gorm.Model
	Name      string `gorm:"column:name;type:varchar(255);" json:"name"`
	Spec      string `gorm:"column:spec;type:varchar(255);" json:"spec"`
	ShellPath string `gorm:"column:shell_path;type:varchar(255);" json:"shell_path"`
	LogPath   string `gorm:"column:log_path;type:varchar(255);" json:"log_path"`
}

func (table TaskBasic) TableName() string {
	return "task_basic"
}
