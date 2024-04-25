package main

import (
	"fmt"
	"net"

)

func main() {
	listener, err := net.Listen("tcp","localhost:1234")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer listener.Close()
	for{
		clientConn, err := listener.Accept()
		if err != nil{
			fmt.Println(err)
			return
		}

		go func(client net.Conn){
			fmt.Print("Accepted")
			for{
				message, err := Decode(client)
				if err != nil{
					fmt.Println(err)
					return 
				}
				fmt.Println(string(message.Bytes()))
			}
		}(clientConn)
	}
	





}
