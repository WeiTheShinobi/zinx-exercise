package main

import (
	"fmt"
	"time"
	"zinx/ziface"
	"zinx/znet"
)

type DemoRouter struct {
	znet.BasicRouter
}

func (r *DemoRouter) PreHandle(request ziface.IRequest) {

}

func (r *DemoRouter) Handle(request ziface.IRequest) {
	fmt.Println("--------")
	fmt.Println("recv from client: ID = ", request.GetMsgId(), ", data = ", string(request.GetData()))
	fmt.Println(time.Now().String())
	fmt.Println(request.GetConnection().RemoteAddr())

	err := request.GetConnection().SendMsg(1, []byte(time.Now().String()))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--------")

	time.Sleep(time.Second)
}

func (r *DemoRouter) PostHandle(request ziface.IRequest) {
}

func main() {
	s := znet.NewServer()
	s.AddRouter(0, &DemoRouter{})
	s.Serve()
}
