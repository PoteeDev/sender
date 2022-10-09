package servers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PoteeDev/sender/providers"
)

type JsonMessage struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
	Mode      string `json:"mode"`
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	var message JsonMessage
	// decode json data
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, fmt.Sprintf("cant't decode json data: %s", err.Error()), http.StatusBadRequest)
		return
	}
	// init tg provider and send message
	providers.NewProvider().Send(message.Recipient, message.Message, message.Mode)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok\n")
}

func HttpServer() {
	http.HandleFunc("/send/", sendMessage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
