package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:90")
	if err != nil {
		log.Fatal("Error connecting:", err)
	}
	defer conn.Close()

	for {
		_, err := conn.Write([]byte("ping\n"))
		if err != nil {
			log.Fatal("Error sending PING:", err)
		}
		fmt.Println("ping")

		var buffer []byte
		//buffer := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(time.Minute))

		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatal("Error receiving response:", err)
		}

		response := string(buffer[:n])
		if response != "pong" {
			log.Fatal("Received unexpected response:", err)
		}

		fmt.Println("pong")
		time.Sleep(time.Minute)
	}
}
