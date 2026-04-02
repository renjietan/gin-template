package core

import (
	"io"

	"example.com/t/types"
	"github.com/gin-gonic/gin"
)

type AppServer struct {
	Config    *types.AppConfig
	Engine    *gin.Engine
	SysConfig *types.SystemConfig // system config cache
}

func NewAppServer(appConfig *types.AppConfig) *AppServer {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	return &AppServer{
		Config: appConfig,
		Engine: gin.Default(),
	}
}

func (s *AppServer) Init(debug bool) {
	// 允许跨域请求 API
	s.Engine.Use(corsMiddleware())
	s.Engine.Use(staticResourceMiddleware())
	//s.Engine.Use(authorizeMiddleware(s, client))
	s.Engine.Use(parameterHandlerMiddleware())
	s.Engine.Use(errorHandler)
	// 添加静态资源访问
	s.Engine.Static("/static", s.Config.StaticDir)
}

func (s *AppServer) Run() {

}
