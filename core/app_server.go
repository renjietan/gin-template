package core

import (
	"io"
	"log"
	"net/http"
	"time"

	"example.com/t/types"
	"github.com/gin-gonic/gin"
)

type AppServer struct {
	Config    *types.AppConfig
	Engine    *gin.Engine
	Server    *http.Server
	SysConfig *types.SystemConfig // system config cache
}

func NewAppServer(appConfig *types.AppConfig) *AppServer {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	engine := gin.Default()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("服务启动失败(app_server.go)：" + err.Error())
	}
	return &AppServer{
		Config: appConfig,
		Engine: engine,
		Server: server,
	}
}

func (s *AppServer) Init(debug bool) {
	//允许跨域请求 API
	s.Engine.Use(corsMiddleware())
	//s.Engine.Use(staticResourceMiddleware())
	//s.Engine.Use(authorizeMiddleware(s, client))
	s.Engine.Use(parameterHandlerMiddleware())
	s.Engine.Use(errorHandler)
	// 添加静态资源访问
	s.Engine.Static("/static", s.Config.StaticDir)
	//InitSwagger(s.Engine)
}

func (s *AppServer) Run() {

}
