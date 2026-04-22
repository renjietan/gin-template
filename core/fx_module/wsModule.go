package fx_module

import (
	"example.com/t/api/ws"
	"go.uber.org/fx"
)

var FxWsModule = fx.Module("fx-ws-module",
	fx.Provide(ws.NewWebSocketManager),
	fx.Invoke(func(ws *ws.WebSocketManager) {
		ws.Run()
	}),
)
