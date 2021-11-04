package zInterface

type Message interface {
	GetMsgID() uint32
	GetMsgLen() uint32
	GetData() []byte
	SetMsgID(uint32)
	SetData([]byte)
	SetMsgLen(uint32)
}
