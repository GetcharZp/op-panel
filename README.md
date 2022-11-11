# op-panel

> 后端：Echo、Gorm的运维面板
> 
> 前端：vue-element-admin

## golang 扩展安装

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

## vue-element-admin 安装

```text
git clone https://gitee.com/PanJiaChen/vue-element-admin.git
cd vue-element-admin
npm install
npm run dev
```

## 系统模块

- [ ] 首页
  - [x] 系统状态
- [x] 网站管理
  - [x] 网站列表
  - [x] 网站新增
  - [x] 网站修改
  - [x] 网站删除
- [x] 软件管理
  - [x] 软件列表
  - [x] 软件操作
- [x] 任务管理
  - [x] 定时任务
  - [x] 任务列表
  - [x] 新增任务
  - [x] 任务详情
  - [x] 修改任务
  - [x] 删除任务
- [x] 系统配置
  - [x] 初始化默认配置
  - [x] 修改系统配置
- [ ] 文件管理
  - [ ] 文件列表
  - [ ] 文件删除
  - [ ] 文件修改
- [x] 其他
  - [x] 登录