package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnId   uint32
	Router   ziface.IRouter
	isClose  bool
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnId:   connId,
		Router:   router,
		isClose:  false,
		ExitChan: make(chan bool, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")

	defer fmt.Println("connId = ", c.ConnId, " Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, utils.GlobalObject.MaxPacketSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf err : ", err)
			return
		}

		request := &Request{
			conn: c,
			data: buf,
		}

		go func(r ziface.IRequest) {
			c.Router.PreHandle(r)
			c.Router.Handle(r)
			c.Router.PostHandle(r)
		}(request)

	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... Conn Id = ", c.ConnId)

	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... Conn Id = ", c.ConnId)

	if !c.isClose {
		err := c.Conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		c.isClose = true
		close(c.ExitChan)
	}
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	panic("implement me")
}
