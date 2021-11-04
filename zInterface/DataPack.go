package zInterface

type DataPack interface {
	// GetHeadLen 获取包头的长度方法
	GetHeadLen() uint32
	// Pack 封包方法
	Pack(msg Message) ([]byte, error)
	// Unpack 拆包方法
	Unpack([]byte) (Message, error)
}
