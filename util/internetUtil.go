package util

import (
	"net"
	"bytes"
)

func HandleSensorData(address string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
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
	data, err := readData(conn)
	checkError(err)
	msg := decode(data)
	// TODO: 如需发信号
	conn.Write([]byte(msg))
}

func readData(conn net.Conn) ([]byte, error) {
	var data, result []byte
	for true {
		_, err := conn.Read(data)
		if err != nil {
			return nil, err
		}
		for _, v := range data {
			temp := []byte{v}
			if !bytes.Equal(temp, []byte("\xAA")) {
				result = append(result, v)
			} else {
				return result[2:], nil
			}
		}
	}
	return nil, nil
}
