package main

import (
	"net"
	"github.com/lunny/log"
	"bufio"
	"io"
	"fmt"
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

func ConnHandlerWithBuf(conn *net.TCPConn) {
	defer conn.Close()

	bufReader := bufio.NewReader(conn)

	_, ok := io.Reader(conn).(*bufio.Reader)
	if ok {
		fmt.Println("b is a implement of *bufio.Reader")
	}

	_, ok2 := io.Reader(conn).(*bufio.Reader)
	if ok2 {
		fmt.Println("b is a implement of *bufio.Reader")
	}

	for {
		_, err := bufReader.ReadByte()
		if err == io.EOF {
			fmt.Printf("client %s is close!\n", conn.RemoteAddr().String())
		}
		//在这里直接退出goroutine，关闭由defer操作完成
		return
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
