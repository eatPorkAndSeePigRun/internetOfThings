package util

import (
	"fmt"
	"strconv"
)

/**
	internetUtil.go调用
 */

func decode(data []byte) (string) {
	//进行校验码校对，若异常则抛出
	err := decodeCheckCode(data)
	checkError(err)
	//进行数据长度校对，若异常则抛出
	length := len(data)
	err = decodeDataLength(data[0], length)
	checkError(err)
	//TODO: 添加命令标识符逻辑
	err = decodeCommandIdentifier(data[1])
	checkError(err)
	//进行传感器名称、传感器模块地址、传感器数据
	sensorNameCode := decodeSensorName(data[1:4])
	sensorModuleAddress := decodeSensorModuleAddress(data[5])
	sensorData := decodeSensorData(data[6:length-2])
	fmt.Println("sensorNameCode:", sensorNameCode)
	fmt.Println("sensorModuleAddress:", sensorModuleAddress)
	fmt.Println("sensorData:", sensorData)
	return "decode success"
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

func decodeSensorName(sensorName []byte) (string) {
	sensorNameCode := strconv.Itoa(int(sensorName[0])) + strconv.Itoa(int(sensorName[1])) + strconv.Itoa(int(sensorName[2]))
	return sensorNameCode
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
