package main

import (
	"zinx/znet"
)

func main() {
	s := znet.NewServer("Demo")
	s.Serve()
}
