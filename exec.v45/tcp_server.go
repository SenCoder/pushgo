package main

import (
	"net"
	"github.com/lunny/log"
)


func ConnHandler(conn *net.TCPConn) {
	defer conn.Close()

	buf := make([]byte, 32)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Error(err)
			break
		}
		log.Info(string(buf[:n]))
	}
}


func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "192.168.0.105:8080")
	listener, _ := net.ListenTCP("tcp", addr)

	for {
		conn, _ := listener.AcceptTCP()

		go ConnHandler(conn)
	}
}
