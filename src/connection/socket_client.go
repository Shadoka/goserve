package connection

import (
	"fmt"
	"net"
	"bufio"
	"io/ioutil"
	"strings"
)

type Client struct {}

func (c Client) Start() {
	fmt.Println("Connecting to localhost...")
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	defer conn.Close()
	fmt.Println("Connection established!")
	for {
		fmt.Println("Waiting for request...")
		request,_ := bufio.NewReader(conn).ReadString('\n')
		trimmedRequest := strings.TrimSuffix(request, "\n")
		fmt.Println("Request from Server: " + trimmedRequest)
		processRequest(trimmedRequest, conn)
		fmt.Println("Request processed!")
	}
}

func processRequest(request string, conn net.Conn) {
	for _,v := range request {
		fmt.Println(v)
	}
	switch request {
	case "cat":
		loadAndSendPic(conn, "coolcat")
	case "dog":
		loadAndSendPic(conn, "cooldog")
	default:
		loadAndSendPic(conn, "notfound")
	}
}

func sendData(content []byte, conn net.Conn) {
	n, err := conn.Write(content)
	fmt.Print("written bytes: ")
	fmt.Println(n)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func loadAndSendPic(conn net.Conn, filename string) {
	fmt.Println("trying to load " + filename + " picture")
	content, err := ioutil.ReadFile("../../resources/" + filename + ".jpg")
	fmt.Println("Size of picture: " + string(len(content)))
	if err != nil {
		fmt.Println(err.Error())
	}
	sendData(content, conn)
}