package core

import (
	"errors"
	"io"
	"net/http"
	"time"

	"example.com/t/types"
	"github.com/gin-gonic/gin"
)

type AppServer struct {
	Config *types.AppConfig
	Engine *gin.Engine
	Server *http.Server
}

func NewAppServer(appConfig *types.AppConfig) *AppServer {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	engine := gin.Default()
	//server := &http.Server{
	//	Addr:           appConfig.Listen,
	//	Handler:        engine,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	return &AppServer{
		Config: appConfig,
		Engine: engine,
		//Server: server,
	}
}

func (s *AppServer) Middlewares(debug bool) {
	//允许跨域请求 API
	s.Engine.Use(corsMiddleware())
	s.Engine.Use(staticResourceMiddleware())
	//s.Engine.Use(authorizeMiddleware(s, client))
	s.Engine.Use(parameterHandlerMiddleware())
	s.Engine.Use(errorHandler)
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
