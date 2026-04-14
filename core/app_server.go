package core

import (
	"errors"
	"log"
	"net/http"
	"time"

	"example.com/t/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AppServer struct {
	Config *types.AppConfig
	Engine *gin.Engine
	Server *http.Server
}

func NewAppServer(appConfig *types.AppConfig) *AppServer {
	gin.SetMode(gin.ReleaseMode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	engine := gin.Default()
	return &AppServer{
		Config: appConfig,
		Engine: engine,
		//Server: server,
	}
}

func (s *AppServer) Middlewares(debug bool, log *logrus.Logger) {
	//允许跨域请求 API
	s.Engine.Use(corsMiddleware())
	// 静态资源
	s.Engine.Use(staticResourceMiddleware())
	//s.Engine.Use(authorizeMiddleware(s, client))
	// 参数处理
	s.Engine.Use(parameterHandlerMiddleware())
	// 错误处理
	s.Engine.Use(errorHandler)
	// 日志
	s.Engine.Use(GinLoggerMiddleware(log))
	// 添加静态资源访问
	s.Engine.Static("/static", s.Config.StaticDir)
	//InitSwagger(s.Engine)
}

func (s *AppServer) Run() {
	server := &http.Server{
		Addr:           s.Config.Listen,
		Handler:        s.Engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 注意：这里使用协程,
	// ListenAndServe 会阻塞 main 主协程，内部会启动一个循环，持续接受和处理连接
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}
