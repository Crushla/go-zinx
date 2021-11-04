package zInterface

import (
	"net"
)

type Connection interface {
	// Start 启动连接
	Start()
	// Stop 停止连接
	Stop()
	// GetTCPConnection 获取当前连接的绑定socket conn
	GetTCPConnection() *net.TCPConn
	// GetConnID GetID 获取当前连接的绑定ID
	GetConnID() uint32
	// RemoteAddr 获取远程客户端的TCP状态 IP Port
	RemoteAddr() net.Addr
	// SendMsg 发送数据，将数据发送给远程客户端
	SendMsg(msgID uint32, data []byte) error
}

type HandleFunc func(conn *net.TCPConn, data []byte, cnt int) error
