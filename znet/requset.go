package znet

import "zinx/ziface"

type Request struct {
	conn ziface.IConnection
	msg  ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetDate() []byte {
	return r.msg.GetMsg()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
