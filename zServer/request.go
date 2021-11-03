package zServer

import "go-zinx/zInterface"

type Request struct {
	//已经和客户端建立好的连接
	conn zInterface.Connection
	//客户端请求的数据
	data []byte
}

func (r *Request) GetConnection() zInterface.Connection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
