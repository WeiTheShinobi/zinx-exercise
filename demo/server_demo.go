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
	fmt.Println("preHandle")

}

func (r *DemoRouter) Handle(request ziface.IRequest) {
	fmt.Println("Handle")
	fmt.Println("recv from client: ID = ", request.GetMsgId(), ", data = ", string(request.GetDate()))

	err := request.GetConnection().SendMsg(1, []byte(time.Now().String()))
	if err != nil {
		fmt.Println(err)
	}
}

func (r *DemoRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("postHandle")
}

func main() {
	s := znet.NewServer()
	s.AddRouter(&DemoRouter{})
	s.Serve()
}
