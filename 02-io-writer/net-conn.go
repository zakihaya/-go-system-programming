package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "yahoo.co.jp:80")
	if err != nil {
		panic(err)
	}
	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: yahoo.co.jp\r\n\r\n")
	io.Copy(os.Stdout, conn)
	io.MultiWriter()
}
