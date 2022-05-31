package znet

import (
	"fmt"
	"io"
	net "net"
	"testing"
	"time"
)

func TestDataPack(t *testing.T) {
	// server
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}

			go func(conn net.Conn) {

				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					io.ReadFull(conn, headData)

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						return
					}

					if msgHead.GetDataLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetDataLen())

						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							return
						}

						fmt.Println("==> Rec Msg: ID = ", msg.Id, " DataLen = ", msg.DataLen, " Data = ", string(msg.Data))
					}
				}

			}(conn)
		}
	}()

	//client
	go func() {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			return
		}

		dp := NewDataPack()

		msg1 := &Message{
			Id:      0,
			DataLen: 5,
			Data:    []byte{'h', 'e', 'l', 'l', 'o'},
		}

		sendData1, err := dp.Pack(msg1)
		if err != nil {
			fmt.Println("client pack msg1 err:", err)
			return
		}

		msg2 := &Message{
			Id:      1,
			DataLen: 4,
			Data:    []byte{'m', 's', 'g', '2'},
		}
		sendData2, err := dp.Pack(msg2)
		if err != nil {
			fmt.Println("client temp msg2 err:", err)
			return
		}

		sendData1 = append(sendData1, sendData2...)
		conn.Write(sendData1)

	}()

	select {
	case <-time.After(time.Second):
		return
	}
}
