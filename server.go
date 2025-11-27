package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)



type clientInfo struct {
	date 	string 
	host 	string 
	ip 		string 
}

var telegramTokenId string 
var telegramChatId string


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	
	telegramTokenId = os.Getenv("TOKEN_ID")
	telegramChatId = os.Getenv("CHAT_ID")
	
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, err := tls.Listen("tcp", "0.0.0.0:4466", tlsConfig)
	if err != nil {
		log.Println(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go registerConn(conn)
	}
}

func registerConn(conn net.Conn) {
	clientInfo := clientInfo{
		date: time.Now().Format("2006-01-02 15:04:05"),
		host: getHostname(conn),
		ip: conn.RemoteAddr().String(),
	}
	fmt.Printf("Client Connected:\n\t%s\n\t%s\n\t%s\n", clientInfo.date, clientInfo.host, clientInfo.ip)
	sendInfo(clientInfo)
}

func getHostname(conn net.Conn) string {
	_, err := conn.Write([]byte(fmt.Sprintf("Hostname")))
	if err != nil {
		log.Println("error al enviar comando", err)
	}
	bufRespuesta := make([]byte, 1024)
	n, err := conn.Read(bufRespuesta)
	if err != nil {
		log.Println("Error al leer respuesta", err)
	}
	hostname := strings.TrimSpace(string(bufRespuesta[:n]))
	return hostname
}

func sendInfo(cI clientInfo) {
	bot, err := tgbotapi.NewBotAPI(telegramTokenId)
	if err != nil {
		log.Println(err)
	}
	telegramChatId, err := strconv.ParseInt(telegramChatId,10,64)
	if err != nil {
		log.Println(err)
	}

	msg := fmt.Sprintf("\nClient Connected:\n\tDate: %s \n\tHostname: %s \n\tIp: %s", cI.date, cI.host, cI.ip)
	sendMsg := tgbotapi.NewMessage(telegramChatId, msg)
	bot.Send(sendMsg)
}
