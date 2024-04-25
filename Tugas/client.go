package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
	"strings"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Println("Welcome")
		fmt.Println("1.Send message to server")
		fmt.Println("2. exit")
		scanner.Scan()
		choose := scanner.Text()
		if choose == "1"{
			messageMenu()
		}else if choose == "2"{
			break;
		}
	}
}


func messageMenu()  {
	scanner := bufio.NewScanner(os.Stdin)
	var Message string
	for{
		fmt.Print("Insert Your message: ")
		scanner.Scan()
		Message = scanner.Text()
		if len(Message) < 10{
			fmt.Println("Message should be more than 10 char")
		}else if strings.Contains(Message,"kasar"){
			fmt.Println("no bad words allowed")
		}else if strings.Compare(Message,"hello world"){
			fmt.Println("gaboleh ini")
		}else{
			break
		}

	}

}

func sendMessagetoServer()  {
	net.
}