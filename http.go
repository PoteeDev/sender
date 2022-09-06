package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PoteeDev/sender/providers"
)

func sendMessage(w http.ResponseWriter, r *http.Request) {
	providerName := strings.TrimPrefix(r.URL.Path, "/send/")

	var message JsonMessage
	// decode json data
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, fmt.Sprintf("cant't decode json data: %s", err.Error()), http.StatusBadRequest)
		return
	}
	// init tg provider and send message
	switch providerName {
	case "telegram":
		token := os.Getenv("BOT_TOKEN")
		err := providers.NewMessage(token, message.ChatID, message.Message, "").Send()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok\n")
}

func HttpServer() {
	http.HandleFunc("/send", sendMessage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
