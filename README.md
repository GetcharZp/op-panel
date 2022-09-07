# op-panel

> 基于Echo、Gorm的运维面板

## 扩展安装

<del>go get github.com/kataras/iris/v12@master</del>
```text
go get github.com/labstack/echo/v4
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
  - [x] 软件列表
  - [ ] 软件安装
  - [ ] 软件状态修改
- [x] 任务管理
  - [x] 定时任务
  - [x] 任务列表
  - [x] 新增任务
  - [x] 修改任务
  - [x] 删除任务
- [x] 系统配置
  - [x] 初始化默认配置
  - [x] 修改系统配置
- [x] 其他
  - [x] 登录