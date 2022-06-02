package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn       *net.TCPConn
	ConnId     uint32
	MsgHandler ziface.IMsgHandler
	isClose    bool
	ExitChan   chan bool
}

func NewConnection(conn *net.TCPConn, connId uint32, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:       conn,
		ConnId:     connId,
		MsgHandler: msgHandler,
		isClose:    false,
		ExitChan:   make(chan bool, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")

	defer fmt.Println("connId = ", c.ConnId, " Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		dp := NewDataPack()

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			break
		}

		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println(err)
				break
			}
		}

		msg.SetMsg(data)
		request := Request{
			conn: c,
			msg:  msg,
		}

		go c.MsgHandler.DoMsgHandler(&request)
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClose {
		return errors.New("conn is closed")
	}

	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("pack err id = ", msgId)
		return errors.New("pack error msg")
	}

	if _, err := c.Conn.Write(binaryMsg); err != nil {
		return errors.New("conn write error")
	}

	return nil
}
