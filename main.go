package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"strings"
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
	HttpVersion              = "HTTP/1.0"
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

func writeResponse(conn net.Conn, response Response) {
	responseString := fmt.Sprintf("%s %s\r\n", HttpVersion, response.Status)

	for key, value := range response.Headers {
		responseString += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	responseString += "\r\n" + response.Body

	_, err := conn.Write([]byte(responseString))
	if err != nil {
		log.Println("Error writing response:", err)
	}
}

func parseRequest(requestText string) *Request {
	lines := strings.Split(requestText, "\n")

	requestLine := strings.Split(lines[0], " ")
	method := requestLine[0]
	path := requestLine[1]

	headers := make(map[string]string)
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			break
		}
		header := strings.Split(line, ": ")
		key := header[0]
		value := header[1]
		headers[key] = value
	}

	body := ""
	if len(lines) > 0 {
		body = lines[len(lines)-1]
	}

	return &Request{
		Method:  method,
		Path:    path,
		Headers: headers,
		Body:    body,
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	requestText, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading request:", err)
		return
	}

	request := parseRequest(requestText)

	response := handleGetRequest(request)

	writeResponse(conn, response)
}

func handleGetRequest(request *Request) Response {
	
	if request.Path == "/" {
		response := Response{
			Status: StatusOK,
			Headers: map[string]string{
				ContentTypeHeader: ContentTypeTextPlain,
			},
			Body: "Resposta de uma requisição GET!",
		}
		return response
	} else if request.Path == "/clientes" {
		response := Response{
			Status: StatusOK,
			Headers: map[string]string{
				ContentTypeHeader: ContentTypeTextPlain,
			},
			Body: "Lista de clientes",
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

