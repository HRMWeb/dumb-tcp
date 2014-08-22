package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := flag.Int("port", 9000, "port to listen on")
	flag.Parse()
	log.Println("Reading input")
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Read %d bytes", len(in))
	log.Printf("Starting listener on %d", *port)
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{
		Port: *port,
	})
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn, in)
	}
}

func handle(conn *net.TCPConn, in []byte) {
	log.Printf("Connection from %s", conn.RemoteAddr())
	// Wait a second and return
	defer conn.Close()
	time.Sleep(time.Second)
	conn.Write(in)
}
