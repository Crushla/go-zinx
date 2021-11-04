package zServer

import "go-zinx/zInterface"

type Request struct {
	//已经和客户端建立好的连接
	conn zInterface.Connection
	//客户端请求的数据
	msg zInterface.Message
}

func (r *Request) GetConnection() zInterface.Connection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}

func (r *Request) GetMsgLen() uint32 {
	return r.msg.GetMsgLen()
}
