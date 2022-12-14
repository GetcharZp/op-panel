package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	db, err := gorm.Open(sqlite.Open("op-panel.db"), &gorm.Config{})
	if err != nil {
		panic("[OPEN DB ERROR] : " + err.Error())
	}
	err = db.AutoMigrate(&ConfigBasic{}, &TaskBasic{}, &SoftBasic{}, &ExecuteQueue{}, &WebBasic{})
	if err != nil {
		panic("[MIGRATE ERROR] : " + err.Error())
	}
	DB = db
}
