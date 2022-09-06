package main

import (
	"log"
	"os"
)

type JsonMessage struct {
	ChatID  string `json:"chat_id"`
	Message string `json:"message"`
	Mode    string `json:"mode"`
}

var modes = []string{"http", "amqp"}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("mode is missing. abaliable modes:", modes)
	}
	mode := os.Args[1]
	switch mode {
	case "http":
		log.Println("Run HTTP Reciver...")
		HttpServer()
	case "amqp":
		log.Println("Run AMQP Reciver...")
		AMQPReciver()
	}

}
