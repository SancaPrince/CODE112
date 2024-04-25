package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:1577")

	if err != nil {
		panic("Somethings Wrong")
	}
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic("Somethings Wrong")
		}
		go handleConnection(clientConn)
	}
}

func handleConnection(clientConn net.Conn) {
	var size uint32
	err := binary.Read(clientConn, binary.LittleEndian, &size)
	if err != nil {
		panic("Somethings Wrong")
	}

	bytMsg := make([]byte, size)

	_, err = clientConn.Read(bytMsg)
	if err != nil {
		panic("Somethings Wrong")
	}
	msg := string(bytMsg)
	//.txt
	reply := ""
	if strings.HasSuffix(msg,".txt"){
		reply = "File Recevied"
	}else{
		reply = "Message Recevied"
	}
	err = binary.Write(clientConn, binary.LittleEndian, uint32(len(msg)))
	if err != nil{
		panic(err)
	}
	_, err = clientConn.Write([]byte(reply))
	if err != nil{
		panic(err)
	} {
		
	}

	fmt.Printf("Received : %s\n", string(bytMsg))

}
