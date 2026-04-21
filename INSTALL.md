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

### NACOS
- 安装:
  go get -u github.com/nacos-group/nacos-sdk-go/v2
- 注意：
  下载完成后，需要执行 go mod tidy

docker run -d  -p 9848:9848 -p 7848:7848 -p 8848:8848 -e MODE=standalone -v E:\docker-volumes\nacos\init.d\custom.properties:/home/nacos/init.d\custom.properties -v E:\docker-volumes\nacos\logs:/home/nacos/logs --restart always --name nacos nacos/nacos-server