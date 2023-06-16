package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
	}
	
	fmt.Println("Servidor rodando")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleRequest(conn)
	} 
}

func handleRequest(conn net.Conn) {
	
}

