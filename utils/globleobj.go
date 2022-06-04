package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

type GlobalObj struct {
	TCPServer ziface.IServer
	Host      string
	TCPPort   int
	Name      string

	Version          string
	MaxConn          int
	MaxPacketSize    uint32
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V1.0",
		TCPPort:          8080,
		Host:             "127.0.0.1",
		MaxConn:          1000,
		MaxPacketSize:    4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	GlobalObject.Reload()
}
