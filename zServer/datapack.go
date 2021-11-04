package zServer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"go-zinx/zInterface"
	"go-zinx/zutils"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	//Datalen uint32 (4)+ ID uint32 (4)
	return 8
}

func (d *DataPack) Pack(msg zInterface.Message) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	//将dataLen写进buffer中
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	//将MsgID写进buffer中
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	//将data数据写进buffer中
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (d *DataPack) Unpack(data []byte) (zInterface.Message, error) {
	buffer := bytes.NewReader(data)
	msg := &Message{}
	if err := binary.Read(buffer, binary.LittleEndian, &msg.MsgLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	//判断datalen是否超出预期
	if zutils.GlobalObject.MaxPackageSize > 0 && msg.MsgLen > zutils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too long")
	}
	return msg, nil
}
