package main

import (
	"net"
	"net/http"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	request, err := http.NewRequest(
		"GET",
		"http://localhost:8888",
		nil)
	if err != nil {
		panic(err)
	}
	err = request.Write(conn)
}
