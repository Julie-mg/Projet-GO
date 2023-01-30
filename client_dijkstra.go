package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func send_init(conn net.Conn, file *os.File) {
	// read matrice in file and send it to the server
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			io.WriteString(conn, fmt.Sprintf("\n\n"))
			break
		}
		fmt.Printf("%v\n", line)
		io.WriteString(conn, fmt.Sprintf("%v\n", line))
	}
}

func main() {

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		return
	}
	defer conn.Close()

	// open matrice
	file, err := os.Open("matriceA.txt")
	if err != nil {
		log.Fatal(err)
	}

	send_init(conn, file)

	for {
		// read data send from server (results Dijkstra)
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("error")
			fmt.Println(err)
			conn.Close()
			break
		} else {
			fmt.Printf(netData)
		}
		// server send 'fin' at the end of its transmission
		if netData == "fin\n" {
			conn.Close()
			break
		}

	}

}
