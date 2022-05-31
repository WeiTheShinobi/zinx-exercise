package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func NewMsgPackage(Id uint32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		Id:      Id,
		Data:    data,
	}
}

func (msg *Message) GetMsgLen() uint32 {
	return msg.DataLen
}

func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

func (msg *Message) GetMsg() []byte {
	return msg.Data
}

func (msg *Message) SetMsgLen(len uint32) {
	msg.DataLen = len
}

func (msg *Message) SetMsgId(msgID uint32) {
	msg.Id = msgID
}

func (msg *Message) SetMsg(data []byte) {
	msg.Data = data
}
