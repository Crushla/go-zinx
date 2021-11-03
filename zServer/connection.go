package zServer

import (
	"github.com/sirupsen/logrus"
	"go-zinx/zInterface"
	"go-zinx/zutils"
	"net"
)

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//连接ID
	ConnID uint32
	//连接状态
	ConnState bool
	//告知当前连接已经退出/停止的channel
	ExitChan chan bool
	//该连接处理方法的Router
	Router zInterface.Router
}

func NewConnetion(conn *net.TCPConn, connID uint32, router zInterface.Router) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		ConnState: false,
		ExitChan:  make(chan bool, 1),
		Router:    router,
	}
}

func (conn *Connection) StartReader() {
	logrus.Info("Reader Goroutine is running...")
	defer logrus.Info("connID = ", conn.ConnID, ", reader is exit, remote addr is ", conn.Conn.RemoteAddr().String())
	defer conn.Stop()

	for {
		buf := make([]byte, zutils.GlobalObject.MaxPackageSize)
		_, err := conn.Conn.Read(buf)
		if err != nil {
			logrus.Error("recv buf err", err)
			continue
		}
		//得到当前conn数据的Request请求数据
		req := Request{
			conn: conn,
			data: buf,
		}
		//从路由中，找到注册绑定的Conn对应的router调用
		go func(request zInterface.Request) {
			conn.Router.PreHandler(request)
			conn.Router.Handler(request)
			conn.Router.PostHandler(request)
		}(&req)
	}

}

func (conn *Connection) StartWriter() {

}

func (conn *Connection) Start() {
	logrus.Info("Conn Start().. ConnID = ", conn.GetConnID())
	//启动当前连接的读数据的业务
	go conn.StartReader()

	// Todo 启动从当前连接写数据的业务
}

func (conn *Connection) Stop() {
	logrus.Info("Conn Stop().. ConnID = ", conn.GetConnID())
	if conn.ConnState == true {
		return
	}
	conn.ConnState = true
	err := conn.Conn.Close()
	if err != nil {
		logrus.Error("connection stop() conn.Conn.Close err:", err)
		return
	}
	close(conn.ExitChan)
}

func (conn *Connection) GetTCPConnection() *net.TCPConn {
	return conn.Conn
}

func (conn *Connection) GetConnID() uint32 {
	return conn.ConnID
}

func (conn *Connection) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

func (conn *Connection) Send(data []byte) error {
	return nil
}
