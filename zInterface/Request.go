package zInterface

type Request interface {
	// GetConnection 得到当前链接
	GetConnection() Connection
	// GetData 得到请求的消息数据
	GetData() []byte
}
