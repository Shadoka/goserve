package connection

import (
	"net"
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
)

type Server struct {
	Connected bool
}

func (s Server) RequestPicture(conn net.Conn) {
	for {
		fmt.Print("Request: ")
		reader := bufio.NewReader(os.Stdin)
    	request, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(request + "\n"))
		if err != nil {
			fmt.Println("Client closed connection!")
			return
		}
		data := make([]byte, 0, 1000000)
		n, _ := bufio.NewReader(conn).Read(data)
		fmt.Println("Read " + string(n) + " bytes")
		fmt.Println("Saving picture as temp.jpg")
		ioutil.WriteFile("temp.jpg", data, 0600)
	}
}

func (s Server) waitForClient(listener net.Listener) {
	for {
		conn,_ := listener.Accept()
		fmt.Println("Connection established!")
		s.Connected = true
		s.RequestPicture(conn)
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