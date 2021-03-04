package socket

import (
	"fmt"
	"goSnippets/logger"
	"net"
	"sync"
	"time"
)

func init() {
	logger.DefaultLogger.Log()

	wg := &sync.WaitGroup{}

	go func() {
		wg.Add(1)
		// defer wg.Done()
		TCPServer(":9090")
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			TCPClient("127.0.0.1:9090")
		}()
	}
	wg.Done()
	wg.Wait()
}

func TCPClient(ipnport string) {
	addr, _ := net.ResolveTCPAddr("tcp", ipnport)
	conn, _ := net.DialTCP("tcp", nil, addr)

	data := make([]byte, 256)
	conn.Read(data)
	fmt.Println(string(data))
}

func TCPServer(port string) {
	addr, _ := net.ResolveTCPAddr("tcp", port)
	listener, _ := net.ListenTCP("tcp", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	timestamp := time.Now().String()
	conn.Write([]byte(timestamp))
	conn.Close()
}
