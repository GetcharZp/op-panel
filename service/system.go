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

func getDefaultUserBasic() *define.UserBasic {
	return &define.UserBasic{
		Name:     helper.RandomString(8),
		Password: helper.RandomString(8),
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

func InitUserConfig() *define.UserBasic {
	dub := getDefaultUserBasic()
	dubByte, _ := json.Marshal(dub)
	ub := new(define.UserBasic)
	cb := new(models.ConfigBasic)
	err := models.DB.Model(new(models.ConfigBasic)).
		Where("key = 'user'").
		Attrs(map[string]interface{}{"key": "user", "value": string(dubByte)}).
		FirstOrCreate(cb).Error
	if err != nil {
		panic("[INIT UserBasic ERROR] : " + err.Error())
	}
	err = json.Unmarshal([]byte(cb.Value), ub)
	if err != nil {
		panic("[UNMARSHAL ERROR] : " + err.Error())
	}
	return ub
}
