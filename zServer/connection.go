package zServer

import (
	"errors"
	"github.com/sirupsen/logrus"
	"go-zinx/zutils"
	"io"
	"net"
)

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//连接ID
	ConnID uint32
	//连接状态,表示是否关闭
	ConnState bool
	//告知当前连接已经退出/停止的channel
	ExitChan chan bool
	//无缓冲管道，用于读写goroutine之间的消息通信
	MsgChan chan []byte
	//消息管理
	MsgHandler *MsgHandler
}

func NewConnetion(conn *net.TCPConn, connID uint32, msgHandler *MsgHandler) *Connection {
	return &Connection{
		Conn:       conn,
		ConnID:     connID,
		ConnState:  false,
		ExitChan:   make(chan bool, 1),
		MsgHandler: msgHandler,
		MsgChan:    make(chan []byte),
	}
}

func (conn *Connection) SendMsg(msgID uint32, data []byte) error {
	if conn.ConnState == true {
		return errors.New("Connection closed when send msg")
	}
	dp := NewDataPack()
	packMsg, err := dp.Pack(NewMsgPackage(msgID, data))
	if err != nil {
		logrus.Error("Pack error msg")
		return errors.New("Pack error msg")
	}
	conn.MsgChan <- packMsg
	return nil
}

func (conn *Connection) StartReader() {
	logrus.Info("Reader Goroutine is running...")
	defer logrus.Info("connID = ", conn.ConnID, ", reader is exit, remote addr is ", conn.Conn.RemoteAddr().String())
	defer conn.Stop()

	for {
		dp := NewDataPack()
		data := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn.GetTCPConnection(), data); err != nil {
			logrus.Error("read error :", err)
			break
		}
		msg, err := dp.Unpack(data)
		if err != nil {
			logrus.Error("unpack error :", err)
			break
		}
		var bytes []byte
		if msg.GetMsgLen() > 0 {
			bytes = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn.GetTCPConnection(), bytes); err != nil {
				logrus.Error("read error :", err)
				break
			}
		}
		msg.SetData(bytes)
		//得到当前conn数据的Request请求数据
		req := Request{
			conn: conn,
			msg:  msg,
		}

		if zutils.GlobalObject.WorkerPoolSize > 0 {
			conn.MsgHandler.SendMsgToTaskQueue(req)
		} else {
			//从路由中，找到注册绑定的Conn对应的router调用
			go conn.MsgHandler.DoMsgHandler(&req)
		}
	}

}

func (conn *Connection) StartWriter() {
	logrus.Info("Writer Goroutine is running")
	defer logrus.Info(conn.RemoteAddr().String(), "writer exit")

	for {
		select {
		case data := <-conn.MsgChan:
			if _, err := conn.Conn.Write(data); err != nil {
				logrus.Error("send data error", err)
				return
			}
		case <-conn.ExitChan:
			return
		default:
		}
	}
}

func (conn *Connection) Start() {
	logrus.Info("Conn Start().. ConnID = ", conn.GetConnID())
	//启动当前连接的读数据的业务
	go conn.StartReader()
	go conn.StartWriter()
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
	conn.ExitChan <- true
	close(conn.ExitChan)
	close(conn.MsgChan)
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
