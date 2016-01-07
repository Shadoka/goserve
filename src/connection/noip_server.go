package main

import (
	"net"
	"bufio"
	"fmt"
	"strings"
)

func listenToClient(conn net.Conn) {
	for {
		fmt.Print("Request: ")
		reader := bufio.NewReader(os.Stdin)
    	request, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(newText + "\n"))
		if err != nil {
			fmt.Println("Client closed connection!")
			return
		}
		//TODO: write it so, that server gets picture data
	}
}

func waitForClient(listener net.Listener) {
	for {
		conn,_ := listener.Accept()
		fmt.Println("Connection established!")
		listenToClient(conn)
	}
}

func main() {
	fmt.Println("Listening for connection on port 8081...")
	ln,_ := net.Listen("tcp", ":8081")
	waitForClient(ln)
}