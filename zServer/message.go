package zServer

type Message struct {
	ID     uint32
	MsgLen uint32
	Data   []byte
}

func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		ID:     id,
		MsgLen: uint32(len(data)),
		Data:   data,
	}
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}

func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(u uint32) {
	m.ID = u
}

func (m *Message) SetData(bytes []byte) {
	m.Data = bytes
}

func (m *Message) SetMsgLen(u uint32) {
	m.MsgLen = u
}
