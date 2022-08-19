package service

import (
	"encoding/json"
	"op-panel/define"
	"op-panel/helper"
	"op-panel/models"
)

func getDefaultSystemConfig() *define.SystemConfig {
	return &define.SystemConfig{
		Port:  ":1888",
		Entry: "/" + helper.RandomString(8),
	}
}

func GetSystemConfig() *define.SystemConfig {
	sc := new(define.SystemConfig)
	cb := new(models.ConfigBasic)
	dsc := getDefaultSystemConfig()
	dscByte, _ := json.Marshal(dsc)
	err := models.DB.Model(new(models.ConfigBasic)).
		Where("key = 'system'").
		Attrs(map[string]interface{}{"key": "system", "value": string(dscByte)}).
		FirstOrCreate(cb).Error
	if err != nil {
		panic("[INIT SYSTEM_CONFIG ERROR] : " + err.Error())
	}
	err = json.Unmarshal([]byte(cb.Value), sc)
	if err != nil {
		panic("[UNMARSHAL ERROR] : " + err.Error())
	}
	return sc
}
