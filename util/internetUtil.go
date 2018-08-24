package util

import (
	"net"
	"bytes"
	"fmt"
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
	return nil,nil
}

func decode(data []byte) (string) {
	err := decodeCheckCode(data)
	checkError(err)
	length := len(data)
	err = decodeDataLength(data[0], length)
	checkError(err)
	err = decodeCommandIdentifier(data[1])
	checkError(err)
	err = decodeSensorName(data[1:4])
	checkError(err)
	sensorModuleAddress := decodeSensorModuleAddress(data[5])
	fmt.Println("sensorModuleAddress:",sensorModuleAddress)
	sensorData := decodeSensorData(data[6:length-2])
	fmt.Println("sensorData:",sensorData)
	return ""
}

func decodeDataLength(dataLength byte, dataRealLength int) (error) {
	if int(dataLength) != (dataRealLength - 2) {
		return &decodeError{"数据长度错误"}
	}
	return nil
}

func decodeCommandIdentifier(commandIdentifier byte) (error) {
	//TODO: 添加命令标识符处理逻辑
	if commandIdentifier == 0xF1 {
		//基板：Read读取命令			传感器模块：发送数据命令
	} else if commandIdentifier == 0xF2 {
		//基板：Test测试命令			传感器模块：接收数据命令
	} else if commandIdentifier == 0xF3 {
		//基板：Test测试确认命令		传感器模块：发送一次数据命令
	} else if commandIdentifier == 0xF4 {
		//基板：查询传感器模块地址	传感器模块：发送传感器地址
	} else {
		return &decodeError{"未能识别命令标识符！"}
	}
	return nil
}

func decodeSensorName(sensorName []byte) (error) {
	//TODO
	return nil
}

func decodeSensorModuleAddress(sensorModuleAddress byte) (int) {
	return int(sensorModuleAddress)
}

func decodeSensorData(sensorData []byte) ([]int) {
	sensorDataLength := int(sensorData[0])
	sensorDataValue := make([]int, sensorDataLength/2)
	for i := 1; i < sensorDataLength; i = i + 2 {
		sensorDataValue[(i-1)/2] = int(sensorData[i])*4096 + int(sensorData[i+1])
	}
	return sensorDataValue
}

func decodeCheckCode(data []byte) (error) {
	checkLength := 0
	length := len(data)
	checkCode := int(data[length-1])
	for i := 0; i < length-2; i++ {
		checkLength = checkLength + int(data[i])
	}
	if checkLength != checkCode {
		return &decodeError{"校验错误"}
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type decodeError struct {
	info string
}

func (e *decodeError) Error() string {
	return e.info
}
