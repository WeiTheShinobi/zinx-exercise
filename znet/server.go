package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       int
	MsgHandler ziface.IMsgHandler
}

func NewServer() ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TCPPort,
		MsgHandler: NewMsgHandler(),
	}
	return s
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient ...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err : ", err)
		return errors.New("CallBackToClient error")
	}

	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server name : %s, Server version : %s, IP : %s:%d\n",
		utils.GlobalObject.Name, utils.GlobalObject.Version, utils.GlobalObject.Host, utils.GlobalObject.TCPPort)

	s.MsgHandler.StartWorkerPool()
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

		var cid uint32
		cid = 0

		for {
			conn, err := listener.AcceptTCP()

			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++

			go dealConn.Start()
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

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
}
