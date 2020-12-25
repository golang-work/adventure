# golang & gin simple projects

该项目是使用golang+gin框架实现的简易版api业务

# 特性

- [x] RESTful API
- [x] gorm
- [x] gin
- [x] token验证 (jwt-go)
- [x] 灵活配置 (viper)
- [x] 日志记录 (zap)
- [x] 单元测试
- [x] abort报错机制
- [x] MVC模式

# 目录结构说明

```
├── api // 业务逻辑
├── bootstrap // 应用初始化
├── config // 配置文件
├── docs // 相关文档
├── foundation // 系统核心
├── models // 模型文件
├── provider // 服务注册
├── routes  // 路由定义
├── storage
│   ├── app     // 应用自身需要的资源文件
│   ├── assets  // 外部需要的表态资源
│   └── logs    // 日志文件
├── support  // 系统框架、全局变量等库文件
├── tests   // 测试文件
├── main.go // 主文件
```

# 安装与运行

## 环境要求

* go version =>  ~1.13
* mysql version => ~5.5

## 安装与运行

1. 新建一个数据库并修改数据库连接信息：`configs/local.yaml`
2. 直接运行：`go run main.go`，也可以编译后再运行

# 项目已实现的api列表

```
[GIN-debug] POST   /api/account/sign-up      --> github.com/golang-work/adventure/api/controllers/v1.(*accountController).SignUp-fm (4 handlers)
[GIN-debug] POST   /api/account/sign-in      --> github.com/golang-work/adventure/api/controllers/v1.(*accountController).SignIn-fm (4 handlers)
[GIN-debug] GET    /api/outside/login-server/list --> github.com/golang-work/adventure/api/controllers/v1.(*outsideController).ListLoginServer-fm (4 handlers)
[GIN-debug] PUT    /api/account/reset-password --> github.com/golang-work/adventure/api/controllers/v1.(*accountController).ResetPassword-fm (5 handlers)
[GIN-debug] POST   /api/account/retrieve-password --> github.com/golang-work/adventure/api/controllers/v1.(*accountController).RetrievePassword-fm (5 handlers)
[GIN-debug] GET    /api/sub-account          --> github.com/golang-work/adventure/api/controllers/v1.(*subAccountCrud).List-fm (5 handlers)
[GIN-debug] POST   /api/sub-account          --> github.com/golang-work/adventure/api/controllers/v1.(*subAccountCrud).Store-fm (5 handlers)
[GIN-debug] DELETE /api/sub-account          --> github.com/golang-work/adventure/api/controllers/v1.(*subAccountCrud).Destroy-fm (5 handlers)
[GIN-debug] PUT    /api/sub-account/recover  --> github.com/golang-work/adventure/api/controllers/v1.(*subAccountCrud).Recover-fm (5 handlers)
```

## 未完待续...

- [ ] 上传资源到远程CDN
- [ ] 实现支付
- [ ] ...

