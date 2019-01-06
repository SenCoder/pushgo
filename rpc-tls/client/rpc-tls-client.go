package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/rpc"
)

var (
	certFile = "../client1.crt" //"../certs/client/client1.crt"
	keyFile  = "../client.key"  //"../certs/client/client.key"
	addr     = "127.0.0.1:8000"
)

func main() {

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}

	if len(cert.Certificate) != 2 {
		log.Fatal("client.crt should have 2 concatenated certificates: client + CA")
	}
	ca, err := x509.ParseCertificate(cert.Certificate[1])
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(ca)

	config := tls.Config{
		InsecureSkipVerify: false,
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
	}

	conn, err := tls.Dial("tcp", addr, &config)

	// conn, err := tcp.Dial("tcp", addr, byte(0)

	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}

	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())
	rpcClient := rpc.NewClient(conn)
	res := new(Result)
	if err := rpcClient.Call("Foo.Bar", "twenty-three ...", &res); err != nil {
		log.Fatalln("Failed to call RPC", err)
	}
	log.Printf("Returned result is %d", res.Data)
}

type Result struct {
	Data int
}
