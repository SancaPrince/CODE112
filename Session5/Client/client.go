package main

import (
	"encoding/binary"
	"net"
)

func main() {
	sendMessagetoServer("Hello World")
}

func sendMessagetoServer(msg string) string{
	serverConn, err := net.Dial("tcp", "127.0.0.1:1577")

	if err != nil {
		panic("Somethings Wrong")
	}
	err = binary.Write(serverConn, binary.LittleEndian, uint32(len(msg)))
	if err != nil {
		panic("Somethings Wrong")
	}
	_, err = serverConn.Write([]byte(msg))
	if err != nil {
		panic("Somethings Wrong")
	}

	var size uint32
	binary.Read(serverConn, binary.LittleEndian, &size)
	if err != nil{
		panic(err)
	}
	reply := make([]byte, size)
	_, err = serverConn.Read((reply))
	if err != nil{
		panic(err)
	}
	return string(reply)
}
