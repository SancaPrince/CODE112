package main

import (
	"fmt"
	"net"

)

func main(){
	serverConn, err := net.Dial("tcp","127.0.0.1:1234")

	if err != nil{
		fmt.Println(err)
		return
	}
	defer serverConn.Close()
		message := "Hello From Client"
		payload := Binary(message)
		payload.WriteTo(serverConn)
	
}