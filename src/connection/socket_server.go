package connection

import (
	"net"
	"bufio"
	"fmt"
	"os"
)

type Server struct {
	Connected bool
}

func listenToClient(conn net.Conn) {
	for {
		fmt.Print("Request: ")
		reader := bufio.NewReader(os.Stdin)
    	request, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(request + "\n"))
		if err != nil {
			fmt.Println("Client closed connection!")
			return
		}
		//TODO: write it so, that server gets picture data
	}
}

func (s Server) waitForClient(listener net.Listener) {
	for {
		conn,_ := listener.Accept()
		fmt.Println("Connection established!")
		s.Connected = true
		listenToClient(conn)
	}
}

func (s Server) Start(port string) {
	fmt.Println("Listening for connection on port 8081...")
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error at listening on port " + port)
		return
	}
	s.waitForClient(ln)
}