package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ls, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			panic(err)
		}
		go func(con net.Conn) {
			for {
				data, _ := bufio.NewReader(con).ReadString('\n')

				if strings.Contains(data, "quit") {
					break
				}
				fmt.Println("data recieved:", data)
				con.Write([]byte("message received \n"))
			}

			con.Close()
			fmt.Println("connection closed")
		}(conn)
	}
}
