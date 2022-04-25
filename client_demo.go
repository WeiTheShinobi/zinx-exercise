package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start conn...")

	conn, err := net.Dial("tcp", "golang.org:http")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		time.Sleep(time.Second)
		_, err := conn.Write([]byte("Hello!"))
		if err != nil {
			fmt.Println("err :", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s\n", buf[:cnt])
	}
}
