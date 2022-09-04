# op-panel

> 运维面板

## 扩展安装

```text
go get github.com/kataras/iris/v12@master
go get -u github.com/golang-jwt/jwt/v4
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get github.com/shirou/gopsutil
go get github.com/robfig/cron/v3
go get github.com/satori/go.uuid
```

## 系统模块

- [ ] 首页
  - [x] 系统状态
- [ ] 网站管理
- [ ] 软件管理
- [ ] 任务管理
  - [x] 定时任务
  - [x] 任务列表
  - [ ] 新增任务
  - [ ] 修改任务
- [x] 系统配置
  - [x] 初始化默认配置
  - [x] 修改系统配置
- [x] 其他
  - [x] 登录