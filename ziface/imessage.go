package ziface

type IMessage interface {
	GetMsgId() uint32
	GetMsgLen() uint32
	GetMsg() []byte

	SetMsgId(uint32)
	SetMsgLen(uint32)
	SetMsg([]byte)
}
