package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":90")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	obj := map[string]string{"name": "Ibrahim", "ping": "pong", "age": "29"}

	for {
		n, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input := strings.TrimSpace(string(n))

		conn.Write([]byte(obj[input]))

	}
}
