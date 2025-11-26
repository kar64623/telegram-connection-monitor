package main

import (
	"log"
	"net"
	"os"
	"strings"
	
	"github.com/joho/godotenv"


)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	IP_SERVER := os.Getenv("IP_SERVER")

	conn, err := net.Dial("tcp", IP_SERVER)
	if err != nil {
		log.Println(err)
	}
	handleConn(conn)
}

func handleConn(conn net.Conn) {
	bufMensaje := make([]byte, 1024)
	n, err := conn.Read(bufMensaje)
	if err != nil {
		log.Println(err)
	}
	mensaje := strings.TrimSpace(string(bufMensaje[:n]))
	if mensaje == "Hostname" {
		hostname, err := os.Hostname()
		if err != nil {
			log.Println(err)
		}
		_, err = conn.Write([]byte(hostname))
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Error, comando no reconocido")
	}
}
