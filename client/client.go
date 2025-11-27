package main

import (
	"log"
	"net"
	"os"
	"strings"
	"crypto/tls"

	"github.com/joho/godotenv"


)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	IP_SERVER := os.Getenv("IP_SERVER")
                
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, 
	}

	conn, err := tls.Dial("tcp", IP_SERVER, tlsConfig)
	if err != nil {
		log.Println("error al realizar el handshake", err) 
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
