package models

import "gorm.io/gorm"

type ExecuteQueue struct {
	gorm.Model
	SoftId    int    `gorm:"column:soft_id;type:int(11);" json:"soft_id"`
	State     int    `gorm:"column:state;type:tinyint(1);" json:"state"`             // 执行状态，1-执行中 2-已完成
	LogPath   string `gorm:"column:log_path;type:varchar(255);" json:"log_path"`     // 日志路径
	ShellPath string `gorm:"column:shell_path;type:varchar(255);" json:"shell_path"` // shell 路径
}

func (table ExecuteQueue) TableName() string {
	return "execute_queue"
}
