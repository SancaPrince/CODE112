package main

import (
	"io"
	"net"

)

func main() {
	listener, err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil{
		panic("Something wrong")
	}
	for{
		clientProxy, err := listener.Accept()
		if err != nil{
			panic("Something wrong")

		}
		go handleProxy(clientProxy)
	}
}

func handleProxy(clientProxy net.Conn){
	serverProxy, err := net.Dial("tcp","127.0.0.1:1577")
	if err != nil{
		panic("Something wrong")
	}
	_, err = io.Copy(serverProxy, clientProxy)
	if err != nil{
		panic("Something wrong")
	}
	go func() {
		_, err := io.Copy(clientProxy, serverProxy)
		if err != nil{
			panic("Something wrong")
		}
	}()
	
}