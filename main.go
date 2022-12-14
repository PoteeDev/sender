package main

import (
	"log"
	"os"

	"github.com/PoteeDev/sender/servers"
)

var modes = []string{"http", "amqp"}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("mode is missing. abaliable modes:", modes)
	}
	mode := os.Args[1]
	switch mode {
	case "http":
		log.Println("Run HTTP Reciver...")
		servers.HttpServer()
	case "amqp":
		log.Println("Run AMQP Reciver...")
		servers.AMQPReciver()
	}

}
