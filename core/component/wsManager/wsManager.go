package wsManager

import (
	"fmt"
	"strings"

	"example.com/t/core"
	"example.com/t/types"
	"example.com/t/utility"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

type WsManager struct {
	Server  *melody.Melody
	Clients map[string]*melody.Session
	config  *types.AppConfig
	Http    *core.AppServer
}

func NewWsManager(config *types.AppConfig, server *core.AppServer) *WsManager {
	return &WsManager{
		Server:  melody.New(),
		Clients: make(map[string]*melody.Session),
		config:  config,
		Http:    server,
	}
}

func (ws *WsManager) InitRouter() {
	ws.Http.Engine.Any("/ws/:id", func(context *gin.Context) {
		err := ws.Server.HandleRequest(context.Writer, context.Request)
		if err != nil {
			return
		}
	})
}

func (ws *WsManager) HandleFunds() {
	ws.Server.HandleConnect(func(session *melody.Session) {
		paths := strings.Split(session.Request.RequestURI, "/")
		id := utility.Tern(len(paths) > 2, paths[2], "")
		session.Set("ws-id", id)
		err := session.Write([]byte("connected"))
		if err != nil {
			return
		}
	})
	ws.Server.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("disconnected", s)
	})
	ws.Server.HandleMessage(func(session *melody.Session, bytes []byte) {
		msg := string(bytes)
		id, _ := session.Get("ws-id")
		fmt.Println("msg", msg)
		fmt.Println("id", id)
	})
}
