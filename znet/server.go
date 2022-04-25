package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      8080,
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)

	go func() {

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			panic(err)
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept err : ", err)
				continue
			}

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("receive buf err : ", err)
						return
					}

					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write buf err : ", err)
						return
					}
				}
			}()
		}

	}()
}

func (s *Server) Stop() {
	panic("implement me")
}

func (s *Server) Serve() {
	s.Start()

	select {}
}
