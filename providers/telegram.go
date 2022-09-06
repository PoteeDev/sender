package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TelegramMessage struct {
	BotToken string
	ChatID   string
	Message  string
	Mode     string
}

func NewMessage(botToken, chatId, message, mode string) *TelegramMessage {
	if botToken == "" {
		log.Fatal("BotToken is empty")
	}
	return &TelegramMessage{
		BotToken: botToken,
		ChatID:   chatId,
		Message:  message,
		Mode:     mode,
	}
}

func (tm *TelegramMessage) Send() error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tm.BotToken)
	body, _ := json.Marshal(map[string]string{
		"chat_id":    tm.ChatID,
		"text":       tm.Message,
		"parce_mode": tm.Mode,
	})
	_, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	log.Println("send message to:", tm.ChatID)
	return nil
}
