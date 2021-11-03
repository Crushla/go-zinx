package zServer

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-zinx/zInterface"
	"go-zinx/zutils"
	"net"
)

type Server struct {
	//服务器名称
	Name string
	//服务器绑定的IP版本
	IPversion string
	//服务器监听IP
	IP string
	//服务器监听端口
	Port int
	//当前的Server添加一个router,Server注册的连接对应处理业务
	Router zInterface.Router
}

func (server *Server) AddRouter(router zInterface.Router) {
	logrus.Info("Add Router Success")
	server.Router = router
}

func NewServer(name string) zInterface.Server {
	return &Server{
		Name:      zutils.GlobalObject.Name,
		IPversion: "tcp4",
		IP:        zutils.GlobalObject.Host,
		Port:      zutils.GlobalObject.Port,
		Router:    nil,
	}
}

func (server *Server) Start() {
	go func() {
		// 获取TCP的Addr
		addr, err := net.ResolveTCPAddr(server.IPversion, fmt.Sprintf("%s:%d", server.IP, server.Port))
		if err != nil {
			logrus.Error("resolve tcp addr error: ", err)
			return
		}
		// 监听服务器地址
		listenner, err := net.ListenTCP(server.IPversion, addr)
		if err != nil {
			logrus.Error("listen tcp error: ", err)
			return
		}
		logrus.Info("start zinx server success,", server.Name, " ,Listenning ...")

		var uid uint32
		uid = 0
		// 阻塞等待客户端连接，并处理业务
		for {
			//如果有客户端连接，阻塞返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				logrus.Error("Accept err", err)
				continue
			}
			dealConn := NewConnetion(conn, uid, server.Router)
			uid++

			go dealConn.Start()
		}
	}()
}

func (server *Server) Stop() {

}

func (server *Server) Run() {
	server.Start()

	//Todo 额外业务

	//阻塞
	select {}
}
