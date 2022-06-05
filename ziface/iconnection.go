package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnId() uint32
	RemoteAddr() net.Addr
	SendMsg(uint32, []byte) error
	SetProperty(key string, val interface{})
	GetProperty(key string) (interface{}, error)
	RemoveProperty(key string)
}

type HandleFunc func(*net.TCPConn, []byte, int) error
