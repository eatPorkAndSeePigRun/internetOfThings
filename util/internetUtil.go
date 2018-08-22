package util

import (
	"net"
)

type InternetUtil struct {
}

func (iu InternetUtil) myMainThread() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:7070")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	// TODO: 接收解析传感器数据,并关闭
	conn.Write([]byte("some smg"))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
