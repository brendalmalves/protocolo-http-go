package main

import (
	"fmt"
	"log"
	"net"
)

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    string
}

type Response struct {
	Status  string
	Headers map[string]string
	Body    string
}

const (
	StatusOK                  = "200 OK"
	StatusMethodNotAllowed   = "405 Method Not Allowed"
	ContentTypeHeader        = "Content-Type"
	ContentTypeTextPlain     = "text/plain"
	StatusNotFound           = "404 Not Found"
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


func handleGetRequest(request *Request) Response {
	
	if request.Path == "/" {
		response := Response{
			Status: StatusOK,
			Headers: map[string]string{
				ContentTypeHeader: ContentTypeTextPlain,
			},
			Body: "Hello, GET request!",
		}
		return response
	} else if request.Path == "/clients" {
		response := Response{
			Status: StatusOK,
			Headers: map[string]string{
				ContentTypeHeader: ContentTypeTextPlain,
			},
			Body: "list clients",
		}
		return response
	} else {
		response := Response{
			Status: StatusNotFound,
			Headers: map[string]string{
				ContentTypeHeader: ContentTypeTextPlain,
			},
			Body: "404 Not Found",
		}
		return response
	}
}

