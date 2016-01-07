package connection

import (
	"fmt"
	"net"
	"bufio"
	"io/ioutil"
)

type Client struct {}

func (c Client) Start() {
	fmt.Println("Connecting to localhost...")
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	fmt.Println("Connection established!")
	for {
		fmt.Println("Waiting for request...")
		request,_ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Request from Server: " + request)
		processRequest(request, conn)
		fmt.Println("Request processed!")
	}
}

func processRequest(request string, conn net.Conn) {
	switch request {
	case "cat":
		sendCatPic(conn)
	case "dog":
		sendDogPic(conn)
	default:
		sendErrorPic(conn)
	}
}

func sendData(content []byte, conn net.Conn) {
	conn.Write(content)
}

func sendCatPic(conn net.Conn) {
	//TODO: check the number of read bytes, seems to be not reading at all
	content, err := ioutil.ReadFile("../resources/coolcat.jpg")
	if err != nil {
		fmt.Println(err.Error())
	}
	sendData(content, conn)
}

func sendDogPic(conn net.Conn) {
	content,_ := ioutil.ReadFile("../resources/cooldog.jpg")
	sendData(content, conn)
}

func sendErrorPic(conn net.Conn) {
	content,_ := ioutil.ReadFile("../resources/notfound.jpg")
	sendData(content, conn)
}