package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type DemoRouter struct {
	znet.BasicRouter
}

func (r *DemoRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("preHandle")
	request.GetConnection().GetTCPConnection().Write([]byte("ping"))
}

func (r *DemoRouter) Handle(request ziface.IRequest) {
	fmt.Println("Handle")
	request.GetConnection().GetTCPConnection().Write([]byte("mid"))
}

func (r *DemoRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("postHandle\n")
	request.GetConnection().GetTCPConnection().Write([]byte("pong\n"))
}

func main() {
	s := znet.NewServer()
	s.AddRouter(&DemoRouter{})
	s.Serve()
}
