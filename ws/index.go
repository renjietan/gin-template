package ws

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"example.com/t/enum"
	"example.com/t/utility"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketManager WebSocket 管理器
type WebSocketManager struct {
	// 所有连接的客户端
	clients map[*websocket.Conn]bool
	// 用于广播消息的通道
	broadcast chan []byte
	// 用于注册新客户端
	register chan *websocket.Conn
	// 用于注销客户端
	unregister chan *websocket.Conn
	// 互斥锁保护 clients map
	mu sync.RWMutex
	// 升级器配置
	upgrader websocket.Upgrader
	closed   atomic.Bool
}

// NewWebSocketManager 创建并初始化 WebSocket 管理器
func NewWebSocketManager() *WebSocketManager {
	manager := &WebSocketManager{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte, 256),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			// 允许跨域
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	// 启动管理器协程
	go manager.run()

	return manager
}

// 客户端-新增
func (m *WebSocketManager) HandleWebSocket(c *gin.Context) {
	conn, err := m.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket 升级失败: %v", err)
		return
	}

	// 注册新客户端
	m.register <- conn

	// 启动读取协程
	go m.read(conn)
}

// 消息发送
func (m *WebSocketManager) SendToClient(conn *websocket.Conn, message []byte) error {
	m.mu.RLock()
	_, exists := m.clients[conn]
	m.mu.RUnlock()

	if !exists {
		return fmt.Errorf("客户端连接不存在")
	}

	return conn.WriteMessage(websocket.TextMessage, message)
}

// 广播
func (m *WebSocketManager) Broadcast(message []byte) {
	m.broadcast <- message
}

// GET 客户端数量
func (m *WebSocketManager) GetClientCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.clients)
}

// 断开连接
func (m *WebSocketManager) Close() {
	if !m.closed.CompareAndSwap(false, true) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	// 关闭所有客户端连接
	for conn := range m.clients {
		conn.Close()
		delete(m.clients, conn)
	}

	close(m.broadcast)
	close(m.register)
	close(m.unregister)
}

// 主循环
func (m *WebSocketManager) run() {
	for {
		select {
		case conn := <-m.register:
			// 注册新客户端
			m.mu.Lock()
			m.clients[conn] = true
			m.mu.Unlock()
			log.Printf("新客户端连接，当前连接数: %d", len(m.clients))

		case conn := <-m.unregister:
			// 注销客户端
			m.mu.Lock()
			if _, ok := m.clients[conn]; ok {
				delete(m.clients, conn)
				conn.Close()
				log.Printf("客户端断开连接，当前连接数: %d", len(m.clients))
			}
			m.mu.Unlock()

		case message := <-m.broadcast:
			// 广播消息给所有客户端
			m.mu.RLock()
			for conn := range m.clients {
				err := conn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Printf("发送消息失败: %v", err)
					// 如果发送失败，将连接加入注销队列
					select {
					case m.unregister <- conn:
					default:
					}
				}
			}
			m.mu.RUnlock()
		}
	}
}

// 读取客户端消息的协程（没写完 - 数据处理）
func (m *WebSocketManager) read(conn *websocket.Conn) {
	defer func() {
		m.unregister <- conn
		conn.Close()
	}()

	// 设置读取超时和 pong 处理
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	for {
		// 读取消息
		messageType, message, err := conn.ReadMessage()
		fmt.Println("消息类型：", messageType)
		// 检查是否是 正常规避 或 读取超时
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket 错误: %v", err)
			}
			break
		}
		// 任意消息都刷新读取超时，避免仅依赖控制帧 pong
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		// 下面 可以开始写我的逻辑了
		log.Printf("收到客户端消息: %s, \n", string(message))

		response := fmt.Sprintf("服务器收到: %s", string(message))
		// 向客户端发送消息
		if err := m.SendToClient(conn, []byte(response)); err != nil {
			log.Printf("发送回显消息失败: %v", err)
			break
		}
	}
}

// TODO：发送时 的数据处理（没写完）
func (m *WebSocketManager) handleRecv(message string) (s string, err any) {
	rec_json, err := utility.JsonStrToMap(string(message))
	if err != nil {
		return "", err
	}
	e := rec_json["event"]
	data_str := rec_json["data"]
	switch e {
	case enum.WS_EVENT_PING:
		str, err := utility.MapToJsonStr(map[string]any{
			"event": enum.WS_EVENT_PONG,
			"data":  data_str,
			"type":  enum.WS_TYPE_SERVER,
		})
		return str, err
	default:
		return utility.MapToJsonStr(map[string]any{
			"event": enum.WS_EVENT_UNKNOWN,
			"data":  data_str,
			"type":  enum.WS_TYPE_SERVER,
		})
	}
}
