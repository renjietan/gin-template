swag init --parseDependency --parseInternal

### 调试
```
go install github.com/go-delve/delve/cmd/dlv@latest
```

### air: 自动重载和自动更新 Swagger 文档
```
go install github.com/air-verse/air@latest
go get -u github.com/air-verse/air@latest
```
### fastTemplate: 模板字符串
```
go get -u github.com/valyala/fasttemplate
```
### swagger API文档
```
- 全局安装：
    - go install github.com/swaggo/swag/cmd/swag@latest
- 项目内安装
    - go get -u github.com/swaggo/swag/cmd/swag
    - go get -u github.com/swaggo/gin-swagger
    - go get -u github.com/swaggo/files
```
### fx 依赖注入
```
go get -u go.uber.org/fx
```
### 日志相关
```aiignore
# 文件切分
github.com/lestrrat-go/file-rotatelogs
# 日志插件
github.com/sirupsen/logrus
# 日志美化
go get github.com/x-cray/logrus-prefixed-formatter
```