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
	//消息管理模块
	MsgHandler *MsgHandler
}

func (server *Server) AddRouter(msgID uint32, router zInterface.Router) {
	server.MsgHandler.AddRouter(msgID, router)
}

func NewServer() zInterface.Server {
	return &Server{
		Name:       zutils.GlobalObject.Name,
		IPversion:  "tcp4",
		IP:         zutils.GlobalObject.Host,
		Port:       zutils.GlobalObject.Port,
		MsgHandler: NewMsgHandler(),
	}
}

func (server *Server) Start() {
	go func() {
		server.MsgHandler.StartWorkerPool()
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
			dealConn := NewConnetion(conn, uid, server.MsgHandler)
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
