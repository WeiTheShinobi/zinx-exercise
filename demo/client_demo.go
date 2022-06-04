package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

func main() {
	fmt.Println("client start conn...")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("msg test")))
		if err != nil {
			fmt.Println(err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println(err)
			return
		}

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			return
		}

		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			return
		}

		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msgHead.GetMsgLen())

			io.ReadFull(conn, msg.Data)

			fmt.Println(msg.Id, " --- ", msg.GetMsgLen(), " --- ", string(msg.GetMsg()))
		}

		time.Sleep(time.Second)
	}
}
