package service

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"log"
	"net/http"
	"op-panel/define"
	"op-panel/helper"
	"op-panel/models"
	"syscall"
	"time"
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

// UpdateSystemConfig 修改系统配置
func UpdateSystemConfig(c echo.Context) error {
	var (
		port  = c.FormValue("port")
		entry = c.FormValue("entry")
		cb    = new(models.ConfigBasic)
		sc    = new(define.SystemConfig)
	)
	// 获取现用配置
	err := models.DB.Where("key = 'system'").First(cb).Error
	if err != nil {
		log.Printf("[DB ERROR] : %v\n", err)
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
	}
	json.Unmarshal([]byte(cb.Value), sc)

	// 设置新的配置
	if port != "" {
		sc.Port = ":" + port
	}
	if entry != "" {
		sc.Entry = "/" + entry
	}
	scByte, _ := json.Marshal(sc)

	// 更新配置
	err = models.DB.Model(new(models.ConfigBasic)).Where("key = 'system'").Update("value", string(scByte)).Error
	if err != nil {
		log.Printf("[DB ERROR] : %v\n", err)
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, echo.Map{
		"code": 200,
		"msg":  "修改成功",
	})

	// 重启服务
	syscall.Kill(define.PID, syscall.SIGINT)
	return nil
}

// SystemState 获取系统状态
func SystemState(c echo.Context) error {
	var (
		cpuUsedPercent  float64
		memUsedPercent  float64
		diskUsed        uint64
		diskUsedPercent float64
	)
	// cpu
	cpuPercents, _ := cpu.Percent(time.Second, true)
	for _, percent := range cpuPercents {
		cpuUsedPercent += percent
	}
	cpuUsedPercent /= float64(len(cpuPercents))
	// mem
	vms, _ := mem.VirtualMemory()
	memUsedPercent = vms.UsedPercent
	// disk
	partitions, _ := disk.Partitions(true)
	for _, partition := range partitions {
		us, _ := disk.Usage(partition.Mountpoint)
		diskUsed += us.Used
	}
	allUsage, _ := disk.Usage("/")
	diskUsedPercent = float64(diskUsed) / float64(allUsage.Total) * 100
	return c.JSON(http.StatusOK, echo.Map{
		"cpu_used_percent":  fmt.Sprintf("%.2f", cpuUsedPercent),
		"mem_used_percent":  fmt.Sprintf("%.2f", memUsedPercent),
		"disk_used_percent": fmt.Sprintf("%.2f", diskUsedPercent),
	})
}
