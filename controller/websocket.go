package controller

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// Upgrader WebSocket 升级器
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 生产环境应配置允许的域名
		return true
	},
}

// WebSocketController WebSocket 控制器基类
type WebSocketController struct {
	Controller
}

// Handle WebSocket 处理入口
func (c *WebSocketController) Handle(ctx echo.Context) error {
	ws, err := Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		response := c.HandleMessage(string(message))
		if err := ws.WriteMessage(messageType, []byte(response)); err != nil {
			break
		}
	}
	return nil
}

// HandleMessage 消息处理（子类覆盖）
func (c *WebSocketController) HandleMessage(msg string) string {
	return "echo: " + msg
}

// WebSocketManager WebSocket 连接管理器
type WebSocketManager struct {
	connections map[*websocket.Conn]bool
	broadcast   chan []byte
	register    chan *websocket.Conn
	unregister  chan *websocket.Conn
	mu          sync.RWMutex
}

// NewWebSocketManager 创建 WebSocket 管理器
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		connections: make(map[*websocket.Conn]bool),
		broadcast:   make(chan []byte, 256),
		register:    make(chan *websocket.Conn),
		unregister:  make(chan *websocket.Conn),
	}
}

// Run 运行管理器
func (m *WebSocketManager) Run() {
	for {
		select {
		case conn := <-m.register:
			m.mu.Lock()
			m.connections[conn] = true
			m.mu.Unlock()

		case conn := <-m.unregister:
			m.mu.Lock()
			if _, ok := m.connections[conn]; ok {
				delete(m.connections, conn)
				conn.Close()
			}
			m.mu.Unlock()

		case message := <-m.broadcast:
			m.mu.RLock()
			for conn := range m.connections {
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					conn.Close()
					delete(m.connections, conn)
				}
			}
			m.mu.RUnlock()
		}
	}
}

// Register 注册连接
func (m *WebSocketManager) Register(conn *websocket.Conn) {
	m.register <- conn
}

// Unregister 注销连接
func (m *WebSocketManager) Unregister(conn *websocket.Conn) {
	m.unregister <- conn
}

// Broadcast 广播消息
func (m *WebSocketManager) Broadcast(message []byte) {
	m.broadcast <- message
}

// GetConnections 获取所有连接
func (m *WebSocketManager) GetConnections() []*websocket.Conn {
	m.mu.RLock()
	defer m.mu.RUnlock()

	connections := make([]*websocket.Conn, 0, len(m.connections))
	for conn := range m.connections {
		connections = append(connections, conn)
	}
	return connections
}

// GetConnectionCount 获取连接数
func (m *WebSocketManager) GetConnectionCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.connections)
}
