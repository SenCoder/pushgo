package main

import (
	"net"
	"github.com/lunny/log"
	"fmt"
	"pushgo/base"
)

func example() {
	addr, _ := net.ResolveTCPAddr("tcp", "192.168.0.105:8080")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Error(err)
		return
	}

	defer func() {
		fmt.Println("==>> defer")
		conn.Close()
	}()

	for {
		//iReader := bufio.NewReader(os.Stdin)
		//b, _ := iReader.ReadBytes('\n')
		b := base.RandomString()
		conn.Write([]byte(b))
		//time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	example()
}